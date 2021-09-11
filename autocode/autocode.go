package autocode

import (
	"gorm.io/gorm"
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"
)

var typeMap map[string]string

func init() {
	typeMap = make(map[string]string)
	typeMap["int"] = "int"
	typeMap["bigint"] = "int64"
	typeMap["varchar"] = "string"
	typeMap["tinyint"] = "int8"
	typeMap["datetime"] = "time.Time"
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

func NewDatabaseModel(dao *gorm.DB, database, modelDir string) error {
	tables := listTable(dao, database)
	if len(tables) > 0 {
		for _, table := range tables {
			err := NewModel(dao, database, table, modelDir)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func NewModel(dao *gorm.DB, database, table, modelDir string) error {
	funcMap := template.FuncMap{
		"camel":    camelName,
		"typeName": typeName,
	}
	f, err := os.Open("model.txt")
	if err != nil {
		return err
	}
	out, err := os.Create(filepath.Join(modelDir, table+".go"))
	if err != nil {
		return err
	}
	bs, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	tpl, err := template.New("test").Funcs(funcMap).Parse(string(bs))
	if err != nil {
		return err
	}
	tpl.Funcs(funcMap)
	columns := listColumns(dao, database, table)
	err = tpl.Execute(out, struct {
		Table   string
		Columns []Column
	}{Table: table, Columns: columns})
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
