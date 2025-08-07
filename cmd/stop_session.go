package cmd

import (
	"fmt"
	"log"
	"nextui-game-tracker/database"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(stopCommand)
	rootCmd.AddCommand(stopAllCommand)
}

var stopFunc = func(cmd *cobra.Command, args []string) {
	sessions, err := database.StopSession()
	if err != nil {
		log.Println("Unable to stop session(s)")
	} else {
		if len(sessions) == 1 {
			log.Println(fmt.Sprintf("Stopped session for Game ID: \"%d\"", sessions[0].GameID.Int64))
		} else if len(sessions) > 1 {
			log.Println("Stopped sessions for multiple games...")
		} else {
			log.Println("No sessions to stop")
		}
	}
}

var stopCommand = &cobra.Command{
	Use:   "stop",
	Short: "Ends the current game session(s)",
	Run:   stopFunc,
}

var stopAllCommand = &cobra.Command{
	Use:   "stop_all",
	Short: "Ends all game sessions (kept for capability with existing gametimectl)",
	Run:   stopFunc,
}
