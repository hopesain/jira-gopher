package jira

import "fmt"

// Jira Credentials
type Credentials struct {
	Email   string // email@email.com
	Token   string // jiraAccessToken
	BaseUrl string // https://your-domain.atlassian.net/rest/api/3
}

func (j *Credentials) Validate() error {
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
