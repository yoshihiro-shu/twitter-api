package user_follwer

import "github.com/yoshihiro-shu/twitter/twitter/utils"

const endpoint = "/2/users/%s/followers"

func Do(userId, bearerToken string) (string, error) {

	url, err := utils.CreateUrlApiV2(endpoint, userId)
	if err != nil {
		return "", err
	}

	return utils.ConnectToEndpointHttpGet(url.String(), bearerToken)
}
