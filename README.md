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

## Sample Command

```golang
# endpoint is "https://api.twitter.com/2/users/:id/tweets" to get timeline by userId recently
go run main.go  get-timeline
```

```
% go run main.go  --help
This commands is available with your laptop!

Usage:
  app [flags]
  app [command]

Examples:

by default you can run "go run main.go"

if you want play with twitter api, you can run "go run main.go  get-timeline " 

Available Commands:
  completion            Generate the autocompletion script for the specified shell
  get-follwers          get followers by user in Twitter
  get-timeline          get timeline by user recently in Twitter
  get-user-liked-tweets get tweets by user liked in Twitter
  help                  Help about any command

Flags:
  -h, --help   help for app

Use "app [command] --help" for more information about a command.
```

## References
https://developer.twitter.com/en/portal/dashboard

https://developer.twitter.com/en/docs/twitter-api

https://github.com/twitterdev/Twitter-API-v2-sample-code