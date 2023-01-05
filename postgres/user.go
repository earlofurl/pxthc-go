package postgres

import (
	"github.com/earlofurl/pxthc"
	"github.com/earlofurl/pxthc/sqlc"
	"github.com/rs/zerolog/log"
	"golang.org/x/net/context"
)

// Ensure service implements interface.
var _ pxthc.UserService = (*UserService)(nil)

type UserService struct {
	store sqlc.Store
}

func NewUserService(store *sqlc.Store) *UserService {
	return &UserService{store: *store}
}

// FindUserByEmail retrieves a user by Email along with their associated auth objects.
// Returns ENOTFOUND if user does not exist.
func (s *UserService) FindUserByEmail(ctx context.Context, email string) (*sqlc.User, error) {
	u, err := s.store.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if u == nil {
		return nil, pxthc.ErrNoRecord
	}
	return u, nil
}

// FindUserByID retrieves a user by ID along with their associated auth objects.
// Returns ENOTFOUND if user does not exist.
func (s *UserService) FindUserByID(ctx context.Context, id int64) (*sqlc.User, error) {
	u, err := s.store.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if u == nil {
		return nil, pxthc.ErrNoRecord
	}
	return u, nil
}

// CreateUser creates a new user.
func (s *UserService) CreateUser(ctx context.Context, r *pxthc.CreateUserRequest) (*sqlc.User, error) {
	log.Debug().Msg("User service Create method called")
	log.Debug().Msgf("userData: %+v", r)

	log.Debug().Msgf("pass: %+v", r.Password)
	hashedPassword, err := HashPassword(r.Password)
	if err != nil {
		return &sqlc.User{}, err
	}

	log.Debug().Msgf("hashed pass: %+v)", hashedPassword)

	arg := &sqlc.CreateUserParams{
		Username:       r.Username,
		HashedPassword: hashedPassword,
		FirstName:      r.FirstName,
		LastName:       r.LastName,
		Email:          r.Email,
		Phone:          r.Phone,
		Role:           r.Role,
	}

	return s.store.CreateUser(ctx, arg)
}

// UpdateUser updates an existing user.
func (s *UserService) UpdateUser(ctx context.Context, arg *sqlc.UpdateUserParams) (*sqlc.User, error) {
	return s.store.UpdateUser(ctx, arg)
}

// TODO: DeactivateUser deactivates an existing user.
