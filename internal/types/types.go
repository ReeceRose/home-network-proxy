package types

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
	Created    int64  `json:"created"`
	Updated    int64  `json:"updated"`
}
