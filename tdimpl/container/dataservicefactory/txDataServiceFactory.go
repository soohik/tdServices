package dataservicefactory

import (
	"tdimpl/config"
	"tdimpl/container"
	"tdimpl/container/datastorefactory"
	"tdimpl/dataservice/txdataservice"
	"tdimpl/tool/gdbc"

	"github.com/pkg/errors"
)

// txDataServiceFactory is a empty receiver for Build method
type txDataServiceFactory struct{}

func (tdsf *txDataServiceFactory) Build(c container.Container, dataConfig *config.DataConfig) (DataServiceInterface, error) {
	//logger.Log.Debug("txDataServiceFactory")
	dsc := dataConfig.DataStoreConfig
	dsi, err := datastorefactory.GetDataStoreFb(dsc.Code).Build(c, &dsc)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	ds := dsi.(gdbc.SqlGdbc)
	tdm := txdataservice.TxDataSql{ds}
	//logger.Log.Debug("udm:", tdm.DB)
	return &tdm, nil

}
