package client

type User struct {
	Username string
	Role     string
	Database string
}

type SecurityObject struct {
	Members SecurityComponent `json:"members,omitempty"`
	Admins  SecurityComponent `json:"admins,omitempty"`
}

type SecurityComponent struct {
	Roles []string `json:"roles,omitempty"`
	Names []string `json:"names,omitempty"`
}

type DataBasesInfo struct {
	DBNames []string `json:"names,omitempty"`
}
