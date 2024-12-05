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

func TestByHTTP(method, url string) (*HttpStatistics, error) {
	initialTime := time.Now().UnixMilli()
	res, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	finalTime := initialTime - time.Now().UnixMilli()
	return &HttpStatistics{
		StatusCode:   int64(res.Response.StatusCode),
		Method:       res.Method,
		Milliseconds: finalTime,
	}, nil
}
