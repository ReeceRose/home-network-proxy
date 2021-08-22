package service

import (
	"fmt"
	"testing"

	"github.com/ReeceRose/home-network-proxy/internal/repository"
	"github.com/ReeceRose/home-network-proxy/internal/service/mocks"
	"github.com/ReeceRose/home-network-proxy/internal/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
)

//go:generate mockery --dir=../ -r --name IIPRepository

var (
	ipData []types.IP = []types.IP{
		{
			ExternalIP: "90.153.0.32",
			Created:    1,
			Updated:    1,
		},
		{
			ExternalIP: "90.153.2.16",
		},
	}
)

type testIPServiceHelper struct {
	ipService        IIPService
	ipRepository     repository.IIPRepository
	ipRepositoryMock *mock.Mock
}

func getInitializedIPService() testIPServiceHelper {
	ipRepository := new(mocks.IIPRepository)

	return testIPServiceHelper{
		ipService:        NewIPService(ipRepository),
		ipRepository:     ipRepository,
		ipRepositoryMock: &ipRepository.Mock,
	}
}

func TestIP_GetAllIP_ReturnsKnownIPs(t *testing.T) {
	helper := getInitializedIPService()
	helper.ipRepositoryMock.On("Find", bson.M{}).Return(ipData, nil)

	res := helper.ipService.GetAllIP("1")

	assert.Equal(t, 2, len(res.Data))
	assert.Equal(t, "", res.Error)
	assert.Equal(t, ipData[0].ExternalIP, res.Data[0].ExternalIP)
	assert.True(t, res.Success)
}

func TestIP_GetAllIP_HandlesError(t *testing.T) {
	helper := getInitializedIPService()
	helper.ipRepositoryMock.On("Find", bson.M{}).Return(nil, fmt.Errorf("failed to get IPs"))

	res := helper.ipService.GetAllIP("1")

	assert.Equal(t, 0, len(res.Data))
	assert.Equal(t, "failed to get all IPs - Request ID: 1", res.Error)
	assert.False(t, res.Success)
}

func TestIP_GetIPByIP_ReturnsKnownIPs(t *testing.T) {
	helper := getInitializedIPService()
	helper.ipRepositoryMock.On("Find", mock.Anything).Return([]types.IP{ipData[0]}, nil)

	res := helper.ipService.GetIPByIP("1", ipData[0].ExternalIP)

	assert.Equal(t, 1, len(res.Data))
	assert.Equal(t, "", res.Error)
	assert.Equal(t, ipData[0].ExternalIP, res.Data[0].ExternalIP)
	assert.True(t, res.Success)
}

func TestIP_GetIPByIP_ReturnsNoIPWithoutError(t *testing.T) {
	helper := getInitializedIPService()
	helper.ipRepositoryMock.On("Find", mock.Anything).Return([]types.IP{}, nil)

	res := helper.ipService.GetIPByIP("1", "90.153.0.36")

	assert.Equal(t, 0, len(res.Data))
	assert.Equal(t, "", res.Error)
	assert.True(t, res.Success)
}

func TestIP_GetIPByIP_HandlesError(t *testing.T) {
	helper := getInitializedIPService()
	helper.ipRepositoryMock.On("Find", mock.Anything).Return(nil, fmt.Errorf("failed to get IP"))

	res := helper.ipService.GetIPByIP("1", ipData[0].ExternalIP)

	assert.Equal(t, 0, len(res.Data))
	assert.Equal(t, "failed to get IP: 90.153.0.32 - Request ID: 1", res.Error)
	assert.False(t, res.Success)
}

func TestIP_InsertIP_InsertsIPIfNotExist(t *testing.T) {
	helper := getInitializedIPService()
	helper.ipRepositoryMock.On("Find", mock.Anything).Return([]types.IP{}, nil)
	helper.ipRepositoryMock.On("Insert", mock.Anything).Return("", nil)

	res := helper.ipService.InsertIP("1", &ipData[0])

	assert.Equal(t, 1, len(res.Data))
	assert.Equal(t, "", res.Error)
	assert.Equal(t, ipData[0].ExternalIP, res.Data[0].ExternalIP)
}

func TestIP_InsertIP_UpdatesUpdateTimeIfIPExists(t *testing.T) {
	helper := getInitializedIPService()
	helper.ipRepositoryMock.On("Find", mock.Anything).Return([]types.IP{ipData[0]}, nil)
	helper.ipRepositoryMock.On("UpdateByID", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	res := helper.ipService.InsertIP("1", &ipData[0])

	assert.Equal(t, 1, len(res.Data))
	assert.Equal(t, "", res.Error)
	assert.Equal(t, ipData[0].ExternalIP, res.Data[0].ExternalIP)
	assert.NotEqual(t, 1, res.Data[0].Updated) // 1 is the initial update time
}

func TestIP_InsertIP_HandlesInsertError(t *testing.T) {
	helper := getInitializedIPService()
	helper.ipRepositoryMock.On("Find", mock.Anything).Return([]types.IP{}, nil)
	helper.ipRepositoryMock.On("Insert", mock.Anything).Return("", fmt.Errorf("failed to insert IP"))

	res := helper.ipService.InsertIP("1", &ipData[0])

	assert.Equal(t, 0, len(res.Data))
	assert.Equal(t, "failed to insert IP: 90.153.0.32 - Request ID: 1", res.Error)
	assert.False(t, res.Success)
	assert.Equal(t, 500, res.StatusCode)
}

func TestIP_InsertIP_HandlesUpdateError(t *testing.T) {
	helper := getInitializedIPService()
	helper.ipRepositoryMock.On("Find", mock.Anything).Return([]types.IP{ipData[0]}, nil)
	helper.ipRepositoryMock.On("UpdateByID", mock.Anything, mock.Anything, mock.Anything).Return(fmt.Errorf("failed to update IP"))

	res := helper.ipService.InsertIP("1", &ipData[0])

	assert.Equal(t, 0, len(res.Data))
	assert.Equal(t, "failed to update existing IP: 90.153.0.32 - Request ID: 1", res.Error)
	assert.False(t, res.Success)
	assert.Equal(t, 500, res.StatusCode)
}
