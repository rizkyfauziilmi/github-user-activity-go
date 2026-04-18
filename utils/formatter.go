package utils

import (
	"fmt"
	"strings"

	"github.com/rizkyfauziilmi/github-user-activity-go/models"
)

func DisplayActivity(events []models.Event) {
	if len(events) == 0 {
		fmt.Println("No recent activity found.")
		return
	}

	for _, event := range events {
		var action string

		switch event.Type {
		case "PushEvent":
			commitCount := len(event.Payload.Commits)
			action = fmt.Sprintf("Pushed %d commit(s) to %s", commitCount, event.Repo.Name)

		case "IssuesEvent":
			// Capitalize first letter of action
			rawAction := event.Payload.Action
			formattedAction := strings.ToUpper(rawAction[:1]) + rawAction[1:]
			action = fmt.Sprintf("%s an issue in %s", formattedAction, event.Repo.Name)

		case "WatchEvent":
			action = fmt.Sprintf("Starred %s", event.Repo.Name)

		case "ForkEvent":
			action = fmt.Sprintf("Forked %s", event.Repo.Name)

		case "CreateEvent":
			action = fmt.Sprintf("Created %s in %s", event.Payload.RefType, event.Repo.Name)

		default:
			typeName := strings.TrimSuffix(event.Type, "Event")
			action = fmt.Sprintf("%s in %s", typeName, event.Repo.Name)
		}

		fmt.Printf("- %s\n", action)
	}
}
