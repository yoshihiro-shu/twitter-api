package user_like_tweet

import "github.com/yoshihiro-shu/twitter/twitter/utils"

const apiEndpoint = "/2/users/%s/liked_tweets"

func Do(userId, bearerToken string) (string, error) {
	url, _ := utils.CreateUrlApiV2(apiEndpoint, userId)

	data, err := utils.ConnectToEndpointHttpGet(url.String(), bearerToken)
	if err != nil {
		return "", err
	}
	return data, nil
}
