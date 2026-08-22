package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/hopesain/gojira/internal/jira"
	"github.com/hopesain/gojira/internal/jira/myself"
	"github.com/hopesain/gojira/internal/jira/projects"
	"github.com/joho/godotenv"
)

var (
	JIRA_EMAIL     string
	JIRA_API_TOKEN string
	JIRA_BASE_URL  string
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stderr, nil))
	slog.SetDefault(logger)

	if err := godotenv.Load(); err != nil {
		slog.Error("failed to load .env file", "error", err)
		os.Exit(1)
	}

	JIRA_EMAIL = os.Getenv("JIRA_EMAIL")
	JIRA_API_TOKEN = os.Getenv("JIRA_API_TOKEN")
	JIRA_BASE_URL = os.Getenv("JIRA_BASE_URL")

	jiraCredentials := jira.JiraCredentials{
		Email:   JIRA_EMAIL,
		Token:   JIRA_API_TOKEN,
		BaseUrl: JIRA_BASE_URL,
	}

	if err := jiraCredentials.Validate(); err != nil {
		slog.Error("validation error", "error", err)
		os.Exit(1)
	}

	accountID, err := myself.UserAccountID(jiraCredentials)
	if err != nil {
		slog.Error("failed to retrieve the user account ID", "error", err)
		os.Exit(1)
	}

	projectType, err := projects.GetProjectTypeDefault("business")
	if err != nil {
		slog.Error("failed to retrieve the default project types and templates", "error", err)
	}

	createProjectPayload := projects.CreateProjectRequest{
		AssigneeType:       "UNASSIGNED",
		Description:        "business as usual",
		Key:                "PB",
		LeadAccountId:      accountID,
		Name:               "Parrot Business",
		ProjectTemplateKey: projectType.ProjectTemplateKey,
		ProjectTypeKey:     projectType.ProjectTypeKey,
	}

	res, err := projects.CreateProject(jiraCredentials, createProjectPayload)
	if err != nil {
		slog.Error("failed to create a project", "error", err)
		os.Exit(1)
	}

	fmt.Println(res)

	fmt.Println("-------------------------------------------------------------------")

	projects, err := projects.GetProjects(jiraCredentials)
	if err != nil {
		slog.Error("failed to retrieve all projects", "error", err)
		os.Exit(1)
	}
	fmt.Println(projects)

}
