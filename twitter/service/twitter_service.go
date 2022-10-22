package twitter_service

import (
	user_follwer "github.com/yoshihiro-shu/twitter/twitter/user/follower"
	user_like_tweet "github.com/yoshihiro-shu/twitter/twitter/user/like/tweet"
	user_tweet_timeline "github.com/yoshihiro-shu/twitter/twitter/user/tweet/timeline"
)

func (t TwitterAPIHandler) GetUserFollowers() (string, error) {
	return user_follwer.Do(t.userId, t.getBearerToken())
}

func (t TwitterAPIHandler) GetUserTweetTimeline() (string, error) {
	return user_tweet_timeline.Do(t.userId, t.getBearerToken())
}

func (t TwitterAPIHandler) GetUserLikeTweets() (string, error) {
	return user_like_tweet.Do(t.userId, t.getBearerToken())
}
