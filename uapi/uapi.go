package uapi

import (
	"net/http"

	domesticstockv1 "github.com/gjkim42/kis-go/uapi/domestic-stock/v1"
	overseasstockv1 "github.com/gjkim42/kis-go/uapi/overseas_stock/v1"
)

type UAPI struct {
	overseasstockv1 *overseasstockv1.OverseasStock
	domesticstockv1 *domesticstockv1.DomesticStock
}

func New(httpclient *http.Client, domain string, appKey, appSecret, accessToken string) *UAPI {
	return &UAPI{
		overseasstockv1: overseasstockv1.New(httpclient, domain+"/uapi", appKey, appSecret, accessToken),
		domesticstockv1: domesticstockv1.New(httpclient, domain+"/uapi", appKey, appSecret, accessToken),
	}
}

func (u *UAPI) OverseasStockV1() *overseasstockv1.OverseasStock {
	return u.overseasstockv1
}

func (u *UAPI) DomesticStockV1() *domesticstockv1.DomesticStock {
	return u.domesticstockv1
}
