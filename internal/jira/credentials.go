package jira

import "fmt"

type JiraCredentials struct {
	Email   string
	Token   string
	BaseUrl string
}

func (j *JiraCredentials) Validate() error {
	var missing []string

	if j.Email == "" {
		missing = append(missing, "email")
	}

	if j.Token == "" {
		missing = append(missing, "token")
	}

	if j.BaseUrl == "" {
		missing = append(missing, "base url")
	}

	if len(missing) != 0 {
		return fmt.Errorf("missing required jira credentials, %v", missing)
	}

	return nil
}
