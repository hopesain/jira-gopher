package projects

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"time"

	"github.com/hopesain/gojira/internal/jira"
)

const (
	MAX_CHARACTER_LENGTH = 10
)

type CreateProjectRequest struct {
	AssigneeType       string `json:"assigneeType"`
	Description        string `json:"description"`
	Key                string `json:"key"`
	LeadAccountId      string `json:"leadAccountId"`
	Name               string `json:"name"`
	ProjectTemplateKey string `json:"projectTemplateKey"`
	ProjectTypeKey     string `json:"projectTypeKey"`
}

func (c *CreateProjectRequest) Validate() error {
	if c.AssigneeType != "UNASSIGNED" && c.AssigneeType != "PROJECT_LEAD" {
		return fmt.Errorf("assignee type must either be UNASSIGNED or PROJECT_LEAD")
	}

	if len(c.Key) > MAX_CHARACTER_LENGTH {
		return fmt.Errorf("the key exceeds maximum length: %v", c.Key)
	}

	var missing []string

	if c.AssigneeType == "" {
		missing = append(missing, "assign type")
	}

	if c.Description == "" {
		missing = append(missing, "description")
	}

	if c.Key == "" {
		missing = append(missing, "key")
	}

	if c.LeadAccountId == "" {
		missing = append(missing, "lead account id")
	}

	if c.Name == "" {
		missing = append(missing, "name")
	}

	if c.ProjectTemplateKey == "" {
		missing = append(missing, "project template key")
	}

	if c.ProjectTypeKey == "" {
		missing = append(missing, "project type key")
	}

	if len(missing) != 0 {
		return fmt.Errorf("missing required fields for creating a project in Jira: %v", missing)
	}

	return nil
}

type CreateProjectResponse struct {
}

func CreateProject(credentials jira.JiraCredentials, payload CreateProjectRequest) error {
	if err := payload.Validate(); err != nil {
		return fmt.Errorf("validation error: %w", err)
	}

	client := &http.Client{
		Timeout: time.Second * 15,
	}

	url := credentials.BaseUrl + "/project"

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal or serialize the payload: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(jsonPayload))
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

	if resp.StatusCode == http.StatusBadRequest {
		return fmt.Errorf("invalid request: %s", string(body))
	}

	if resp.StatusCode == http.StatusUnauthorized {
		return fmt.Errorf("incorrect or missing authentication credentials")
	}

	if resp.StatusCode == http.StatusForbidden {
		return fmt.Errorf("you do not have permission to create a project")
	}

	if resp.StatusCode == http.StatusConflict {
		return fmt.Errorf("a project with this key or name already exists")
	}

	slog.Info("jira request completed",
		"path", url,
		"status", resp.Status,
		"status_code", resp.StatusCode,
		"body", string(body),
	)

	return nil
}

func GetProjects(credentials jira.JiraCredentials) error {
	client := &http.Client{
		Timeout: time.Second * 15,
	}

	url := credentials.BaseUrl + "/project/search"

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("failed to build the request: %w", err)
	}

	req.SetBasicAuth(credentials.Email, credentials.Token)
	req.Header.Set("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read the response body: %w", err)
	}

	if resp.StatusCode == http.StatusBadRequest {
		return fmt.Errorf("invalid request")
	}

	if resp.StatusCode == http.StatusUnauthorized {
		return fmt.Errorf("incorrect or missing authentication credentials")
	}

	if resp.StatusCode == http.StatusNotFound {
		return fmt.Errorf("no projects matching the search criteria are found")
	}

	slog.Info("jira request completed",
		"path", url,
		"status", resp.Status,
		"status_code", resp.StatusCode,
		"body", string(body),
	)

	return nil

}
