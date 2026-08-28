package issues

type CreateIssueRequest struct {
	Fields CreateIssueFields `json:"fields"`
}

type CreateIssueFields struct {
	Project   ProjectRef   `json:"project"`
	IssueType IssueTypeRef `json:"issuetype"`
	Summary   string       `json:"summary"`
}

type ProjectRef struct {
	ID string `json:"id"`
}

type IssueTypeRef struct {
	ID string `json:"id"`
}

type CreateIssueResponse struct {
	ID   string `json:"id"`
	Key  string `json:"key"`
	Self string `json:"self"`
}
