package config

import "fmt"

// JIRA CREDENTIALS
type Credentials struct {
	Email   string
	Token   string
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
