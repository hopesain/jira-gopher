package jira

import "fmt"

type Client struct {
	Credentials
	User string
	Age  int
}

func NewClient(jiraCredentials Credentials) (*Client, error) {
	if err := jiraCredentials.Validate(); err != nil {
		return nil, fmt.Errorf("validation error: %w", err)
	}
	return &Client{Credentials: jiraCredentials}, nil
}
