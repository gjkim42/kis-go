package client

import (
	"net/http"

	"github.com/gjkim42/kis-go/uapi"
)

const (
	mockDomain = "https://openapivts.koreainvestment.com:29443"
)

type clientset struct {
	uapi *uapi.UAPI
}

func New(httpclient *http.Client, domain string, appKey, appSecret, accessToken string) *clientset {
	if httpclient == nil {
		httpclient = http.DefaultClient
	}
	if domain == "" {
		domain = mockDomain
	}
	return &clientset{
		uapi: uapi.New(httpclient, domain, appKey, appSecret, accessToken),
	}
}

func (c *clientset) UAPI() *uapi.UAPI {
	return c.uapi
}
