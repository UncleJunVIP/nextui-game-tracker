package cmd

import (
	"fmt"
	"log"
	"nextui-game-tracker/database"

	"github.com/spf13/cobra"
)

var gamePath string

func init() {
	startCommand.Flags().StringVarP(&gamePath, "game", "g", "",
		"path of the game's ROM file to start a play session")
	startCommand.MarkFlagRequired("game")
	rootCmd.AddCommand(startCommand)
}

var startCommand = &cobra.Command{
	Use:   "start",
	Short: "Start a new game session, forcibly closing any existing sessions",
	Run: func(cmd *cobra.Command, args []string) {
		err := database.StartSession(gamePath)
		if err != nil {
			log.Printf("Unable to start session for \"%s\": %s\n", gamePath, err.Error())
		} else {
			log.Println(fmt.Sprintf("Started a new game session for \"%s\"", gamePath))
		}
	},
}
