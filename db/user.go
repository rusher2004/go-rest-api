package db

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/rusher2004/go-rest-api/datastore"
)

type Email struct {
	Address   string
	CreatedAt time.Time
	ID        int
	UpdatedAt time.Time
	UUID      uuid.UUID
	Verified  bool
}

type User struct {
	Email Email
	ID    string
	Name  string
	UUID  uuid.UUID
}

var ErrUserNotFound = errors.New("user not found")

func (c *Client) QueryUserByID(id uuid.UUID) (datastore.User, error) {
	query := `
		SELECT
			u.uuid,
			u.name,
			e.address
		FROM users u
		INNER JOIN emails e ON e.user_id = u.id
		WHERE u.id = $1
	`
	row := c.db.QueryRow(query, id)

	var user User
	if err := row.Scan(&user.UUID, &user.Name, &user.Email.Address); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return datastore.User{}, ErrUserNotFound
		}
		return datastore.User{}, fmt.Errorf("error scanning user row: %w", err)
	}

	return datastore.User{
		Email: user.Email.Address,
		Name:  user.Name,
		UUID:  user.UUID,
	}, nil
}
