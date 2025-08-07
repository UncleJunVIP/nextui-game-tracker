package cmd

import (
	"log"
	"nextui-game-tracker/database"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(resumeCommand)
}

var resumeCommand = &cobra.Command{
	Use:   "resume",
	Short: "Finds the rom associated with last closed session and starts a new session for it",
	Run: func(cmd *cobra.Command, args []string) {
		resumedID, err := database.ResumeSession()
		if err != nil {
			log.Println("Unable to resume session!")
		}
		if resumedID == -1 {
			log.Println("No game session to resume!")
		} else {
			log.Printf("Resumed session for Game ID: \"%d\"", resumedID)
		}
	},
}
