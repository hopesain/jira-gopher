package myself

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/hopesain/gojira/internal/jira"
)

func GetCurrentUser(credentials jira.JiraCredentials) (MyselfResponse, error) {
	client := &http.Client{
		Timeout: time.Second * 15,
	}

	url := credentials.BaseUrl + "/myself"

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return MyselfResponse{}, fmt.Errorf("failed to build the request: %w", err)
	}

	req.SetBasicAuth(credentials.Email, credentials.Token)
	req.Header.Set("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return MyselfResponse{}, fmt.Errorf("request failed: %w", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(req.Body)
	if err != nil {
		return MyselfResponse{}, fmt.Errorf("failed to read the response body")
	}

	if resp.StatusCode == http.StatusUnauthorized {
		return MyselfResponse{}, fmt.Errorf("incorrect authorization credentials. check your email and token: %v", string(body))
	}

	var response MyselfResponse

	json.NewDecoder(resp.Body).Decode(&response)

	return response, nil

}

func UserAccountID(credentials jira.JiraCredentials) (accountID string, err error) {
	userInformation, err := GetCurrentUser(credentials)
	if err != nil {
		return "", fmt.Errorf("failed to retrieve the user information: %w", err)
	}

	accountID = userInformation.AccountID

	return accountID, nil
}
