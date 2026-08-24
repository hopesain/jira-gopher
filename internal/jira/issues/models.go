package issues

type IssueType struct {
	Self             string          `json:"self"`
	ID               string          `json:"id"`
	Description      string          `json:"description"`
	IconURL          string          `json:"iconUrl"`
	Name             string          `json:"name"`
	UntranslatedName string          `json:"untranslatedName"`
	Subtask          bool            `json:"subtask"`
	AvatarID         int             `json:"avatarId,omitempty"`
	HierarchyLevel   int             `json:"hierarchyLevel"`
	Scope            *IssueTypeScope `json:"scope,omitempty"`
}

type IssueTypeScope struct {
	Type    string        `json:"type"`
	Project ScopedProject `json:"project"`
}

type ScopedProject struct {
	ID string `json:"id"`
}

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
