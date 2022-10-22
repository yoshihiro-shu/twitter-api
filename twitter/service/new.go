package twitter_service

import (
	"fmt"
	"os"

	user_follwer "github.com/yoshihiro-shu/twitter/twitter/user/follower"
	user_like_tweet "github.com/yoshihiro-shu/twitter/twitter/user/like/tweet"
	user_tweet_timeline "github.com/yoshihiro-shu/twitter/twitter/user/tweet/timeline"
)

type TwitterAPIHandler struct {
	userId       string
	userName     string
	aPIKey       string
	aPIKeySecret string
	bearerToken  string
}

func New() *TwitterAPIHandler {
	return &TwitterAPIHandler{
		userId:       os.Getenv("UserId"),
		userName:     os.Getenv("UserName"),
		aPIKey:       os.Getenv("APIKey"),
		aPIKeySecret: os.Getenv("APIKeySecret"),
		bearerToken:  os.Getenv("BearerToken"),
	}
}

func (t TwitterAPIHandler) getBearerToken() string {
	return fmt.Sprintf("Bearer %s", t.bearerToken)
}

func (t TwitterAPIHandler) GetUserFollowers() (string, error) {
	return user_follwer.Do(t.userId, t.getBearerToken())
}

func (t TwitterAPIHandler) GetUserTweetTimeline() (string, error) {
	return user_tweet_timeline.Do(t.userId, t.getBearerToken())
}

func (t TwitterAPIHandler) GetUserLikeTweets() (string, error) {
	return user_like_tweet.Do(t.userId, t.getBearerToken())
}
