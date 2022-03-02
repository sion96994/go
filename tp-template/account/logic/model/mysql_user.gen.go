// Code generated by 'micro gen' command.
// DO NOT EDIT!

package model

import (
	"database/sql"
	"unsafe"

	"github.com/henrylee2cn/goutil/coarsetime"
	tp "github.com/henrylee2cn/teleport"
	"github.com/xiaoenai/tp-micro/model/mysql"
	"github.com/xiaoenai/tp-micro/model/sqlx"

	"github.com/sion96994/go/tp-template/account/args"
)

// User user info
type User args.User

// ToUser converts to *User type.
func ToUser(_u *args.User) *User {
	return (*User)(unsafe.Pointer(_u))
}

// ToArgsUser converts to *args.User type.
func ToArgsUser(_u *User) *args.User {
	return (*args.User)(unsafe.Pointer(_u))
}

// ToUserSlice converts to []*User type.
func ToUserSlice(a []*args.User) []*User {
	return *(*[]*User)(unsafe.Pointer(&a))
}

// ToArgsUserSlice converts to []*args.User type.
func ToArgsUserSlice(a []*User) []*args.User {
	return *(*[]*args.User)(unsafe.Pointer(&a))
}

// TableName implements 'github.com/xiaoenai/tp-micro/model'.Cacheable
func (*User) TableName() string {
	return "user"
}

func (_u *User) isZeroPrimaryKey() bool {
	var _id int64
	if _u.Id != _id {
		return false
	}
	return true
}

var userDB, _ = mysqlHandler.RegCacheableDB(new(User), cacheExpire, args.UserSql)

// GetUserDB returns the User DB handler.
func GetUserDB() *mysql.CacheableDB {
	return userDB
}

// InsertUser insert a User data into database.
// NOTE:
//  Primary key: 'id';
//  Without cache layer.
func InsertUser(_u *User, tx ...*sqlx.Tx) (int64, error) {
	_u.UpdatedAt = coarsetime.FloorTimeNow().Unix()
	if _u.CreatedAt == 0 {
		_u.CreatedAt = _u.UpdatedAt
	}
	return _u.Id, userDB.Callback(func(tx sqlx.DbOrTx) error {
		var (
			query            string
			isZeroPrimaryKey = _u.isZeroPrimaryKey()
		)
		if isZeroPrimaryKey {
			query = "INSERT INTO `user` (`name`,`access_token`,`updated_at`,`created_at`)VALUES(:name,:access_token,:updated_at,:created_at);"
		} else {
			query = "INSERT INTO `user` (`id`,`name`,`access_token`,`updated_at`,`created_at`)VALUES(:id,:name,:access_token,:updated_at,:created_at);"
		}
		r, err := tx.NamedExec(query, _u)
		if isZeroPrimaryKey && err == nil {
			_u.Id, err = r.LastInsertId()
		}
		return err
	}, tx...)
}

// UpsertUser insert or update the User data by primary key.
// NOTE:
//  Primary key: 'id';
//  With cache layer;
//  Insert data if the primary key is specified;
//  Update data based on _updateFields if no primary key is specified;
//  _updateFields' members must be db field style (snake format);
//  Automatic update 'updated_at' field;
//  Don't update the primary keys, 'created_at' key and 'deleted_ts' key;
//  Update all fields except the primary keys, 'created_at' key and 'deleted_ts' key, if _updateFields is empty.
func UpsertUser(_u *User, _updateFields []string, tx ...*sqlx.Tx) (int64, error) {
	if _u.UpdatedAt == 0 {
		_u.UpdatedAt = coarsetime.FloorTimeNow().Unix()
	}
	if _u.CreatedAt == 0 {
		_u.CreatedAt = _u.UpdatedAt
	}
	err := userDB.Callback(func(tx sqlx.DbOrTx) error {
		var (
			query            string
			isZeroPrimaryKey = _u.isZeroPrimaryKey()
		)
		if isZeroPrimaryKey {
			query = "INSERT INTO `user` (`name`,`access_token`,`updated_at`,`created_at`)VALUES(:name,:access_token,:updated_at,:created_at)"
		} else {
			query = "INSERT INTO `user` (`id`,`name`,`access_token`,`updated_at`,`created_at`)VALUES(:id,:name,:access_token,:updated_at,:created_at)"
		}
		query += " ON DUPLICATE KEY UPDATE "
		if len(_updateFields) == 0 {
			query += "`name`=VALUES(`name`),`access_token`=VALUES(`access_token`),`updated_at`=VALUES(`updated_at`);"
		} else {
			for _, s := range _updateFields {
				if s == "updated_at" || s == "created_at" || s == "deleted_ts" || s == "id" {
					continue
				}
				query += "`" + s + "`=VALUES(`" + s + "`),"
			}
			if query[len(query)-1] != ',' {
				return nil
			}
			query += "`updated_at`=VALUES(`updated_at`),`deleted_ts`=0;"
		}
		r, err := tx.NamedExec(query, _u)
		if isZeroPrimaryKey && err == nil {
			var rowsAffected int64
			rowsAffected, err = r.RowsAffected()
			if rowsAffected == 1 {
				_u.Id, err = r.LastInsertId()
			}
		}
		return err
	}, tx...)
	if err != nil {
		return _u.Id, err
	}
	err = userDB.DeleteCache(_u)
	if err != nil {
		tp.Errorf("%s", err.Error())
	}
	return _u.Id, nil
}

