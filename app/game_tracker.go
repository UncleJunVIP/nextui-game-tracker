package main

import (
	"nextui-game-tracker/cmd"
	"nextui-game-tracker/database"
)

func main() {
	defer database.CloseDatabase()
	_ = cmd.Execute()
}
