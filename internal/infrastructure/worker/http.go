package worker

import (
	"net/http"
	"time"
)

type HttpStatistics struct {
	StatusCode   int64  `json:"statusCode"`
	Method       string `json:"method"`
	Milliseconds int64  `json:"milliseconds"`
}

type HttpMessageContent struct {
	Method string `json:"method"`
	URL    string `json:"url"`
}

func TestByHTTP(method, url string) (*HttpStatistics, error) {
	start := time.Now()

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	end := time.Since(start).Milliseconds()
	return &HttpStatistics{
		StatusCode:   int64(res.StatusCode),
		Method:       method,
		Milliseconds: end,
	}, nil
}
