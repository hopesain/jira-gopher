package main

import (
	"io"
	"log/slog"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stderr, nil))
	slog.SetDefault(logger)

	if err := godotenv.Load(); err != nil {
		slog.Error("failed to load .env file", "error", err)
		os.Exit(1)
	}

	baseURL := os.Getenv("JIRA_BASE_URL") // e.g. https://your-domain.atlassian.net
	email := os.Getenv("JIRA_EMAIL")
	token := os.Getenv("JIRA_API_TOKEN")

	if baseURL == "" || email == "" || token == "" {
		slog.Error("missing required env vars", "need", "JIRA_BASE_URL, JIRA_EMAIL, JIRA_API_TOKEN")
		os.Exit(1)
	}

	callJira(baseURL, email, token, "/rest/api/3/myself")
	callJira(baseURL, email, token, "/rest/api/3/project/search")
	callJira(baseURL, email, token, "/rest/api/3/serverInfo")
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