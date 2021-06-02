// Package sql represents SQL database implementation of the user data persistence layer
package sqldb

import (
	"database/sql"
	"tdapi/config"
	"tdapi/container/logger"
	"tdapi/dataservice"
	"tdapi/model"
	"tdapi/tool/gdbc"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

const (
	DELETE_USER        string = "delete from userinfo where username=?"
	QUERY_USER_BY_ID   string = "SELECT * FROM userinfo where uid =?"
	QUERY_USER_BY_NAME        = "SELECT * FROM userinfo where username =?"
	QUERY_ALLCLIENT           = "SELECT phone,account,tddata,tdfile,create_time,update_time from td.phone"
	//QUERY_ALLCLIENT = "SELECT phone,account from td.phone"
	UPDATE_USER = "update userinfo set username=?, department=?, created=? where uid=?"
	INSERT_USER = "INSERT userinfo SET username=?,department=?,created=?"
)

// UserDataSql is the SQL implementation of UserDataInterface
type UserDataSql struct {
	DB gdbc.SqlGdbc
}

func (uds *UserDataSql) InitUseCase(appConfig *config.AppConfig) {

}

func (uds *UserDataSql) Remove(username string) (int64, error) {

	stmt, err := uds.DB.Prepare(DELETE_USER)
	if err != nil {
		return 0, errors.Wrap(err, "")
	}
	defer stmt.Close()

	res, err := stmt.Exec(username)
	if err != nil {
		return 0, errors.Wrap(err, "")
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "")
	}

	logger.Log.Debug("remove:row affected ", rowsAffected)
	return rowsAffected, nil
}

func (uds *UserDataSql) Find(id int) (*model.User, error) {
	return nil, nil
	// rows, err := uds.DB.Query(QUERY_USER_BY_ID, id)
	// if err != nil {
	// 	return nil, errors.Wrap(err, "")
	// }
	// defer rows.Close()
	// return retrieveUser(rows)
}
func retrieveUser(rows *sql.Rows) (*model.Phone, error) {
	if rows.Next() {
		return rowsToUser(rows)
	}
	return nil, nil
}

func rowsToUser(rows *sql.Rows) (*model.Phone, error) {

	user := &model.Phone{}
	err := rows.Scan(&user)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	// created, err := time.Parse(tool.FORMAT_ISO8601_DATE, ds)
	// if err != nil {
	// 	return nil, errors.Wrap(err, "")
	// }

	logger.Log.Debug("rows to User:", user)
	return user, nil
}
func (uds *UserDataSql) FindByName(name string) (*model.User, error) {
	return nil, nil
	// //logger.Log.Debug("call FindByName() and name is:", name)
	// rows, err := uds.DB.Query(QUERY_USER_BY_NAME, name)
	// if err != nil {
	// 	return nil, errors.Wrap(err, "")
	// }
	// defer rows.Close()
	// return retrieveUser(rows)
}

func (uds *UserDataSql) FindAll() ([]model.User, error) {

	return nil, nil
}

func (uds *UserDataSql) Update(user *model.User) (int64, error) {

	stmt, err := uds.DB.Prepare(UPDATE_USER)

	if err != nil {
		return 0, errors.Wrap(err, "")
	}
	defer stmt.Close()
	res, err := stmt.Exec(user.Name, user.Department, user.Created, user.Id)
	if err != nil {
		return 0, errors.Wrap(err, "")
	}
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		return 0, errors.Wrap(err, "")
	}
	logger.Log.Debug("update: rows affected: ", rowsAffected)

	return rowsAffected, nil
}

func (uds *UserDataSql) Insert(user *model.User) (*model.User, error) {

	stmt, err := uds.DB.Prepare(INSERT_USER)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	defer stmt.Close()
	res, err := stmt.Exec(user.Name, user.Department, user.Created)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	user.Id = int(id)
	logger.Log.Debug("user inserted:", user)
	return user, nil
}

func (uds *UserDataSql) GetAllClient() ([]model.Phone, error) {
	rows, err := uds.DB.Query(QUERY_ALLCLIENT)
	if err != nil {
		logger.Log.Errorf("UserDataSql interface GetAllClient failed:%+v\n", err)
		return nil, errors.Wrap(err, "")
	}
	defer rows.Close()
	phones := []model.Phone{}

	for rows.Next() {
		phone := model.Phone{}
		uds.DB.ScanRows(rows, &phone)
		phones = append(phones, phone)
	}

	if err != nil {
		logger.Log.Errorf("rowsToUser failed:%+v\n", err)
		return phones, errors.Wrap(err, "")
	}
	//need to check error for rows.Next()
	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, "")
	}
	logger.Log.Debug("find user list:", phones)
	return phones, nil
}

func (uds *UserDataSql) EnableTx(tx dataservice.TxDataInterface) {
	uds.DB = tx.GetTx()
}
