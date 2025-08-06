package cmd

import (
	"fmt"
	"log"
	"nextui-game-tracker/database"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(endCommand)
}

var endCommand = &cobra.Command{
	Use:   "end",
	Short: "Ends the current game session(s)",
	Run: func(cmd *cobra.Command, args []string) {
		sessions, err := database.EndSession()
		if err != nil {
			log.Println("Unable to end sessions")
		} else {
			if len(sessions) == 1 {
				log.Println(fmt.Sprintf("Ended session for Game ID: \"%d\"", sessions[0].GameID.Int64))
			} else if len(sessions) > 1 {
				log.Println("Ended sessions for multiple games...")
			} else {
				log.Println("No sessions to end")
			}
		}
	},
}
