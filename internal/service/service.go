package service

import (
	"github.com/ReeceRose/home-network-proxy/internal/types"
)

// IIPService is an interface which provides method signatures for a IP service
type IIPService interface {
	GetIPByIP(requestID string, ip string) types.IPResponse
	GetAllIP(requestID string) types.IPResponse
	InsertIP(requestID string, data *types.IP) types.IPResponse
}
