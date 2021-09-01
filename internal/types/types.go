package types

import "github.com/google/uuid"

type IPResponse struct {
	Data       []IP
	StatusCode int
	Error      string
	Success    bool
}

// IP is the database model used to store external IPs (and some other useful information)
type IP struct {
	ID         string `json:"id"`
	ExternalIP string `json:"externalIP"`
	Created    string `json:"created"`
	Updated    string `json:"updated"`
	UserId     string `json:"userId"`
}

// AgentInformation contains an ID which is used to differentiate between different agents
type ReportingToolAgent struct {
	ID uuid.UUID
}
