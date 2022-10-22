# Getting Start

- Get Twiiter API Account From [References](https://developer.twitter.com/en/docs/twitter-api)

- SET the Secret Infomation in env file

```
touch .env
```

```.env
APIKey=XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
APIKeySecret=XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
BearerToken=XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
UserId=XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
````

```
# to get userid
curl --request GET 'https://api.twitter.com/2/users/by/username/USER_NAME --header 'Authorization: Bearer XXXXXX'
```

## sample

```golang
go run main.go  get-timeline
```

## References
https://developer.twitter.com/en/portal/dashboard

https://developer.twitter.com/en/docs/twitter-api

https://github.com/twitterdev/Twitter-API-v2-sample-code