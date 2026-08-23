package projects

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/hopesain/gojira/internal/jira"
)

func GetProjects(credentials jira.JiraCredentials) (ProjectsResponse, error) {
	client := &http.Client{
		Timeout: time.Second * 15,
	}

	url := credentials.BaseUrl + "/project/search"

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return ProjectsResponse{}, fmt.Errorf("failed to build the request: %w", err)
	}

	req.SetBasicAuth(credentials.Email, credentials.Token)
	req.Header.Set("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return ProjectsResponse{}, fmt.Errorf("request failed: %w", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return ProjectsResponse{}, fmt.Errorf("failed to read the response body: %w", err)
	}

	if resp.StatusCode == http.StatusBadRequest {
		return ProjectsResponse{}, fmt.Errorf("invalid request: %v", string(body))
	}

	if resp.StatusCode == http.StatusUnauthorized {
		return ProjectsResponse{}, fmt.Errorf("incorrect or missing authentication credentials: %v", string(body))
	}

	if resp.StatusCode == http.StatusNotFound {
		return ProjectsResponse{}, fmt.Errorf("no projects matching the search criteria are found: %v", string(body))
	}

	var response ProjectsResponse

	if err := json.Unmarshal(body, &response); err != nil {
		return ProjectsResponse{}, fmt.Errorf("failed to decode the response body: %w", err)
	}

	return response, nil

	

}

func CreateProject(credentials jira.JiraCredentials, payload CreateProjectRequest) (CreateProjectResponse, error) {
	if err := payload.Validate(); err != nil {
		return CreateProjectResponse{}, fmt.Errorf("validation error: %w", err)
	}

	client := &http.Client{
		Timeout: time.Second * 15,
	}

	url := credentials.BaseUrl + "/project"

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return CreateProjectResponse{}, fmt.Errorf("failed to marshal or serialize the payload: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(jsonPayload))
	if err != nil {
		return CreateProjectResponse{}, fmt.Errorf("failed to build the request: %w", err)
	}

	req.SetBasicAuth(credentials.Email, credentials.Token)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return CreateProjectResponse{}, fmt.Errorf("request failed: %w", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return CreateProjectResponse{}, fmt.Errorf("failed to read the response body: %w", err)
	}

	if resp.StatusCode == http.StatusBadRequest {
		return CreateProjectResponse{}, fmt.Errorf("invalid request: %v", string(body))
	}

	if resp.StatusCode == http.StatusUnauthorized {
		return CreateProjectResponse{}, fmt.Errorf("incorrect or missing authentication credentials: %v", string(body))
	}

	if resp.StatusCode == http.StatusForbidden {
		return CreateProjectResponse{}, fmt.Errorf("you do not have permission to create a project: %v", string(body))
	}

	var response CreateProjectResponse

	if err := json.Unmarshal(body, &response); err != nil {
		return CreateProjectResponse{}, fmt.Errorf("failed to decode the response body")
	}

	return response, nil

}
