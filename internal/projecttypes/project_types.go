package projecttypes

import (
	"fmt"
	"net/http"

	"github.com/hopesain/jira-gopher/internal/config"
)

type ProjectTypesService struct {
	credentials config.Credentials
	httpClient  *http.Client
}

func New(credentials config.Credentials, httpClient *http.Client) *ProjectTypesService {
	return &ProjectTypesService{
		credentials: credentials,
		httpClient: httpClient,
	}
}

func (p *ProjectTypesService) GetDefaultProjectType(projectType string) (ProjectTypeDefault, error) {
	defaultProjectType, exists := defaultProjectTypes[projectType]
	if !exists {
		return ProjectTypeDefault{}, fmt.Errorf("no default template found for project type: %v", projectType)
	}
	return defaultProjectType, nil
}
