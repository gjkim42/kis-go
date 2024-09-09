package v1

import (
	"net/http"

	"github.com/gjkim42/kis-go/uapi/overseas_stock/v1/trading"
)

type DomesticStock struct {
	trading trading.Interface
}

func New(httpclient *http.Client, url string, appKey, appSecret, accessToken string) *DomesticStock {
	return &DomesticStock{
		trading: trading.New(httpclient, url+"/domestic-stock/v1", appKey, appSecret, accessToken),
	}
}

func (o *DomesticStock) Trading() trading.Interface {
	return o.trading
}
