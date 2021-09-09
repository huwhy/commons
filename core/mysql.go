package core

import (
	"git.huwhy.cn/huwhy/commons/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func NewSQL(conf *config.Mysql) *gorm.DB {
	mysqlConfig := mysql.Config{
		DSN:                       conf.Dsn(), // DSN data source name
		DefaultStringSize:         191,        // string 类型字段的默认长度
		DisableDatetimePrecision:  true,       // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,       // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,       // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,      // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
	}); err != nil {
		return nil
	} else {
		sqlDB, _ := db.DB()
		db.Logger = logger.Default.LogMode(logger.Info)
		sqlDB.SetMaxIdleConns(conf.MaxIdle)
		sqlDB.SetMaxOpenConns(conf.MaxOpen)
		return db
	}
}
