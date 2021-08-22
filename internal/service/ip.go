package service

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/ReeceRose/home-network-proxy/internal/repository"
	"github.com/ReeceRose/home-network-proxy/internal/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ipService struct {
	ipRepository repository.IIPRepository
}

var (
	_ IIPService = (*ipService)(nil)
)

// NewIPService returns an instanced ip service
func NewIPService(ipRepository repository.IIPRepository) IIPService {
	return &ipService{
		ipRepository: ipRepository,
	}
}

// GetAllIP returns all known external IPs
func (s *ipService) GetAllIP(requestID string) types.IPResponse {
	log.Default().Println("attemping to get all IPs - Request ID: " + requestID)
	data, err := s.ipRepository.Find(bson.M{})
	if err != nil {
		return types.IPResponse{
			Data:       []types.IP{},
			StatusCode: http.StatusInternalServerError,
			Error:      fmt.Sprintf("failed to get all IPs - Request ID: %s", requestID),
			Success:    false,
		}
	}

	log.Default().Println("successfully got all IPs - Request ID: " + requestID)

	return types.IPResponse{
		Data:       data,
		StatusCode: http.StatusOK,
		Success:    true,
	}
}

func (s *ipService) GetIPByIP(requestID string, ip string) types.IPResponse {
	log.Default().Printf("attempt to get ip by ip: %s - Request ID: %s", ip, requestID)
	data, err := s.ipRepository.Find(bson.M{"externalIP": ip})
	if err != nil {
		return types.IPResponse{
			Data:       []types.IP{},
			StatusCode: http.StatusInternalServerError,
			Error:      fmt.Sprintf("failed to get IP: %s - Request ID: %s", ip, requestID),
			Success:    false,
		}
	}

	log.Default().Printf("successfully got IP: %s - Request ID: %s", ip, requestID)

	return types.IPResponse{
		Data:       data,
		StatusCode: http.StatusOK,
		Success:    true,
	}
}

// AddHealth inserts new health data for a given agent
func (s *ipService) InsertIP(requestID string, data *types.IP) types.IPResponse {
	log.Default().Printf("attemping to insert IP: %s - Request ID: %s", data.ExternalIP, requestID)

	now := time.Now().UTC().UnixNano()
	data.Updated = now

	res := s.GetIPByIP(requestID, data.ExternalIP)
	if res.Success && len(res.Data) >= 1 {
		first := res.Data[0]
		first.Updated = now
		data = &first

		err := s.ipRepository.UpdateByID(&first)

		if err != nil {
			return types.IPResponse{
				Data:       []types.IP{},
				StatusCode: http.StatusInternalServerError,
				Error:      fmt.Sprintf("failed to update existing IP: %s - Request ID: %s", data.ExternalIP, requestID),
			}
		}

		log.Default().Printf("successfully updated IP: %s - Request ID: %s", data.ExternalIP, requestID)
	} else {
		data.ID = primitive.NewObjectID()
		data.Created = now

		_, err := s.ipRepository.Insert(data)

		if err != nil {
			return types.IPResponse{
				Data:       []types.IP{},
				StatusCode: http.StatusInternalServerError,
				Error:      fmt.Sprintf("failed to insert IP: %s - Request ID: %s", data.ExternalIP, requestID),
				Success:    false,
			}
		}

		log.Default().Printf("successfully inserted IP: %s - Request ID: %s\n", data.ExternalIP, requestID)
	}

	return types.IPResponse{
		Data:       []types.IP{*data},
		StatusCode: http.StatusOK,
		Success:    true,
	}
}
