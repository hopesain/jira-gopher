package issues

import (
	"fmt"
	"net/http"

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

func (i *IssuesService) Create(issue string) (string, error) {
	if issue == "" {
		return "", fmt.Errorf("issue name is required")
	}

	if i.credentials.Email == "" {
		return "", fmt.Errorf("it is required")
	}

	return issue + i.credentials.BaseUrl, nil
}

func (i *IssuesService) BatchCreate(names []string) ([]string, error) {
	if len(names) == 0 {
		return nil, fmt.Errorf("names cannot be empty")
	}
	return []string{"hope", "mary", "saviour"}, nil
}
