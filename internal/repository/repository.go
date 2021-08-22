package repository

import (
	"github.com/ReeceRose/home-network-proxy/internal/database"
	"github.com/ReeceRose/home-network-proxy/internal/types"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// IIPRepository is an interface which provides method signatures for a IP repository
type IIPRepository interface {
	Find(query interface{}) ([]types.IP, error)
	FindWithFilter(query interface{}, options *options.FindOptions) ([]types.IP, error)
	Insert(data *types.IP) (string, error)
	UpdateByID(data *types.IP) error
}

type baseRepository struct {
	db             database.Database
	collection     *mongo.Collection
	collectionName string
}
