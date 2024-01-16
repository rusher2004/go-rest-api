package oldstore

import (
	"context"

	"github.com/google/uuid"
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
	return server.User{
		Email: "rusher@example.com",
		ID:    in.ID,
		Name:  "rusher",
	}, nil
}
func (d DataStore) GetUserList(ctx context.Context, in server.GetUserListInput) ([]server.User, error) {
	return []server.User{
		{
			Email: "rusher@example.com",
			ID:    uuid.MustParse("07a86fdf-fe39-45ea-af57-e4aa9e8068b0"),
			Name:  "rusher",
		},
	}, nil
}
func (d DataStore) PostUser(ctx context.Context, in server.PostUserInput) (server.User, error) {
	return server.User{
		Email: in.Email,
		ID:    uuid.New(),
		Name:  in.Name,
	}, nil
}
func (d DataStore) PutUser(ctx context.Context, in server.PutUserInput) (server.User, error) {
	return server.User(in), nil
}
