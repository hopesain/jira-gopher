package projects

import "fmt"

type ProjectTypeDefault struct {
	ProjectType        string
	ProjectTypeKey     string
	ProjectTemplateKey string
}

type ProjectTypeResponse struct {
	Key                string `json:"key"`
	FormattedKey       string `json:"formattedKey"`
	DescriptionI18nKey string `json:"descriptionI18nKey"`
	Icon               string `json:"icon"`
	Color              string `json:"color"`
}

type ProjectsResponse struct {
	Self       string    `json:"self"`
	NextPage   string    `json:"nextPage,omitempty"`
	MaxResults int       `json:"maxResults"`
	StartAt    int       `json:"startAt"`
	Total      int       `json:"total"`
	IsLast     bool      `json:"isLast"`
	Values     []Project `json:"values"`
}

type Project struct {
	Expand          string           `json:"expand,omitempty"`
	Self            string           `json:"self"`
	ID              string           `json:"id"`
	Key             string           `json:"key"`
	Name            string           `json:"name"`
	AvatarUrls      AvatarUrls       `json:"avatarUrls"`
	ProjectTypeKey  string           `json:"projectTypeKey"`
	Simplified      bool             `json:"simplified"`
	Style           string           `json:"style"`
	IsPrivate       bool             `json:"isPrivate"`
	EntityID        string           `json:"entityId,omitempty"`
	UUID            string           `json:"uuid,omitempty"`
	ProjectCategory *ProjectCategory `json:"projectCategory,omitempty"`
	Insight         *ProjectInsight  `json:"insight,omitempty"`
}

type AvatarUrls struct {
	Size48 string `json:"48x48"`
	Size24 string `json:"24x24"`
	Size16 string `json:"16x16"`
	Size32 string `json:"32x32"`
}

type ProjectCategory struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Self        string `json:"self"`
}

type ProjectInsight struct {
	TotalIssueCount     int    `json:"totalIssueCount"`
	LastIssueUpdateTime string `json:"lastIssueUpdateTime"`
}

type CreateProjectRequest struct {
	AssigneeType       string `json:"assigneeType"`
	Description        string `json:"description"`
	Key                string `json:"key"`
	LeadAccountId      string `json:"leadAccountId"`
	Name               string `json:"name"`
	ProjectTemplateKey string `json:"projectTemplateKey"`
	ProjectTypeKey     string `json:"projectTypeKey"`
}

const (
	MAX_CHARACTER_LENGTH = 10
)

func (c *CreateProjectRequest) Validate() error {
	if c.AssigneeType != "UNASSIGNED" && c.AssigneeType != "PROJECT_LEAD" {
		return fmt.Errorf("assignee type must either be UNASSIGNED or PROJECT_LEAD")
	}

	if len(c.Key) > MAX_CHARACTER_LENGTH {
		return fmt.Errorf("the key exceeds maximum length: %v", c.Key)
	}

	var missing []string

	if c.AssigneeType == "" {
		missing = append(missing, "assign type")
	}

	if c.Description == "" {
		missing = append(missing, "description")
	}

	if c.Key == "" {
		missing = append(missing, "key")
	}

	if c.LeadAccountId == "" {
		missing = append(missing, "lead account id")
	}

	if c.Name == "" {
		missing = append(missing, "name")
	}

	if c.ProjectTemplateKey == "" {
		missing = append(missing, "project template key")
	}

	if c.ProjectTypeKey == "" {
		missing = append(missing, "project type key")
	}

	if len(missing) != 0 {
		return fmt.Errorf("missing required fields for creating a project in Jira: %v", missing)
	}

	return nil
}

type CreateProjectResponse struct {
	Self string `json:"self"`
	ID   int    `json:"id"`
	Key  string `json:"key"`
}
