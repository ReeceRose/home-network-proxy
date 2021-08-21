package repository

import (
	"fmt"
	"log"

	"github.com/ReeceRose/home-network-proxy/internal/consts"
	"github.com/ReeceRose/home-network-proxy/internal/database"
	"github.com/ReeceRose/home-network-proxy/internal/types"
	"github.com/ReeceRose/home-network-proxy/internal/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ipRepository struct {
	*baseRepository
}

var (
	_ IIPRepository = (*ipRepository)(nil)
)

// NewIPRepository returns an instanced ip repository
func NewIPRepository() IIPRepository {
	db, _ := database.Instance()

	return &ipRepository{
		baseRepository: &baseRepository{
			db:             db,
			collection:     db.Client().Database(utils.GetVariable(consts.DB_NAME)).Collection("ip"),
			collectionName: "ip",
		},
	}
}

// Find all IPs given a certain query
func (r *ipRepository) Find(query interface{}) ([]types.IP, error) {
	return r.FindWithFilter(query, nil)
}

// FindWithFilter returns all IPs given a certain query and find options
func (r *ipRepository) FindWithFilter(query interface{}, options *options.FindOptions) ([]types.IP, error) {
	cursor, err := r.collection.Find(r.db.Context(), query, options)
	if err != nil {
		msg := fmt.Sprintf("failed to read data from collection: %s with query: %s (%s)", r.collectionName, query, err.Error())
		log.Default().Println(msg)
		return nil, fmt.Errorf(msg)
	}

	var data []types.IP
	defer cursor.Close(r.db.Context())
	for cursor.Next(r.db.Context()) {
		var record types.IP
		if err = cursor.Decode(&record); err != nil {
			log.Default().Printf("failed to read record on %s with query: %s\n", r.collectionName, query)
		}
		data = append(data, record)
	}

	return data, nil
}

// Insert a single IP record into the database
func (r *ipRepository) Insert(data *types.IP) (string, error) {
	res, err := r.collection.InsertOne(r.db.Context(), data)
	if err != nil {
		msg := fmt.Sprintf("failed to insert data into collection: %s", r.collectionName)
		log.Default().Println(msg)
		return "", fmt.Errorf(msg)
	}
	return fmt.Sprintf("%x", res.InsertedID), nil
}

// UpdateById updates an existing ip record in the database
func (r *ipRepository) UpdateByID(data *types.IP) error {
	_, err := r.collection.UpdateByID(r.db.Context(), data.ID,
		bson.M{
			"$set": data,
		},
	)
	if err != nil {
		log.Default().Println(err.Error())
		return fmt.Errorf(err.Error())
	}
	return nil
}
