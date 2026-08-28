package projecttypes

type ProjectTypeDefault struct {
	ProjectType        string
	ProjectTypeKey     string
	ProjectTemplateKey string
}

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
