package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gjkim42/kis-go/rest"
)

type OAuth2 interface {
	TokenP(ctx context.Context, typ, appKey, appSecret string) (*TokenPResponse, error)
}

type oAuth2 struct {
	restclient *rest.Client
}

func NewOAuth2(httpclient *http.Client, domain string) OAuth2 {
	if httpclient == nil {
		httpclient = http.DefaultClient
	}
	if domain == "" {
		domain = mockDomain
	}
	return &oAuth2{
		restclient: rest.NewClient(httpclient, domain+"/oauth2", "", "", rest.ClientOptions{}),
	}
}

type tokenPBody struct {
	GrantType string `json:"grant_type"`
	AppKey    string `json:"appkey"`
	AppSecret string `json:"appsecret"`
}

type TokenPResponse struct {
	AccessToken             string `json:"access_token"`
	TokenType               string `json:"token_type"`
	ExpiresIn               int    `json:"expires_in"`
	AccessTokenTokenExpired string `json:"access_token_token_expired"`
}

func (o *oAuth2) TokenP(ctx context.Context, typ, appKey, appSecret string) (*TokenPResponse, error) {
	body := &tokenPBody{
		GrantType: typ,
		AppKey:    appKey,
		AppSecret: appSecret,
	}
	res, err := o.restclient.Post().At("tokenP").Body(body).Do(ctx)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		b, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("failed to get token: %s, code: %d", string(b), res.StatusCode)
	}

	var response TokenPResponse
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, err
	}

	return &response, nil
}
