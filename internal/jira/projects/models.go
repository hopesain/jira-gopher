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
