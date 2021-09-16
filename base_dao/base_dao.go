package base_dao

import (
	"fmt"
	"github.com/huwhy/commons/basemodel"
	"github.com/huwhy/commons/errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type BaseDao struct {
	DB  *gorm.DB
	LOG *zap.SugaredLogger
}

func NewBaseDao(db *gorm.DB, log *zap.SugaredLogger) BaseDao {
	return BaseDao{
		DB:  db,
		LOG: log,
	}
}

func (dao *BaseDao) HandleErr(err error) {
	if err != nil {
		panic(errors.NewDaoError(err.Error()))
	}
}

func (dao *BaseDao) Count(sql string, args []interface{}) (int, error) {
	countSql := "select count(1) from (" + sql + ") temp"
	var count int
	dao.LOG.Info("dao.paging.count start: ", countSql, args)
	rs := dao.DB.Raw(countSql, args...).Scan(&count)
	if rs.Error != nil {
		return 0, rs.Error
	}
	return count, nil
}

func (dao *BaseDao) Paging(term *basemodel.Term, sql, orderBy string, args []interface{}, data interface{}) error {
	dao.LOG.Info("dao.paging.paging start", sql, orderBy, args)
	var err error
	if term.LastId == 0 {
		term.Total, err = dao.Count(sql, args)
		if err != nil {
			return err
		}
	}
	if orderBy != "" {
		sql += " order by " + orderBy
	}
	sql += fmt.Sprintf(" limit %v, %v", term.GetOffset(), term.Size)
	dao.LOG.Info("dao.paging.find start")
	rs := dao.DB.Raw(sql, args...).Find(data)
	if rs.Error != nil {
		return rs.Error
	}
	dao.LOG.Info("dao.paging.find end")
	return nil
}

func (dao *BaseDao) List(sql, orderBy string, limit int, args []interface{}, data interface{}) error {
	if orderBy != "" {
		sql += " order by " + orderBy
	}
	sql += fmt.Sprintf(" limit 0, %d", limit)
	rs := dao.DB.Raw(sql, args...).Find(data)
	if rs.Error != nil {
		return rs.Error
	}
	return nil
}

func (dao *BaseDao) GetOneById(id int, po interface{}) (bool, error) {
	rs := dao.DB.Where("id=?", id).First(po)
	if rs.Error != nil {
		if rs.Error == gorm.ErrRecordNotFound {
			return false, nil
		} else {
			return false, rs.Error
		}
	}
	return true, nil
}

func (dao *BaseDao) GetOneById64(id int64, po interface{}) (bool, error) {
	rs := dao.DB.Where("id=?", id).First(po)
	if rs.Error != nil {
		if rs.Error == gorm.ErrRecordNotFound {
			return false, nil
		} else {
			return false, rs.Error
		}
	}
	return true, nil
}

func (dao *BaseDao) SaveById(id int, po interface{}) bool {
	var rs *gorm.DB
	if id > 0 {
		rs = dao.DB.Model(po).Updates(po)
	} else {
		rs = dao.DB.Create(po)
	}
	dao.HandleErr(rs.Error)
	return rs.RowsAffected > 0
}

func (dao *BaseDao) SaveById64(id int64, po interface{}) bool {
	var rs *gorm.DB
	if id > 0 {
		rs = dao.DB.Model(po).Updates(po)
	} else {
		rs = dao.DB.Create(po)
	}
	dao.HandleErr(rs.Error)
	return rs.RowsAffected > 0
}

func (dao *BaseDao) DeleteInt64(id int64, table string) bool {
	rs := dao.DB.Exec("update "+table+" set is_deleted=1 where id=?", id)
	dao.HandleErr(rs.Error)
	return rs.RowsAffected > 0
}

func (dao *BaseDao) DeleteInt(id int, table string) bool {
	rs := dao.DB.Exec("update "+table+" set is_deleted=1 where id=?", id)
	dao.HandleErr(rs.Error)
	return rs.RowsAffected > 0
}

func InExp(length int) string {
	inSql := "("
	for i := 0; i < length; i++ {
		if i > 0 {
			inSql += ",?"
		} else {
			inSql += "?"
		}
	}
	inSql += ")"
	return inSql
}
