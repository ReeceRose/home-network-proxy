package types

type IPRequest struct {
	ExternalIP string `json:"externalIP"`
}

type IP struct {
	ExternalIP string `json:"externalIP" bson:"externalIP"`
	Created    int64  `json:"created" bson:"_created"`
	Updated    int64  `json:"updated" bson:"_updated"`
}
