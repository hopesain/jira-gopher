package projects

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