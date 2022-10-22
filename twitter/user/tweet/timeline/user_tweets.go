package user_tweet_timeline

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

var tweetParams = map[string]string{"tweet.fields": "created_at"}

func Do(userId, bearerToken string) (string, error) {
	url := createUrl(userId)

	data, err := connectToEndPoint(url.String(), bearerToken)
	if err != nil {
		return "", err
	}
	return data, nil
}

func createUrl(userId string) *url.URL {
	endpoint := fmt.Sprintf("https://api.twitter.com/2/users/%s/tweets", userId)

	url, _ := url.Parse(endpoint)
	setParams(url)

	return url
}

func setParams(u *url.URL) {
	q := u.Query()
	for i, v := range tweetParams {
		q.Set(i, v)
	}

	u.RawQuery = q.Encode()
}

func setBearerOauth(r *http.Request, bearerToken string) {
	r.Header.Set("Authorization", fmt.Sprintf("Bearer %s", bearerToken))
}

func connectToEndPoint(url, bearerToken string) (string, error) {
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
