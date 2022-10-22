package twitter_service

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	user_tweet_timeline "github.com/yoshihiro-shu/twitter/twitter/user/tweet/timeline"
)

type TwitterAPIHandler struct {
	UserId       string `json:"userId"`
	UserName     string `json:"userName"`
	APIKey       string `json:"consumer_key"`
	APIKeySecret string `json:"consumer_secret"`
	BearerToken  string `json:"bearer_token"`
}

func New() *TwitterAPIHandler {
	return &TwitterAPIHandler{
		UserId:       os.Getenv("UserId"),
		UserName:     os.Getenv("UserName"),
		APIKey:       os.Getenv("APIKey"),
		APIKeySecret: os.Getenv("APIKeySecret"),
		BearerToken:  os.Getenv("BearerToken"),
	}
}

func (t TwitterAPIHandler) getBearerToken() string {
	return fmt.Sprintf("Bearer %s", t.BearerToken)
}

func (t TwitterAPIHandler) GetRequestApi() (string, error) {
	url := "https://api.twitter.com/2/tweets/search/recent?query=from:twitterdev"
	clinet := &http.Client{
		Timeout: time.Second * 10,
	}

	req, _ := http.NewRequest(http.MethodGet, url, nil)

	req.Header.Add("Authorization", t.getBearerToken())

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

func (t TwitterAPIHandler) GetMyTweets() (string, error) {
	url := "https://api.twitter.com/2/users/%s/tweets"

	clinet := &http.Client{
		Timeout: time.Second * 10,
	}

	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf(url, t.UserId), nil)

	req.Header.Add("Authorization", t.getBearerToken())

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

func (t TwitterAPIHandler) GetUserTweetTimeline() (string, error) {
	return user_tweet_timeline.Do(t.UserId, t.BearerToken)
}
