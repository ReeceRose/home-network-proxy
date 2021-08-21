package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type IPResponse struct {
	Data       []IP
	StatusCode int
	Error      string
	Success    bool
}

// IP is the database model used to store external IPs (and some other useful information)
type IP struct {
	ID         primitive.ObjectID `json:"id" bson:"_id"`
	ExternalIP string             `json:"externalIP" bson:"externalIP"`
	Created    int64              `json:"created" bson:"_created"`
	Updated    int64              `json:"updated" bson:"_updated"`
}
