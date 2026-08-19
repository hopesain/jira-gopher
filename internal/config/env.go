package config

import (
	"os"
)

var (
	JIRA_API_TOKEN string
)

func GetEnvironmentVariables() {
	JIRA_API_TOKEN = os.Getenv("JIRA_API_TOKEN")
}
