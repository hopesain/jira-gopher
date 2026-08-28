package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/hopesain/jira-gopher/internal/config"
	"github.com/hopesain/jira-gopher/internal/issues"
	"github.com/hopesain/jira-gopher/jira"
	"github.com/joho/godotenv"
)

// This is just a testing playground.
func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stderr, nil))
	slog.SetDefault(logger)

	if err := godotenv.Load(); err != nil {
		slog.Error("failed to load .env file", "error", err)
		os.Exit(1)
	}

	JIRA_EMAIL := os.Getenv("JIRA_EMAIL")
	JIRA_API_TOKEN := os.Getenv("JIRA_API_TOKEN")
	JIRA_BASE_URL := os.Getenv("JIRA_BASE_URL")

	creds := config.Credentials{
		Email:   JIRA_EMAIL,
		Token:   JIRA_API_TOKEN,
		BaseUrl: JIRA_BASE_URL,
	}

	client, err := jira.NewClient(creds)
	if err != nil {
		fmt.Println("Unable to create a new client", err)
		os.Exit(1)
	}

	fmt.Println("------------------ Create Issue -------------------------- ")

	issue := issues.CreateIssueRequest{
		Fields: issues.CreateIssueFields{
			Project: issues.ProjectRef{
				ID: "projectID",
			},
			IssueType: issues.IssueTypeRef{
				ID: "issueTypeID",
			},
			Summary: "New Jira Issue",
		},
	}

	issueCreate, err := client.Issues.Create(issue)
	if err != nil {
		fmt.Println("error: ", err)
		os.Exit(1)
	}
	fmt.Println("Create Issue Response: ", issueCreate)

	fmt.Println("---------------------- User Account ID ---------------------")

	accountID, err := client.Myself.UserAccountID()
	if err != nil {
		fmt.Println("err to retrieve the id", err)
		os.Exit(1)
	}
	fmt.Println("Account ID: ", accountID)

	ptd, err := client.ProjectTypes.GetDefaultProjectType("software")
	if err != nil {
		fmt.Println("error: ", err)
		os.Exit(1)
	}
	fmt.Println("Project Type: ", ptd)

	fmt.Println("----------------- Get Project -----------------------")
	projectResponse, err := client.Projects.Get("projectIDorKey")
	if err != nil {
		fmt.Println("error: ", err)
		os.Exit(1)
	}

	fmt.Println("---------------------- Issue Types ------------------------")

	issueTypes := projectResponse.IssueTypes
	for _, issueType := range issueTypes {
		fmt.Println(issueType)
		fmt.Println("---------------- Another One ----------------------")
	}

	fmt.Println("----------- Get All Projects -------------")
	allProjects, err := client.Projects.GetAll()
	if err != nil {
		fmt.Println("error", err)
	}
	fmt.Println(allProjects)
}
