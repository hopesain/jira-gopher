package issues

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"time"

	"github.com/hopesain/jira-gopher/internal/jira"
)

func CreateIssue(credentials jira.JiraCredentials, payload CreateIssueRequest) error {
	client := &http.Client{
		Timeout: time.Second * 15,
	}

	url := credentials.BaseUrl + "/issue"

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal or serialize the payload: %w", err)
	}

	req, err := http.NewRequest(
		http.MethodPost,
		url,
		bytes.NewReader(jsonPayload),
	)
	if err != nil {
		return fmt.Errorf("failed to build the request: %w", err)
	}

	req.SetBasicAuth(credentials.Email, credentials.Token)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read the response body: %w", err)
	}

	slog.Info(
		"create issue response",
		"status", resp.Status,
		"status code", resp.StatusCode,
		"body", string(body),
	)

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf(
			"failed to create issue: %s",
			resp.Status,
		)
	}

	return nil
}


