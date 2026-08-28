package jira

import (
	"fmt"
	"net/http"
	"time"

	"github.com/hopesain/jira-gopher/internal/config"
	"github.com/hopesain/jira-gopher/internal/issues"
	"github.com/hopesain/jira-gopher/internal/issuetypes"
	"github.com/hopesain/jira-gopher/internal/myself"
	"github.com/hopesain/jira-gopher/internal/projects"
	"github.com/hopesain/jira-gopher/internal/projecttypes"
)

type Client struct {
	credentials  config.Credentials
	httpClient   *http.Client
	Projects     *projects.ProjectsService
	ProjectTypes *projecttypes.ProjectTypesService
	Issues       *issues.IssuesService
	IssueTypes   *issuetypes.IssuesTypesService
	Myself       *myself.MyselfService
}

func NewClient(credentials config.Credentials) (*Client, error) {
	if err := credentials.Validate(); err != nil {
		return nil, fmt.Errorf("validation error: %w", err)
	}

	httpClient := &http.Client{Timeout: time.Second * 20}

	return &Client{
		credentials:  credentials,
		httpClient:   httpClient,
		Projects:     projects.New(credentials, httpClient),
		ProjectTypes: projecttypes.New(credentials, httpClient),
		Issues:       issues.New(credentials, httpClient),
		IssueTypes:   issuetypes.New(credentials, httpClient),
		Myself:       myself.New(credentials, httpClient),
	}, nil
}
