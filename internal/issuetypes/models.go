package issuetypes

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