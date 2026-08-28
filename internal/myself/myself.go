package myself

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/hopesain/jira-gopher/internal/config"
	"github.com/hopesain/jira-gopher/internal/jira"
)

type MyselfService struct {
	credentials config.Credentials
	httpClient  *http.Client
}

func New(creds config.Credentials, httpClient *http.Client) *MyselfService {
	return &MyselfService{
		credentials: creds,
		httpClient: httpClient,
	}
} 

func (m *MyselfService) GetCurrentUser() (GetCurrentUserResponse, error) {
	url := m.credentials.BaseUrl + "/myself"

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return GetCurrentUserResponse{}, fmt.Errorf("failed to build the request: %w", err)
	}

	req.SetBasicAuth(m.credentials.Email, m.credentials.Token)
	req.Header.Set("Accept", "application/json")

	resp, err := m.httpClient.Do(req)
	if err != nil {
		return GetCurrentUserResponse{}, fmt.Errorf("request failed: %w", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return GetCurrentUserResponse{}, fmt.Errorf("failed to read the response body")
	}

	if resp.StatusCode == http.StatusUnauthorized {
		return GetCurrentUserResponse{}, &jira.HttpResponseError{
			Status:     resp.Status,
			StatusCode: resp.StatusCode,
			Message:    "incorrect or missing authentication credentials",
			Body:       body,
		}
	}

	if resp.StatusCode != http.StatusOK {
		return GetCurrentUserResponse{}, &jira.HttpResponseError{
			Status:     resp.Status,
			StatusCode: resp.StatusCode,
			Message:    "something went wrong",
			Body:       body,
		}
	}

	var response GetCurrentUserResponse

	if err := json.Unmarshal(body, &response); err != nil {
		return GetCurrentUserResponse{}, fmt.Errorf("failed to decode the response body: %w", err)
	}

	return response, nil

}

func (m *MyselfService) UserAccountID() (accountID string, err error) {
	userInformation, err := m.GetCurrentUser()
	if err != nil {
		return "", fmt.Errorf("failed to retrieve the user information: %w", err)
	}
	accountID = userInformation.AccountID
	return accountID, nil
}
