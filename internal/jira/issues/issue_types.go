package issues

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"time"

	"github.com/hopesain/gojira/internal/jira"
)

// Get Issue Types for User
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

	var response []IssueType

	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("failed to decode the response body: %w", err)
	}

	if resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusAccepted {
		slog.Info(
			"successfully fetched issue types",
			"body", string(body),
			"status", resp.Status,
			"status code", resp.StatusCode,
		)

		return response, nil

	}

	return nil, fmt.Errorf("failed to fetch issue types for user: %v", string(body))
}
