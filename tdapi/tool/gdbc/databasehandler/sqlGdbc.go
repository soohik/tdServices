// Package handles low level database access including transaction through *sql.Tx or *sql.DB

package databasehandler

import (
	"context"
	"database/sql"

	"github.com/go-kivik/kivik"
	"gorm.io/gorm"
)

// SqlDBTx is the concrete implementation of sqlGdbc by using *sql.DB
type SqlDBTx struct {
	DB *gorm.DB
}

// SqlConnTx is the concrete implementation of sqlGdbc by using *sql.Tx
type SqlConnTx struct {
	DB *gorm.DB
}

func (sdt *SqlDBTx) Exec(query string, args ...interface{}) (sql.Result, error) {
	return nil, nil
}

func (sdt *SqlDBTx) Prepare(query string) (*sql.Stmt, error) {
	return nil, nil
}

func (sdt *SqlDBTx) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return sdt.DB.Raw(query, args...).Rows()
}

func (sdt *SqlDBTx) QueryRow(query string, args ...interface{}) *sql.Row {
	return nil
}

func (sdt *SqlDBTx) ScanRows(rows *sql.Rows, dest interface{}) error {
	return sdt.DB.ScanRows(rows, dest)
}

func (sdb *SqlConnTx) Exec(query string, args ...interface{}) (sql.Result, error) {
	return nil, nil
}

func (sdb *SqlConnTx) Prepare(query string) (*sql.Stmt, error) {
	return nil, nil
}

func (sdb *SqlConnTx) Query(query string, args ...interface{}) (*sql.Rows, error) {

	return nil, nil
}

func (sdb *SqlConnTx) QueryRow(query string, args ...interface{}) *sql.Row {
	return nil
}

func (sdb *SqlConnTx) ScanRows(rows *sql.Rows, dest interface{}) error {
	return nil
}

// The followings are dummy implementation for noSqlGdbc.
// After implementing the following, SqlDBTx will also implements NoSqlGdbc interface.
// This making the courseDataServiceFactory possible.
// The functions will never be called, it just makes the interface available
// If you don't need something similar to courseDataServiceFactory, you can remove the following code.
func (sdt *SqlDBTx) QueryNoSql(ctx context.Context, ddoc string, view string) (*kivik.Rows, error) {
	return nil, nil
}

func (sdt *SqlDBTx) Put(ctx context.Context, docID string, doc interface{}, options ...kivik.Options) (rev string, err error) {
	return "", nil
}

func (sdt *SqlDBTx) Get(ctx context.Context, docID string, options ...kivik.Options) (*kivik.Row, error) {
	return nil, nil
}

func (sdt *SqlDBTx) Find(ctx context.Context, query interface{}) (*kivik.Rows, error) {
	return nil, nil
}

func (sdt *SqlDBTx) AllDocs(ctx context.Context, options ...kivik.Options) (*kivik.Rows, error) {
	return nil, nil
}