// UpdateUserByPrimary update the User data in database by primary key.
// NOTE:
//  Primary key: 'id';
//  With cache layer;
//  _updateFields' members must be db field style (snake format);
//  Automatic update 'updated_at' field;
//  Don't update the primary keys, 'created_at' key and 'deleted_ts' key;
//  Update all fields except the primary keys, 'created_at' key and 'deleted_ts' key, if _updateFields is empty.
func UpdateUserByPrimary(_u *User, _updateFields []string, tx ...*sqlx.Tx) error {
	_u.UpdatedAt = coarsetime.FloorTimeNow().Unix()
	err := userDB.Callback(func(tx sqlx.DbOrTx) error {
		query := "UPDATE `user` SET "
		if len(_updateFields) == 0 {
			query += "`name`=:name,`access_token`=:access_token,`updated_at`=:updated_at WHERE `id`=:id AND `deleted_ts`=0 LIMIT 1;"
		} else {
			for _, s := range _updateFields {
				if s == "updated_at" || s == "created_at" || s == "deleted_ts" || s == "id" {
					continue
				}
				query += "`" + s + "`=:" + s + ","
			}
			if query[len(query)-1] != ',' {
				return nil
			}
			query += "`updated_at`=:updated_at WHERE `id`=:id AND `deleted_ts`=0 LIMIT 1;"
		}
		_, err := tx.NamedExec(query, _u)
		return err
	}, tx...)
	if err != nil {
		return err
	}
	err = userDB.DeleteCache(_u)
	if err != nil {
		tp.Errorf("%s", err.Error())
	}
	return nil
}

// UpdateUserByName update the User data in database by 'name' unique key.
// NOTE:
//  With cache layer;
//  _updateFields' members must be db field style (snake format);
//  Automatic update 'updated_at' field;
//  Don't update the primary keys, 'created_at' key and 'deleted_ts' key;
//  Update all fields except the primary keys, 'name' unique key, 'created_at' key and 'deleted_ts' key, if _updateFields is empty.
func UpdateUserByName(_u *User, _updateFields []string, tx ...*sqlx.Tx) error {
	_u.UpdatedAt = coarsetime.FloorTimeNow().Unix()
	err := userDB.Callback(func(tx sqlx.DbOrTx) error {
		query := "UPDATE `user` SET "
		if len(_updateFields) == 0 {
			query += "`name`=:name,`access_token`=:access_token,`updated_at`=:updated_at WHERE `name`=:name AND `deleted_ts`=0 LIMIT 1;"
		} else {
			for _, s := range _updateFields {
				if s == "updated_at" || s == "created_at" || s == "deleted_ts" || s == "name" || s == "id" {
					continue
				}
				query += "`" + s + "`=:" + s + ","
			}
			if query[len(query)-1] != ',' {
				return nil
			}
			query += "`updated_at`=:updated_at WHERE `name`=:name AND `deleted_ts`=0 LIMIT 1;"
		}
		_, err := tx.NamedExec(query, _u)
		return err
	}, tx...)
	if err != nil {
		return err
	}
	err = userDB.DeleteCache(_u, "name")
	if err != nil {
		tp.Errorf("%s", err.Error())
	}
	return nil
}

