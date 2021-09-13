package autocode

import (
	"gorm.io/gorm"
	"os"
	"path"
	"path/filepath"
	"text/template"
)

var typeMap map[string]string
var baseColMap map[string]bool

func init() {
	typeMap = make(map[string]string)
	typeMap["int"] = "int"
	typeMap["bigint"] = "int64"
	typeMap["varchar"] = "string"
	typeMap["tinyint"] = "int8"
	typeMap["datetime"] = "time.Time"
	typeMap["bit"] = "bool"
	typeMap["char"] = "string"
	typeMap["date"] = "time.Time"

	baseColMap = make(map[string]bool)
	baseColMap["is_deleted"] = true
	baseColMap["modifier"] = true
	baseColMap["modified"] = true
	baseColMap["creator"] = true
	baseColMap["created"] = true
}

type Column struct {
	ColumnName string `gorm:"COLUMN_NAME"`
	DataType   string `gorm:"DATA_TYPE"`
	Comment    string `gorm:"COLUMN_COMMENT"`
	Length     int64  `gorm:"CHARACTER_MAXIMUM_LENGTH"`
}

func (c *Column) TableName() string {
	return "COLUMNS"
}

func NewDatabaseModel(dao *gorm.DB, database, baseDir, modPath string) error {
	tables := listTable(dao, database)
	if len(tables) > 0 {
		for _, table := range tables {
			if table != "table_name" {
				err := NewTable(dao, database, table, baseDir, modPath)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func NewTable(dao *gorm.DB, database, table, baseDir, modPath string) error {
	err := NewModel(dao, database, table, baseDir, modPath)
	if err != nil {
		return err
	}
	err = NewDao(table, baseDir, modPath)
	if err != nil {
		return err
	}
	err = NewBiz(table, baseDir, modPath)
	if err != nil {
		return err
	}
	err = NewApi(table, baseDir, modPath)
	if err != nil {
		return err
	}
	return nil
}

func NewModel(dao *gorm.DB, database, table, baseDir, modPath string) error {
	funcMap := template.FuncMap{
		"camel":    camelName,
		"typeName": typeName,
	}
	filePath := filepath.Join(baseDir, "model", table+".go")
	err := os.MkdirAll(path.Dir(filePath), os.ModePerm)
	if err != nil {
		return err
	}
	out, err := os.Create(filePath)
	if err != nil {
		return err
	}
	tpl, err := template.New("model").Funcs(funcMap).Parse(modelTemp)
	if err != nil {
		return err
	}
	tpl.Funcs(funcMap)
	columns := listColumns(dao, database, table)
	var cols []Column
	for _, c := range columns {
		if ok := baseColMap[c.ColumnName]; !ok {
			cols = append(cols, c)
		}
	}
	err = tpl.Execute(out, struct {
		Table   string
		ModPath string
		Columns []Column
	}{Table: table, ModPath: modPath, Columns: cols})
	if err != nil {
		return err
	}
	return nil
}

func NewDao(table, baseDir, ModPath string) error {
	funcMap := template.FuncMap{
		"camel":    camelName,
		"typeName": typeName,
	}
	filePath := filepath.Join(baseDir, "dao", table+"_dao.go")
	err := os.MkdirAll(path.Dir(filePath), os.ModePerm)
	if err != nil {
		return err
	}
	out, err := os.Create(filePath)
	if err != nil {
		return err
	}
	tpl, err := template.New("dao").Funcs(funcMap).Parse(daoTemp)
	if err != nil {
		return err
	}
	tpl.Funcs(funcMap)
	err = tpl.Execute(out, struct {
		Table   string
		ModPath string
	}{Table: table, ModPath: ModPath})
	if err != nil {
		return err
	}
	return nil
}

func NewBiz(table, baseDir, ModPath string) error {
	funcMap := template.FuncMap{
		"camel":    camelName,
		"typeName": typeName,
	}
	filePath := filepath.Join(baseDir, "biz", table+"_biz.go")
	err := os.MkdirAll(path.Dir(filePath), os.ModePerm)
	if err != nil {
		return err
	}
	out, err := os.Create(filePath)
	if err != nil {
		return err
	}
	tpl, err := template.New("dao").Funcs(funcMap).Parse(bizTemp)
	if err != nil {
		return err
	}
	tpl.Funcs(funcMap)
	err = tpl.Execute(out, struct {
		Table   string
		ModPath string
	}{Table: table, ModPath: ModPath})
	if err != nil {
		return err
	}
	return nil
}

func NewApi(table, baseDir, ModPath string) error {
	funcMap := template.FuncMap{
		"camel":    camelName,
		"typeName": typeName,
	}
	filePath := filepath.Join(baseDir, "api", table+"_api.go")
	err := os.MkdirAll(path.Dir(filePath), os.ModePerm)
	if err != nil {
		return err
	}
	out, err := os.Create(filePath)
	if err != nil {
		return err
	}
	tpl, err := template.New("api").Funcs(funcMap).Parse(apiTemp)
	if err != nil {
		return err
	}
	tpl.Funcs(funcMap)
	err = tpl.Execute(out, struct {
		Table   string
		ModPath string
	}{Table: table, ModPath: ModPath})
	if err != nil {
		return err
	}
	return nil
}

func listTable(dao *gorm.DB, database string) []string {
	var tables []string
	err := dao.Raw("select table_name from information_schema.tables where table_schema=?",
		database).Find(&tables).Error
	if err != nil {
		panic(err)
	}
	return tables
}

func listColumns(dao *gorm.DB, database, table string) []Column {
	var columns []Column
	err := dao.Raw("select column_name, data_type, CHARACTER_MAXIMUM_LENGTH length, COLUMN_COMMENT comment "+
		"from information_schema.columns where table_schema =? and table_name=?", database, table).Find(&columns).Error
	if err != nil {
		panic(err)
	}
	return columns
}

func typeName(str string, length int64) string {
	if str == "tinyint" && length == 1 {
		return "bool"
	}
	return typeMap[str]
}

func camelName(str string, firstCap bool) string {
	var v string
	for i, c := range []rune(str) {
		if c != '_' {
			if i == 0 && firstCap {
				if c >= 97 && c <= 122 {
					v += string(c - 32)
				} else {
					v += string(c)
				}
			} else if i > 0 && str[i-1] == '_' {
				if c >= 97 && c <= 122 {
					v += string(c - 32)
				} else {
					v += string(c)
				}
			} else {
				v += string(c)
			}
		}
	}
	return v
}

func capFirst(str string) string {
	var v string
	runes := []rune(str)
	if str[0] >= 97 && str[0] <= 122 {
		v += string(str[0] - 32)
	}
	for i, c := range runes {
		if i > 0 {
			v += string(c)
		}
	}
	return v
}
