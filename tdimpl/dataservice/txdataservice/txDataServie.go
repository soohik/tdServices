// Package txdataservice represents transaction support on data service layer
//事务
package txdataservice

import (
	"tdimpl/dataservice"
	"tdimpl/tool/gdbc"
)

// TxDataSql is the generic implementation for transaction for SQL database
// You only need to do it once for each SQL database
type TxDataSql struct {
	DB gdbc.SqlGdbc
}

func (tds *TxDataSql) TxEnd(txFunc func() error) error {
	return tds.DB.TxEnd(txFunc)
}

func (tds *TxDataSql) TxBegin() (dataservice.TxDataInterface, error) {

	sqlTx, error := tds.DB.TxBegin()
	tdi := TxDataSql{sqlTx}
	tds.DB = tdi.DB
	return &tdi, error
}
func (tds *TxDataSql) GetTx() gdbc.SqlGdbc {
	return tds.DB
}
