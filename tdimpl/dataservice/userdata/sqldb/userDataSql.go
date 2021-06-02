// Package sql represents SQL database implementation of the user data persistence layer
package sqldb

import (
	"tdimpl/dataservice"
	"tdimpl/model"
	"tdimpl/tool/gdbc"

	_ "github.com/go-sql-driver/mysql"

	"github.com/pkg/errors"
)

const (
	DELETE_USER        string = "delete from userinfo where username=?"
	QUERY_USER_BY_ID   string = "SELECT * FROM userinfo where uid =?"
	QUERY_USER_BY_NAME        = "SELECT * FROM userinfo where username =?"
	QUERY_USER                = "SELECT * FROM userinfo "
	UPDATE_USER               = "update userinfo set username=?, department=?, created=? where uid=?"
	INSERT_USER               = "INSERT userinfo SET username=?,department=?,created=?"
)

// UserDataSql is the SQL implementation of UserDataInterface
type UserDataSql struct {
	DB gdbc.SqlGdbc
}

func (uds *UserDataSql) Insert(user *model.Phone) (*model.Phone, error) {

	stmt, err := uds.DB.Prepare(INSERT_USER)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	defer stmt.Close()
	res, err := stmt.Exec(user.Phone)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	user.Id = int(id)

	return user, nil
}

func (uds *UserDataSql) EnableTx(tx dataservice.TxDataInterface) {
	uds.DB = tx.GetTx()
}