// DeleteUserByPrimary delete a User data in database by primary key.
// NOTE:
//  Primary key: 'id';
//  With cache layer.
func DeleteUserByPrimary(_id int64, deleteHard bool, tx ...*sqlx.Tx) error {
	var err error
	if deleteHard {
		// Immediately delete from the hard disk.
		err = userDB.Callback(func(tx sqlx.DbOrTx) error {
			_, err := tx.Exec("DELETE FROM `user` WHERE `id`=? AND `deleted_ts`=0;", _id)
			return err
		}, tx...)

	} else {
		// Delay delete from the hard disk.
		ts := coarsetime.FloorTimeNow().Unix()
		err = userDB.Callback(func(tx sqlx.DbOrTx) error {
			_, err := tx.Exec("UPDATE `user` SET `updated_at`=?, `deleted_ts`=? WHERE `id`=? AND `deleted_ts`=0;", ts, ts, _id)
			return err
		}, tx...)
	}

	if err != nil {
		return err
	}
	err = userDB.DeleteCache(&User{
		Id: _id,
	})
	if err != nil {
		tp.Errorf("%s", err.Error())
	}
	return nil
}

// DeleteUserByName delete a User data in database by 'name' unique key.
// NOTE:
//  With cache layer.
func DeleteUserByName(_name string, deleteHard bool, tx ...*sqlx.Tx) error {
	var err error
	if deleteHard {
		// Immediately delete from the hard disk.
		err = userDB.Callback(func(tx sqlx.DbOrTx) error {
			_, err := tx.Exec("DELETE FROM `user` WHERE `name`=? AND `deleted_ts`=0;", _name)
			return err
		}, tx...)

	} else {
		// Delay delete from the hard disk.
		ts := coarsetime.FloorTimeNow().Unix()
		err = userDB.Callback(func(tx sqlx.DbOrTx) error {
			_, err := tx.Exec("UPDATE `user` SET `updated_at`=?, `deleted_ts`=? WHERE `name`=? AND `deleted_ts`=0;", ts, ts, _name)
			return err
		}, tx...)
	}

	if err != nil {
		return err
	}
	err = userDB.DeleteCache(&User{
		Name: _name,
	}, "name")
	if err != nil {
		tp.Errorf("%s", err.Error())
	}
	return nil
}

// GetUserByPrimary query a User data from database by primary key.
// NOTE:
//  Primary key: 'id';
//  With cache layer;
//  If @return bool=false error=nil, means the data is not exist.
func GetUserByPrimary(_id int64) (*User, bool, error) {
	var _u = &User{
		Id: _id,
	}
	err := userDB.CacheGet(_u)
	switch err {
	case nil:
		if _u.CreatedAt == 0 || _u.DeletedTs != 0 {
			return nil, false, nil
		}
		return _u, true, nil
	case sql.ErrNoRows:
		return nil, false, nil
	default:
		return nil, false, err
	}
}

// GetUserByName query a User data from database by 'name' unique key.
// NOTE:
//  With cache layer;
//  If @return bool=false error=nil, means the data is not exist.
func GetUserByName(_name string) (*User, bool, error) {
	var _u = &User{
		Name: _name,
	}
	err := userDB.CacheGet(_u, "name")
	switch err {
	case nil:
		if _u.CreatedAt == 0 || _u.DeletedTs != 0 {
			return nil, false, nil
		}
		return _u, true, nil
	case sql.ErrNoRows:
		return nil, false, nil
	default:
		return nil, false, err
	}
}

// GetUserByWhere query a User data from database by WHERE condition.
// NOTE:
//  Without cache layer;
//  If @return bool=false error=nil, means the data is not exist.
func GetUserByWhere(whereCond string, arg ...interface{}) (*User, bool, error) {
	var _u = new(User)
	err := userDB.Get(_u, "SELECT `id`,`name`,`access_token`,`updated_at`,`created_at` FROM `user` WHERE "+insertZeroDeletedTsField(whereCond)+" LIMIT 1;", arg...)
	switch err {
	case nil:
		return _u, true, nil
	case sql.ErrNoRows:
		return nil, false, nil
	default:
		return nil, false, err
	}
}

// SelectUserByWhere query some User data from database by WHERE condition.
// NOTE:
//  Without cache layer.
func SelectUserByWhere(whereCond string, arg ...interface{}) ([]*User, error) {
	var objs = new([]*User)
	err := userDB.Select(objs, "SELECT `id`,`name`,`access_token`,`updated_at`,`created_at` FROM `user` WHERE "+insertZeroDeletedTsField(whereCond), arg...)
	return *objs, err
}

// CountUserByWhere count User data number from database by WHERE condition.
// NOTE:
//  Without cache layer.
func CountUserByWhere(whereCond string, arg ...interface{}) (int64, error) {
	var count int64
	err := userDB.Get(&count, "SELECT count(*) FROM `user` WHERE "+insertZeroDeletedTsField(whereCond), arg...)
	return count, err
}
