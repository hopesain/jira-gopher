package main

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"

	"github.com/hopesain/gojira/internal/jira"
	"github.com/hopesain/gojira/internal/jira/myself"
	"github.com/joho/godotenv"
)

var (
	JIRA_EMAIL     string
	JIRA_API_TOKEN string
	JIRA_BASE_URL  string
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stderr, nil))
	slog.SetDefault(logger)

	if err := godotenv.Load(); err != nil {
		slog.Error("failed to load .env file", "error", err)
		os.Exit(1)
	}

	JIRA_EMAIL = os.Getenv("JIRA_EMAIL")
	JIRA_API_TOKEN = os.Getenv("JIRA_API_TOKEN")
	JIRA_BASE_URL = os.Getenv("JIRA_BASE_URL")

	jiraCredentials := jira.JiraCredentials{
		Email:   JIRA_EMAIL,
		Token:   JIRA_API_TOKEN,
		BaseUrl: JIRA_BASE_URL,
	}

	if err := jiraCredentials.Validate(); err != nil {
		slog.Error("validation error", "error", err)
		os.Exit(1)
	}

	fmt.Println("All credentials available")

	if err := myself.GetCurrentUser(jiraCredentials); err != nil {
		slog.Error("failed to get current user information", "error", err)
		os.Exit(1)
	}

}

// callJira hits a single Jira REST endpoint and logs the result. It's a
// throwaway helper for smoke-testing multiple endpoints — not meant to
// survive once real architecture goes in.
func callJira(baseURL, email, token, path string) {
	req, err := http.NewRequest(http.MethodGet, baseURL+path, nil)
	if err != nil {
		slog.Error("failed to build request", "path", path, "error", err)
		return
	}

	req.SetBasicAuth(email, token)
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		slog.Error("request failed", "path", path, "error", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		slog.Error("failed to read response body", "path", path, "error", err)
		return
	}

	slog.Info("jira request completed",
		"path", path,
		"status", resp.Status,
		"status_code", resp.StatusCode,
		"body", string(body),
	)
}
