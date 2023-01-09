package pxthc

import (
	"context"
	"encoding/json"
	"github.com/earlofurl/pxthc/sqlc"
	"github.com/gobuffalo/nulls"
	"io"
	"time"
)

type UserService interface {
	CreateUser(ctx context.Context, user *CreateUserRequest) (*sqlc.User, error)
	FindUserByEmail(ctx context.Context, email string) (*sqlc.User, error)
	FindUserByID(ctx context.Context, id int64) (*sqlc.User, error)
	UpdateUser(ctx context.Context, upd *sqlc.UpdateUserParams) (*sqlc.User, error)
}

// UserFilter represents a filter passed to FindUsers().
type UserFilter struct {
	// Filtering fields.
	ID    *int    `json:"id"`
	Email *string `json:"email"`

	// Restrict to subset of results.
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

// UserUpdate represents a set of fields to be updated via UpdateUser().
type UserUpdate struct {
	FirstName *string `json:"first_name"`
	LastName  *string `json:"last_name"`
	Email     *string `json:"email"`
	Phone     *string `json:"phone"`
}

// GetUserByEmailRequest represents the params for getting a user by email.
type GetUserByEmailRequest struct {
	Email string `json:"email" validate:"required"`
}

// GetUserByIDRequest represents the params for getting a user by email.
type GetUserByIDRequest struct {
	ID int64 `json:"id" validate:"required"`
}

// CreateUserRequest represents a request to create a new user. Used in handler.
type CreateUserRequest struct {
	Username  string       `json:"username" validate:"required"`
	Email     string       `json:"email" validate:"required,email"`
	Password  string       `json:"password" validate:"required,min=6"`
	FirstName string       `json:"first_name" validate:"required"`
	LastName  string       `json:"last_name" validate:"required"`
	Phone     nulls.String `json:"phone" validate:"required"`
	Role      string       `json:"role" validate:"required"`
}

// CreatedUserResponse represents a response to a successful user creation.
type CreatedUserResponse struct {
	Username          string       `json:"username"`
	FirstName         string       `json:"first_name"`
	LastName          string       `json:"last_name"`
	Email             string       `json:"email"`
	Phone             nulls.String `json:"phone"`
	Role              string       `json:"role"`
	PasswordChangedAt time.Time    `json:"password_changed_at"`
	CreatedAt         time.Time    `json:"created_at"`
}

type UpdateUserRequest struct {
	Username  string       `json:"username" validate:"required"`
	Email     string       `json:"email" validate:"required,email"`
	Password  string       `json:"password" validate:"required,min=6"`
	FirstName string       `json:"first_name" validate:"required"`
	LastName  string       `json:"last_name" validate:"required"`
	Phone     nulls.String `json:"phone" validate:"required"`
	Role      string       `json:"role" validate:"required"`
}

func (r *GetUserByIDRequest) Bind(body io.ReadCloser) error {
	return json.NewDecoder(body).Decode(r)
}

func (r *GetUserByEmailRequest) Bind(body io.ReadCloser) error {
	return json.NewDecoder(body).Decode(r)
}

func (r *UpdateUserRequest) Bind(body io.ReadCloser) error {
	return json.NewDecoder(body).Decode(r)
}

func (r *CreateUserRequest) Bind(body io.ReadCloser) error {
	return json.NewDecoder(body).Decode(r)
}
