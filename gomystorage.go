package gomystorage

import (
	"net/http"
	"net/url"
	"path"
)

const host = "api.moysklad.ru"

type ApiClient struct {
	baseUrl url.URL
	client  http.Client
	token   string
}

func New(token string) *ApiClient {
	baseUrl := url.URL{
		Scheme: "https",
		Host:   host,
		Path:   path.Join("/api", "remap", "1.2"),
	}
	return &ApiClient{
		baseUrl: baseUrl,
		client:  http.Client{},
		token:   token,
	}
}
