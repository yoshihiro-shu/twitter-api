package user_tweet_timeline

import (
	"github.com/yoshihiro-shu/twitter/twitter/utils"
)

const apiEndpoint = "/2/users/%s/tweets"

var tweetParams = map[string]string{"tweet.fields": "created_at"}

func Do(userId, bearerToken string) (string, error) {
	url, _ := utils.CreateUrlApiV2(apiEndpoint, userId)

	utils.SetParams(url, tweetParams)

	data, err := utils.ConnectToEndpointHttpGet(url.String(), bearerToken)
	if err != nil {
		return "", err
	}
	return data, nil
}
