package twitter_service

import (
	"fmt"
	"os"
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
