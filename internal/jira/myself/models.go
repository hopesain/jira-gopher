package myself

type MyselfResponse struct {
	Self              string            `json:"self"`
	AccountID         string            `json:"accountId"`
	AccountType       string            `json:"accountType"`
	EmailAddress      string            `json:"emailAddress"`
	AvatarUrls        AvatarUrls        `json:"avatarUrls"`
	DisplayName       string            `json:"displayName"`
	Active            bool              `json:"active"`
	TimeZone          string            `json:"timeZone"`
	Locale            string            `json:"locale"`
	Groups            Groups            `json:"groups"`
	ApplicationRoles  ApplicationRoles  `json:"applicationRoles"`
	Expand            string            `json:"expand"`
}

type AvatarUrls struct {
	Size48 string `json:"48x48"`
	Size24 string `json:"24x24"`
	Size16 string `json:"16x16"`
	Size32 string `json:"32x32"`
}

type Groups struct {
	Size  int           `json:"size"`
	Items []GroupItem   `json:"items"`
}

type GroupItem struct {
	Name string `json:"name"`
}

type ApplicationRoles struct {
	Size  int                `json:"size"`
	Items []ApplicationRole  `json:"items"`
}

type ApplicationRole struct {
	Key  string `json:"key"`
	Name string `json:"name"`
}