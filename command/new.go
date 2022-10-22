package command

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	get_user_timeline "github.com/yoshihiro-shu/twitter/command/twiiter-api-v2/get-user-timeline/commands"
	twitter_service "github.com/yoshihiro-shu/twitter/twitter/service"
)

func NewApplicationCommand(t *twitter_service.TwitterAPIHandler) *cobra.Command {
	// this command doesn't work
	var command = &cobra.Command{
		Use:   "app",
		Short: "Manage applications",
		Example: `Cobra is a CLI library for Go that empowers applications.
		This application is a tool to generate the needed files
		to quickly create a Cobra application.1`,
		Run: func(c *cobra.Command, args []string) {
			// c.HelpFunc()(c, args)
			fmt.Println("APP")
			os.Exit(1)
		},
	}

	command.AddCommand(get_user_timeline.NewApplicationGetTwiiterTimeline(t))

	return command
}

func Execute(cmd *cobra.Command) {
	if err := cmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
