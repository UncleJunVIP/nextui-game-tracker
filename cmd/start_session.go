package cmd

import (
	"fmt"
	"log"
	"nextui-game-tracker/database"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(startCommand)
}

var startCommand = &cobra.Command{
	Use:   "start <game-path>",
	Short: "Start a new game session, forcibly closing any existing sessions",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		gamePath := args[0]
		err := database.StartSession(gamePath)
		if err != nil {
			log.Printf("Unable to start session for \"%s\": %s\n", gamePath, err.Error())
		} else {
			log.Println(fmt.Sprintf("Started a new game session for \"%s\"", gamePath))
		}
	},
}
