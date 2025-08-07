package cmd

import (
	"github.com/spf13/cobra"
	"nextui-game-tracker/database"
)

var (
	isDev bool

	rootCmd = &cobra.Command{
		Use:   "game-tracker",
		Short: "A CLI tool for tracking NextUI game sessions.",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			database.InitializeDB(isDev)
		},
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().
		BoolVarP(&isDev, "dev", "d", false,
			"enable development mode")
}
