package oldstore

import (
	"context"

	"github.com/rusher2004/go-rest-api/server"
)

type DataStore struct {
	lambda any
}

func NewDataStore(l any) DataStore {
	return DataStore{l}
}

func (d DataStore) DeleteUser(context.Context, server.DeleteUserInput) error {
	return server.ErrNotImplemented
}
func (d DataStore) GetUser(ctx context.Context, in server.GetUserInput) (server.User, error) {
	return server.User{}, nil
}
func (d DataStore) GetUserList(ctx context.Context, in server.GetUserListInput) ([]server.User, error) {
	return []server.User{}, nil
}
func (d DataStore) PostUser(ctx context.Context, in server.PostUserInput) (server.User, error) {
	return server.User{}, nil
}
func (d DataStore) PutUser(ctx context.Context, in server.PutUserInput) (server.User, error) {
	return server.User{}, nil
}
