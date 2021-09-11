package autocode

import (
	"git.huwhy.cn/commons/config"
	"git.huwhy.cn/commons/core"
	"gorm.io/gorm"
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

func TestTemplate(t *testing.T) {
	err := NewModel(getDao(), "trend", "shares_day_data", ".", "autocode")
	if err != nil {
		t.Fatal(err)
	}
	t.Log("success")
}

func TestNewDatabaseModel(t *testing.T) {
	err := NewDatabaseModel(getDao(), "trend", ".", "autocode")
	if err != nil {
		t.Fatal(err)
	}
	t.Log("success")
}
