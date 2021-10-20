package autocode

import (
	"github.com/huwhy/commons/autocode/model"
	"github.com/huwhy/commons/config"
	"github.com/huwhy/commons/core"
	"gorm.io/gorm"
	"path/filepath"
	"testing"
)

func TestListTable(t *testing.T) {
	tables := listTable(getDao(), "trend")
	t.Log(tables)
}

func TestListColumn(t *testing.T) {
	columns := listColumns(getDao(), "trend", "shares_day_data")
	t.Log(columns)
}

func TestCamel(t *testing.T) {
	t.Log(camelName("a_test_adf", false))
}

func TestCapFirst(t *testing.T) {
	t.Log(capFirst("a_test_adf"))
}

func getDao() *gorm.DB {
	var mysql = &config.Mysql{
		Host:     "localhost",
		Username: "root",
		Password: "abc123",
		Database: "trend",
		Params:   "charset=utf8mb4&parseTime=True&loc=Local",
		MaxIdle:  5,
		MaxOpen:  5,
	}
	return core.NewSQL(mysql)
}

func TestModel(t *testing.T) {
	err := NewModel(getDao(), "trend", "shares", ".", "huwhy.cn/demo")
	if err != nil {
		t.Fatal(err)
	}
	t.Log("success")
}

func TestDao(t *testing.T) {
	baseDir, err := filepath.Abs(".")
	if err != nil {
		panic(err)
	}
	err = NewDao("member", baseDir, "huwhy.cn/demo")
	if err != nil {
		t.Fatal(err)
	}
}

func TestBiz(t *testing.T) {
	baseDir, err := filepath.Abs(".")
	if err != nil {
		panic(err)
	}
	err = NewBiz("member", baseDir, "huwhy.cn/demo")
	if err != nil {
		t.Fatal(err)
	}
}

func TestApi(t *testing.T) {
	baseDir, err := filepath.Abs(".")
	if err != nil {
		panic(err)
	}
	err = NewApi("member", baseDir, "huwhy.cn/demo")
	if err != nil {
		t.Fatal(err)
	}
}

func TestNewTable(t *testing.T) {
	err := NewTable(getDao(), "trend", "shares_day_data", ".", "huwhy.cn/demo")
	if err != nil {
		t.Fatal(err)
	}
}

func TestNewDatabaseModel(t *testing.T) {
	err := NewDatabaseModel(getDao(), "trend", ".", "huwhy.cn/demo")
	if err != nil {
		t.Fatal(err)
	}
	t.Log("success")
}

func TestTestWhereIn(t *testing.T) {
	sql := "select * from shares where id in ?"
	var ids []int
	ids = append(ids, 1, 2, 3)
	var pos []model.Shares
	rs := getDao().Raw(sql, ids).Find(&pos)
	if rs.Error != nil {
		t.Error(rs.Error)
	}
	t.Log(pos)

	getDao().Model(&model.Shares{}).Where("id in ?", ids).Find(&pos)
	t.Log(pos)
}
