package trading

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/gjkim42/kis-go/rest"
)

type Interface interface {
	Get(ctx context.Context, path string, header map[string]string, query map[string]string) (map[string]any, error)
}

type trading struct {
	restclient *rest.Client
}

func New(httpclient *http.Client, url, appKey, appSecret, accessToken string) *trading {
	return &trading{
		restclient: rest.NewClient(httpclient, url+"/trading", appKey, appSecret, rest.ClientOptions{
			Header: map[string]string{
				"Authorization": "Bearer " + accessToken,
			},
		}),
	}
}

func (t *trading) Get(ctx context.Context, path string, header map[string]string, query map[string]string) (map[string]any, error) {
	slog.Info("Starting Get...", "path", path, "header", header, "query", query)
	defer slog.Info("Finished Get", "path", path, "header", header, "query", query)
	res, err := t.restclient.Get().At(path).Headers(header).Queries(query).Do(ctx)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		b, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("failed to get: %s, path: %s, code: %d", string(b), path, res.StatusCode)
	}

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	data := make(map[string]any)
	err = json.Unmarshal(b, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
