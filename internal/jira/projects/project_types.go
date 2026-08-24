package projects

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/hopesain/jira-gopher/internal/jira"
)

var defaultProjectTypes = map[string]ProjectTypeDefault{
	"software": {
		ProjectType:        "software",
		ProjectTypeKey:     "software",
		ProjectTemplateKey: "com.pyxis.greenhopper.jira:gh-simplified-agility-kanban",
	},
	"business": {
		ProjectType:        "business",
		ProjectTypeKey:     "business",
		ProjectTemplateKey: "com.atlassian.jira-core-project-templates:jira-core-simplified-task-tracking",
	},
	"service desk": {
		ProjectType:        "service desk",
		ProjectTypeKey:     "service_desk",
		ProjectTemplateKey: "com.atlassian.servicedesk:simplified-it-service-management",
	},
	"customer service": {
		ProjectType:        "customer service",
		ProjectTypeKey:     "customer_service",
		ProjectTemplateKey: "com.atlassian.jcs:customer-service-management",
	},
}

func GetProjectTypeDefault(projectType string) (ProjectTypeDefault, error) {
	ptd, exists := defaultProjectTypes[projectType]
	if exists {
		return ptd, nil
	}
	return ProjectTypeDefault{}, fmt.Errorf("no default template found for project type: %v", projectType)
}

func GetAllProjectTypes(credentials jira.JiraCredentials) ([]ProjectTypeResponse, error) {
	client := &http.Client{
		Timeout: time.Second * 15,
	}

	url := credentials.BaseUrl + "/project/type"

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to build the request: %w", err)
	}

	req.SetBasicAuth(credentials.Email, credentials.Token)
	req.Header.Set("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to get all project types: %w", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read the response body: %w", err)
	}

	if resp.StatusCode == http.StatusUnauthorized {
		return nil, fmt.Errorf("incorrect authentication credentials: %v", string(body))
	}

	var response []ProjectTypeResponse

	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("failed to decode the response body: %w", err)
	}

	return response, nil

}
