package apiv2

import (
	"net/http"
	"time"
)

type API struct {
	baseURL    string
	token      string
	ownerSlug  string
	httpClient *http.Client
}

func New(token, ownerSlug string, httpTimeout time.Duration) API {
	t := http.DefaultTransport.(*http.Transport).Clone()
	t.MaxIdleConnsPerHost = 10

	httpClient := &http.Client{
		Timeout:   httpTimeout,
		Transport: t,
	}

	return API{
		baseURL:    "https://circleci.com/api/v2",
		token:      token,
		ownerSlug:  ownerSlug,
		httpClient: httpClient,
	}
}
