package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/rizkyfauziilmi/github-user-activity-go/utils"
	"github.com/rizkyfauziilmi/github-user-activity-go/validator"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("Usage: <username>")
		return
	}

	username := strings.Join(args[1:], " ")

	if err := validator.ValidateGitHubUsername(username); err != nil {
		fmt.Printf("Username not valid: %s\n", err)
		return
	}

	events, err := utils.FetchGitHubActivity(username)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	utils.DisplayActivity(events)
}
