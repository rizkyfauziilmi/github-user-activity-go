package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rizkyfauziilmi/github-user-activity-go/models"
)

func FetchGitHubActivity(username string) ([]models.Event, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s/events", username)

	req, _ := http.NewRequest("GET", url, nil)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == 404 {
			return nil, fmt.Errorf("user not found. Please check the username")
		}
		return nil, fmt.Errorf("error fetching data: %d", resp.StatusCode)
	}

	var events []models.Event
	if err := json.NewDecoder(resp.Body).Decode(&events); err != nil {
		return nil, err
	}

	return events, nil
}
