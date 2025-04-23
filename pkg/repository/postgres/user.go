package postgres

import (
	"HR-monitor/pkg/models"
	"HR-monitor/pkg/repository"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type postgresUserRepository struct {
	db *pgxpool.Pool
}

func NewPostgresUserRepository(db *pgxpool.Pool) repository.UserRepository {
	return &postgresUserRepository{db: db}
}

func (r *postgresUserRepository) CreateUser(ctx context.Context, user models.User) error {
	query := `
	INSERT INTO users (email, password_hash, role, created_at, updated_at, name)
	VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING id
	`
	err := r.db.QueryRow(ctx, query,
		user.Email,
		user.PasswordHash,
		user.Role,
		user.CreatedAt,
		user.UpdatedAt,
		user.Name,
	).Scan(&user.ID)
	if err != nil {
		return err
	}
	return nil

}

func (r *postgresUserRepository) GetUserByID(ctx context.Context, id int) (models.User, error) {
	query := `
	SELECT id, email, password_hash, role, created_at, updated_at, name
	FROM users
	WHERE id = $1
	`
	var user models.User
	err := r.db.QueryRow(ctx, query, id).Scan(
		&user.ID,
		&user.Email,
		&user.PasswordHash,
		&user.Role,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.Name,
	)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (r *postgresUserRepository) GetUserByEmail(ctx context.Context, email string) (models.User, error) {
	query := `
	SELECT id, email, password_hash, role, created_at, updated_at, name 
	FROM users
	WHERE email = $1
	`
	var user models.User
	err := r.db.QueryRow(ctx, query, email).Scan(
		&user.ID,
		&user.Email,
		&user.PasswordHash,
		&user.Role,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.Name,
	)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (r *postgresUserRepository) DeleteUserByID(ctx context.Context, id int) error {
	query := `
	DELETE FROM users
	WHERE id = $1
	` 
	_, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}