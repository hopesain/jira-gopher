package issues

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/hopesain/jira-gopher/internal/jira"
)

// Get issue types for user
func GetIssueTypes(credentials jira.JiraCredentials) ([]IssueType, error) {
	client := &http.Client{
		Timeout: time.Second * 15,
	}

	url := credentials.BaseUrl + "/issuetype"

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to build the request: %w", err)
	}

	req.SetBasicAuth(credentials.Email, credentials.Token)
	req.Header.Set("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read the request body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, &jira.HttpResponseError{
			Status:     resp.Status,
			StatusCode: resp.StatusCode,
			Message:    "something went wrong",
			Body:       body,
		}
	}

	var response []IssueType

	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("failed to decode the response body: %w", err)
	}

	return response, nil
}
