package commands

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	twitter_service "github.com/yoshihiro-shu/twitter/twitter/service"
)

const cliName = "get-follwers"

func NewApplicationGetUserFollowers(t *twitter_service.TwitterAPIHandler) *cobra.Command {
	var command = cobra.Command{
		Use:   cliName,
		Short: "get followers by user in Twitter",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			data, err := t.GetUserFollowers()
			if err != nil {
				log.Fatalln(err)
			}

			fmt.Println(data)
		},
	}
	return &command
}
