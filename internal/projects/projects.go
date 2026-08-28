package projects

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/hopesain/jira-gopher/internal"
	"github.com/hopesain/jira-gopher/internal/config"
)

type ProjectsService struct {
	credentials config.Credentials
	httpClient  *http.Client
}

func New(credentials config.Credentials, httpClient *http.Client) *ProjectsService {
	return &ProjectsService{
		credentials: credentials,
		httpClient:  httpClient,
	}
}

func (p *ProjectsService) Get() (ProjectsResponse, error) {
	url := p.credentials.BaseUrl + "/project/search"

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return ProjectsResponse{}, fmt.Errorf("failed to build the request: %w", err)
	}

	req.SetBasicAuth(p.credentials.Email, p.credentials.Token)
	req.Header.Set("Accept", "application/json")

	resp, err := p.httpClient.Do(req)
	if err != nil {
		return ProjectsResponse{}, fmt.Errorf("request failed: %w", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return ProjectsResponse{}, fmt.Errorf("failed to read the response body: %w", err)
	}

	if resp.StatusCode == http.StatusBadRequest {
		return ProjectsResponse{}, &internal.HttpResponseError{
			Status:     resp.Status,
			StatusCode: resp.StatusCode,
			Message:    "invalid request",
			Body:       body,
		}
	}

	if resp.StatusCode == http.StatusUnauthorized {
		return ProjectsResponse{}, &internal.HttpResponseError{
			Status:     resp.Status,
			StatusCode: resp.StatusCode,
			Message:    "incorrect or missing authentication credentials",
			Body:       body,
		}
	}

	if resp.StatusCode == http.StatusNotFound {
		return ProjectsResponse{}, &internal.HttpResponseError{
			Status:     resp.Status,
			StatusCode: resp.StatusCode,
			Message:    "no projects matching the search criteria are found",
			Body:       body,
		}
	}

	if resp.StatusCode != http.StatusOK {
		return ProjectsResponse{}, &internal.HttpResponseError{
			Status:     resp.Status,
			StatusCode: resp.StatusCode,
			Message:    "something went wrong",
			Body:       body,
		}
	}

	var response ProjectsResponse

	if err := json.Unmarshal(body, &response); err != nil {
		return ProjectsResponse{}, fmt.Errorf("failed to decode the response body: %w", err)
	}

	return response, nil

}

func (p *ProjectsService) Create(payload CreateProjectRequest) (CreateProjectResponse, error) {
	if err := payload.Validate(); err != nil {
		return CreateProjectResponse{}, fmt.Errorf("validation error: %w", err)
	}

	url := p.credentials.BaseUrl + "/project"

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return CreateProjectResponse{}, fmt.Errorf("failed to marshal or serialize the payload: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(jsonPayload))
	if err != nil {
		return CreateProjectResponse{}, fmt.Errorf("failed to build the request: %w", err)
	}

	req.SetBasicAuth(p.credentials.Email, p.credentials.Token)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	resp, err := p.httpClient.Do(req)
	if err != nil {
		return CreateProjectResponse{}, fmt.Errorf("request failed: %w", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return CreateProjectResponse{}, fmt.Errorf("failed to read the response body: %w", err)
	}

	if resp.StatusCode == http.StatusBadRequest {
		return CreateProjectResponse{}, &internal.HttpResponseError{
			Status:     resp.Status,
			StatusCode: resp.StatusCode,
			Message:    "invalid request",
			Body:       body,
		}
	}

	if resp.StatusCode == http.StatusUnauthorized {
		return CreateProjectResponse{}, &internal.HttpResponseError{
			Status:     resp.Status,
			StatusCode: resp.StatusCode,
			Message:    "incorrect or missing authentication credentials",
			Body:       body,
		}
	}

	if resp.StatusCode == http.StatusForbidden {
		return CreateProjectResponse{}, &internal.HttpResponseError{
			Status:     resp.Status,
			StatusCode: resp.StatusCode,
			Message:    "user does not have permission to create a project",
			Body:       body,
		}
	}

	if resp.StatusCode != http.StatusCreated {
		return CreateProjectResponse{}, &internal.HttpResponseError{
			Status:     resp.Status,
			StatusCode: resp.StatusCode,
			Message:    "something went wrong",
			Body:       body,
		}
	}

	var response CreateProjectResponse

	if err := json.Unmarshal(body, &response); err != nil {
		return CreateProjectResponse{}, fmt.Errorf("failed to decode the response body")
	}

	return response, nil

}

func CreateProject(credentials config.Credentials, payload CreateProjectRequest) (CreateProjectResponse, error) {
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
		return CreateProjectResponse{}, &internal.HttpResponseError{
			Status:     resp.Status,
			StatusCode: resp.StatusCode,
			Message:    "invalid request",
			Body:       body,
		}
	}

	if resp.StatusCode == http.StatusUnauthorized {
		return CreateProjectResponse{}, &internal.HttpResponseError{
			Status:     resp.Status,
			StatusCode: resp.StatusCode,
			Message:    "incorrect or missing authentication credentials",
			Body:       body,
		}
	}

	if resp.StatusCode == http.StatusForbidden {
		return CreateProjectResponse{}, &internal.HttpResponseError{
			Status:     resp.Status,
			StatusCode: resp.StatusCode,
			Message:    "user does not have permission to create a project",
			Body:       body,
		}
	}

	if resp.StatusCode != http.StatusCreated {
		return CreateProjectResponse{}, &internal.HttpResponseError{
			Status:     resp.Status,
			StatusCode: resp.StatusCode,
			Message:    "something went wrong",
			Body:       body,
		}
	}

	var response CreateProjectResponse

	if err := json.Unmarshal(body, &response); err != nil {
		return CreateProjectResponse{}, fmt.Errorf("failed to decode the response body")
	}

	return response, nil

}
