package jira

import (
	"github.com/hopesain/jira-gopher/internal/config"
	"github.com/hopesain/jira-gopher/internal/issues"
	"github.com/hopesain/jira-gopher/internal/projects"
)

type Credentials = config.Credentials

type ProjectsResponse = projects.ProjectsResponse
type GetProjectResponse = projects.GetProjectResponse

type CreateIssueResponse = issues.CreateIssueResponse