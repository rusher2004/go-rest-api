package datastore

import "github.com/google/uuid"

type DBClient interface {
	QueryUserByID(uuid.UUID) (User, error)
}

type DataStore struct {
	lambdaClient any
	db           DBClient
}

func NewDataStore(cl any, db DBClient) DataStore {
	return DataStore{cl, db}
}
