# Getting Start

- Get Twiiter API Account From References

- SET the Secret Infomation in env file

```
touch .env
```

```.env
APIKey=XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
APIKeySecret=XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
BearerToken=XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
````

## sample

```golang
go run main.go  get-timeline
```

```
# to get userid
curl --request GET 'https://api.twitter.com/2/users/by/username/USER_NAME --header 'Authorization: Bearer XXXXXX'
```

## References
https://developer.twitter.com/en/portal/dashboard

https://developer.twitter.com/en/docs/twitter-api
