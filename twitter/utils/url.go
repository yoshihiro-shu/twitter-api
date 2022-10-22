package utils

import (
	"fmt"
	"net/url"
)

func CreateUrl(urlFormat string, args ...interface{}) (*url.URL, error) {
	endpoint := fmt.Sprintf(urlFormat, args...)

	url, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	return url, nil
}

func SetParams(url *url.URL, params map[string]string) {
	q := url.Query()
	for i, v := range params {
		q.Set(i, v)
	}

	url.RawQuery = q.Encode()
}
