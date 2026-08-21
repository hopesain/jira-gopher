package projects

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"time"

	"github.com/hopesain/gojira/internal/jira"
)

func GetAllProjectTypes(credentials jira.JiraCredentials) error {
	client := &http.Client{
		Timeout: time.Second * 15,
	}

	url := credentials.BaseUrl + "/project/type"

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("failed to build the request: %w", err)
	}

	req.SetBasicAuth(credentials.Email, credentials.Token)
	req.Header.Set("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to get all project types: %w", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read the response body: %w", err)
	}

	if resp.StatusCode == http.StatusUnauthorized {
		return fmt.Errorf("incorrect authentication credentials")
	}

	slog.Info("jira request completed",
		"path", url,
		"status", resp.Status,
		"status_code", resp.StatusCode,
		"body", string(body),
	)

	return nil

}
