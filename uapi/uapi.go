package uapi

import (
	"net/http"

	v1 "github.com/gjkim42/kis-go/uapi/overseas_stock/v1"
)

type UAPI struct {
	overseasstockv1 *v1.OverseasStock
}

func New(httpclient *http.Client, domain string, appKey, appSecret, accessToken string) *UAPI {
	return &UAPI{
		overseasstockv1: v1.New(httpclient, domain+"/uapi", appKey, appSecret, accessToken),
	}
}

func (u *UAPI) OverseasStockV1() *v1.OverseasStock {
	return u.overseasstockv1
}
