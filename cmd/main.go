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

	fmt.Println("All credentials available")
	accountID := "712020:825cff40-821e-40d6-b058-c9d47a88702b"

	projectType, err := projects.GetProjectTypeDefault("software")
	if err != nil {
		slog.Error("failed to resolve project type default", "error", err)
		os.Exit(1)
	}
	fmt.Println(projectType, accountID)

	// createProjectPayload := projects.CreateProjectRequest{
	// 	AssigneeType:       "UNASSIGNED",
	// 	Description:        "programming guidelines",
	// 	Key:                "PGL",
	// 	LeadAccountId:      accountID,
	// 	Name:               "programming guidelines",
	// 	ProjectTemplateKey: projectType.ProjectTemplateKey,
	// 	ProjectTypeKey:     projectType.ProjectTypeKey,
	// }

	// if err := projects.CreateProject(jiraCredentials, createProjectPayload); err != nil {
	// 	slog.Error("failed to create a project", "error", err)
	// 	os.Exit(1)
	// }

	id, err := myself.UserAccountID(jiraCredentials)
	println("Account ID", id)

	pTypes, err := projects.GetAllProjectTypes(jiraCredentials)
	if err != nil {
		slog.Error("failed to load all project types", "error", err)
	}

	fmt.Println(pTypes)

}
