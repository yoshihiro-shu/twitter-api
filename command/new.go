package command

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	get_user_follower "github.com/yoshihiro-shu/twitter/command/twiiter-api-v2/get-user-followers/commands"
	get_user_liked_tweets "github.com/yoshihiro-shu/twitter/command/twiiter-api-v2/get-user-liked-tweets/commands"
	get_user_timeline "github.com/yoshihiro-shu/twitter/command/twiiter-api-v2/get-user-timeline/commands"
	twitter_service "github.com/yoshihiro-shu/twitter/twitter/service"
)

func NewApplicationCommand(t *twitter_service.TwitterAPIHandler) *cobra.Command {
	var command = &cobra.Command{
		Use:   "app",
		Short: "This commands is available with your laptop!",
		Example: `
by default you can run "go run main.go"

if you want play with twitter api, you can run "go run main.go  get-timeline " `,
		Run: func(c *cobra.Command, args []string) {
			// c.HelpFunc()(c, args)
			fmt.Println("Hello !!! Now You Can Enjoy with this commands!!!")
			os.Exit(1)
		},
	}

	command.AddCommand(get_user_timeline.NewApplicationGetTwiiterTimeline(t))
	command.AddCommand(get_user_follower.NewApplicationGetUserFollowers(t))
	command.AddCommand(get_user_liked_tweets.NewApplicationGetUserLikedTweets(t))

	return command
}

func Execute(cmd *cobra.Command) {
	if err := cmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
