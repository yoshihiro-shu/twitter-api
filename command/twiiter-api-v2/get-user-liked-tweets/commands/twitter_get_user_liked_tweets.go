package commands

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	twitter_service "github.com/yoshihiro-shu/twitter/twitter/service"
)

const cliName = "get-user-liked-tweets"

func NewApplicationGetUserLikedTweets(t *twitter_service.TwitterAPIHandler) *cobra.Command {
	var command = cobra.Command{
		Use:   cliName,
		Short: "get twitter timeline by user",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			data, err := t.GetUserLikeTweets()
			if err != nil {
				log.Fatalln(err)
			}

			fmt.Println(data)
		},
	}
	return &command
}
