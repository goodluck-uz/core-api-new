package postgres

import (
	"context"
	"fmt"

	"github.com/goodluck-uz/core-api/api/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

var logPath2 = "storage/postgres/user.go"

type userRepo struct {
	db *pgxpool.Pool
}

func NewUserRepo(db *pgxpool.Pool) *userRepo {
	return &userRepo{
		db: db,
	}
}

// Create ...
func (r *userRepo) Create(ctx context.Context, req *models.CreateUserRequest) (*models.User, error) {
	var user models.User
	err := r.db.QueryRow(ctx, `
		INSERT INTO users (first_name, last_name, about, avatar, banner, username, password, country_id, role, category_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		RETURNING id, first_name, last_name, about, avatar, banner, username, password, country_id, role, category_id
	`, req.FirstName, req.LastName, req.About, req.Avatar, req.Banner, req.Username, req.Password, req.CountryID, req.Role, req.CategoryID).Scan(&user.ID, &user.FirstName, &user.LastName, &user.About, &user.Avatar, &user.Banner, &user.Username, &user.Password, &user.CountryID, &user.Role,
		&user.CategoryID)
	if err != nil {
		return nil, fmt.Errorf(logPath2, " userRepo.Create: %w", err)
	}
	return &user, nil
}

// Update ...
func (r *userRepo) Update(ctx context.Context, req *models.UpdateUserRequest) (*models.User, error) {
	var user models.User
	err := r.db.QueryRow(ctx, `
		UPDATE users
		SET first_name = $1, last_name = $2, about = $3, avatar = $4, banner = $5, username = $6, password = $7, country_id = $8, role = $9, category_id = $10
		WHERE id = $11
		RETURNING id, first_name, last_name, about, avatar, banner, username, password, country_id, role, category_id
	`, req.FirstName, req.LastName, req.About, req.Avatar, req.Banner, req.Username, req.Password, req.CountryID, req.Role, req.CategoryID, req.ID).Scan(&user.ID, &user.FirstName, &user.LastName, &user.About, &user.Avatar, &user.Banner, &user.Username, &user.Password, &user.CountryID, &user.Role, &user.CategoryID)
	if err != nil {
		return nil, fmt.Errorf(logPath2, " userRepo.Update: %w", err)
	}
	return &user, nil
}

// GetByID ...
func (r *userRepo) GetByID(ctx context.Context, id int) (*models.User, error) {
	var user models.User
	err := r.db.QueryRow(ctx, `
		SELECT id, first_name, last_name, about, avatar, banner, username, password, country_id, role, category_id
		FROM users
		WHERE id = $1
	`, id).Scan(&user.ID, &user.FirstName, &user.LastName, &user.About, &user.Avatar, &user.Banner, &user.Username, &user.Password, &user.CountryID, &user.Role, &user.CategoryID)
	if err != nil {
		return nil, fmt.Errorf(logPath2, " userRepo.GetByID: %w", err)
	}
	return &user, nil
}

// Delete ...
func (r *userRepo) Delete(ctx context.Context, id int) error {
	_, err := r.db.Exec(ctx, `
		DELETE FROM users
		WHERE id = $1
	`, id)
	if err != nil {
		return fmt.Errorf(logPath2, " userRepo.Delete: %w", err)
	}
	return nil
}

// GetList ...
func (r *userRepo) GetList(ctx context.Context, req *models.GetListUserRequest) (*models.ListUserResponse, error) {
	var users []*models.User
	rows, err := r.db.Query(ctx, `
		SELECT id, first_name, last_name, about, avatar, banner, username, password, country_id, role, category_id
		FROM users
		LIMIT $1 OFFSET $2
	`, req.Limit, req.Offset)
	if err != nil {
		return nil, fmt.Errorf(logPath2, " userRepo.GetList: %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.About, &user.Avatar, &user.Banner, &user.Username, &user.Password, &user.CountryID, &user.Role, &user.CategoryID)
		if err != nil {
			return nil, fmt.Errorf(logPath2, " userRepo.GetList: %w", err)
		}
		users = append(users, &user)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf(logPath2, " userRepo.GetList: %w", err)
	}
	return &models.ListUserResponse{
		Users: users,
	}, nil
}
