package v1

import (
	"net/http"

	"github.com/gjkim42/kis-go/uapi/overseas_stock/v1/trading"
)

type OverseasStock struct {
	trading trading.Interface
}

func New(httpclient *http.Client, url string, appKey, appSecret, accessToken string) *OverseasStock {
	return &OverseasStock{
		trading: trading.New(httpclient, url+"/overseas-stock/v1", appKey, appSecret, accessToken),
	}
}

func (o *OverseasStock) Trading() trading.Interface {
	return o.trading
}
