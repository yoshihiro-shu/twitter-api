package utils

import (
	"io"
	"net/http"
	"time"
)

// request client for http method is GET
func ConnectToEndpointHttpGet(url, bearerToken string) (string, error) {
	clinet := &http.Client{
		Timeout: time.Second * 10,
	}
	req, _ := http.NewRequest(http.MethodGet, url, nil)

	setBearerOauth(req, bearerToken)

	res, err := clinet.Do(req)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()
	b, err := io.ReadAll(res.Body)
	if err != nil {
		return "", nil
	}

	return string(b), nil
}
