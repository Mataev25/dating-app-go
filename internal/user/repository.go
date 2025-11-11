package user

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

type Repository interface {
	CreateUser(ctx context.Context, user *User) error
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	GetUserByID(ctx context.Context, id string) (*User, error)
}

type userRepository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(
	ctx context.Context, 
	user *User,
) error {
	user.ID = uuid.New().String()
	hashPassword, err := bcrypt.GenerateFromPassword(
		[]byte(user.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return fmt.Errorf("faild to hash password: %w", err) 
	}
	user.Password = string(hashPassword)

	now := time.Now()
	user.CreatedAt = now
	user.UpdatedAt = now

	query := `INSERT INTO users (
		id, email, password, name, age, gender,
		description, looking_for, created_at, updated_at
		 ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	_, err = r.db.ExecContext(
			 ctx, query, user.ID, user.Email,
			 user.Password, user.Name, user.Age, user.Gender,
		     user.Description, user.LookingFor, user.CreatedAt,
			 user.UpdatedAt,
	)

	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	return  nil
}

func (r *userRepository) GetUserByEmail(
	ctx context.Context, 
	email string,
) (*User, error) {
	var user User
	query := `SELECT * FROM users WHERE email = ?`
	err := r.db.GetContext(ctx, &user, query, email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found: %w", err)
		}
		return nil, fmt.Errorf(
			"failed to get user by email: %w",
			 err,
			)
	}
	return &user, nil
}

func (r *userRepository) GetUserByID(
	ctx context.Context, 
	id string,
) (*User, error) {
	var user User
	query := `SELECT * FROM users WHERE id = ?`
	err := r.db.GetContext(ctx, &user, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found: %w", err)
		}
		return nil, fmt.Errorf(
			"failed to get user by id: %w", 
			err,
			)
	}
	return &user, nil
}

















