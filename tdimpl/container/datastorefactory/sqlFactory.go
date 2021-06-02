package datastorefactory

import (
	"database/sql"
	"tdimpl/config"
	"tdimpl/tool/gdbc/databasehandler"

	"tdimpl/container"

	"github.com/pkg/errors"
)

// sqlFactory is receiver for Build method
type sqlFactory struct{}

// implement Build method for SQL database
func (sf *sqlFactory) Build(c container.Container, dsc *config.DataStoreConfig) (DataStoreInterface, error) {
	//logger.Log.Debug("sqlFactory")
	key := dsc.Code
	//if it is already in container, return
	if value, found := c.Get(key); found {
		//logger.Log.Debugf("found db value %+v\n:", value)
		sdb := value.(*sql.DB)
		sdt := databasehandler.SqlDBTx{DB: sdb}
		//l/ogger.Log.Debug("found db in container for key:", key)
		return &sdt, nil
	}

	db, err := sql.Open(dsc.DriverName, dsc.UrlAddress)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	// check the connection
	err = db.Ping()
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	dt := databasehandler.SqlDBTx{DB: db}
	c.Put(key, db)
	return &dt, nil

}
