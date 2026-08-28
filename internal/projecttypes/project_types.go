package projecttypes

import (
	"fmt"
	"net/http"

	"github.com/hopesain/jira-gopher/internal/config"
)

type ProjectTypes struct {
	credentials config.Credentials
	httpClient  *http.Client
}

func (p *ProjectTypes) GetDefaultProjectType(projectType string) (ProjectTypeDefault, error) {
	defaultProjectType, exists := defaultProjectTypes[projectType]
	if !exists {
		return ProjectTypeDefault{}, fmt.Errorf("no default template found for project type: %v", projectType)
	}
	return defaultProjectType, nil
}
