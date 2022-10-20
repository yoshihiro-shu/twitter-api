package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

const (
	url = "https://api.twitter.com/2/tweets/search/recent?query=from:twitterdev"
)

type TwitterAPIHandler struct {
	APIKey       string `json:"consumer_key"`
	APIKeySecret string `json:"consumer_secret"`
	BearerToken  string `json:"bearer_token"`
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apiKey := os.Getenv("APIKey")
	apiKeySecret := os.Getenv("APIKeySecret")
	bearerToken := os.Getenv("BearerToken")

	t := NewTwitterAPIHadler(apiKey, apiKeySecret, bearerToken)

	data, err := t.getRequestApi(url)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(data)

}

func NewTwitterAPIHadler(apikey, apiKeySecret, beareToken string) *TwitterAPIHandler {
	return &TwitterAPIHandler{
		APIKey:       apikey,
		APIKeySecret: apiKeySecret,
		BearerToken:  beareToken,
	}
}

func (t TwitterAPIHandler) getBearerToken() string {
	return fmt.Sprintf("Bearer %s", t.BearerToken)
}

func (t TwitterAPIHandler) getRequestApi(url string) (string, error) {
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
