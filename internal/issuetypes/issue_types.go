package issuetypes

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/hopesain/jira-gopher/internal"
	"github.com/hopesain/jira-gopher/internal/config"
)

type IssuesTypesService struct {
	credentials config.Credentials
	httpClient  *http.Client
}

func New(credentials config.Credentials, httpClient *http.Client) *IssuesTypesService {
	return &IssuesTypesService{
		credentials: credentials,
		httpClient:  httpClient,
	}
}

func (i *IssuesTypesService) Get() ([]IssueType, error) {
	url := i.credentials.BaseUrl + "/issuetype"

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to build the request: %w", err)
	}

	req.SetBasicAuth(i.credentials.Email, i.credentials.Token)
	req.Header.Set("Accept", "application/json")

	resp, err := i.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read the request body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, &internal.HttpResponseError{
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
