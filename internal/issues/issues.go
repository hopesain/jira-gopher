package issues

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/hopesain/jira-gopher/internal"
	"github.com/hopesain/jira-gopher/internal/config"
)

type IssuesService struct {
	credentials config.Credentials
	httpClient  *http.Client
}

func New(credentials config.Credentials, httpClient *http.Client) *IssuesService {
	return &IssuesService{
		credentials: credentials,
		httpClient:  httpClient,
	}
}

func (i *IssuesService) Create(payload CreateIssueRequest) (CreateIssueResponse, error) {
	url := i.credentials.BaseUrl + "/issue"

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return CreateIssueResponse{}, fmt.Errorf("failed to marshal or serialize the payload: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(jsonPayload))
	if err != nil {
		return CreateIssueResponse{}, fmt.Errorf("failed to build the request: %w", err)
	}

	req.SetBasicAuth(i.credentials.Email, i.credentials.Token)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	resp, err := i.httpClient.Do(req)
	if err != nil {
		return CreateIssueResponse{}, fmt.Errorf("request failed: %w", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return CreateIssueResponse{}, fmt.Errorf("failed to read the response body: %w", err)
	}

	if resp.StatusCode == http.StatusBadRequest {
		return CreateIssueResponse{}, &internal.HttpResponseError{
			Status:     resp.Status,
			StatusCode: resp.StatusCode,
			Message:    "request failed",
			Body:       body,
		}
	}

	if resp.StatusCode == http.StatusUnauthorized {
		return CreateIssueResponse{}, &internal.HttpResponseError{
			Status:     resp.Status,
			StatusCode: resp.StatusCode,
			Message:    "incorrect or missing authentication credentials",
			Body:       body,
		}
	}

	if resp.StatusCode == http.StatusForbidden {
		return CreateIssueResponse{}, &internal.HttpResponseError{
			Status:     resp.Status,
			StatusCode: resp.StatusCode,
			Message:    "user does not have necessary permissions to create a task",
			Body:       body,
		}
	}

	if resp.StatusCode == http.StatusUnprocessableEntity {
		return CreateIssueResponse{}, &internal.HttpResponseError{
			Status:     resp.Status,
			StatusCode: resp.StatusCode,
			Message:    "configuration problems",
			Body:       body,
		}
	}

	if resp.StatusCode != http.StatusCreated {
		return CreateIssueResponse{}, &internal.HttpResponseError{
			Status:     resp.Status,
			StatusCode: resp.StatusCode,
			Message:    "something went wrong",
			Body:       body,
		}
	}

	var response CreateIssueResponse

	if err := json.Unmarshal(body, &response); err != nil {
		return CreateIssueResponse{}, fmt.Errorf("failed to decode the response body: %w", err)
	}

	return response, nil
}
