package postgres

import (
	"context"
	"fmt"

	"github.com/goodluck-uz/core-api/api/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

var logPath1 = "storage/postgres/category.go"

type categoryRepo struct {
	db *pgxpool.Pool
}

func NewCategoryRepo(db *pgxpool.Pool) *categoryRepo {
	return &categoryRepo{
		db: db,
	}
}

// Create ...
func (r *categoryRepo) Create(ctx context.Context, req *models.CreateCategoryRequest) (*models.Category, error) {
	var category models.Category
	err := r.db.QueryRow(ctx, `
		INSERT INTO categories (status)
		VALUES ($1)
		RETURNING id, status
	`, req.Status).Scan(&category.ID, &category.Status)
	if err != nil {
		return nil, fmt.Errorf(logPath1, " categoryRepo.Create: %w", err)
	}
	return &category, nil
}

// Update ...
func (r *categoryRepo) Update(ctx context.Context, req *models.UpdateCategoryRequest) (*models.Category, error) {
	var category models.Category
	err := r.db.QueryRow(ctx, `
		UPDATE categories
		SET status = $1
		WHERE id = $2
		RETURNING id, status
	`, req.Status, req.ID).Scan(&category.ID, &category.Status)
	if err != nil {
		return nil, fmt.Errorf(logPath1, " categoryRepo.Update: %w", err)
	}
	return &category, nil
}

// GetByID ...
func (r *categoryRepo) GetByID(ctx context.Context, id int) (*models.Category, error) {
	var category models.Category
	err := r.db.QueryRow(ctx, `
		SELECT id, status
		FROM categories
		WHERE id = $1
	`, id).Scan(&category.ID, &category.Status)
	if err != nil {
		return nil, fmt.Errorf(logPath1, " categoryRepo.GetByID: %w", err)
	}
	return &category, nil
}

// Delete ...
func (r *categoryRepo) Delete(ctx context.Context, id int) error {
	_, err := r.db.Exec(ctx, `
		DELETE FROM categories
		WHERE id = $1
	`, id)
	if err != nil {
		return fmt.Errorf(logPath1, " categoryRepo.Delete: %w", err)
	}
	return nil
}

// GetList ...
func (r *categoryRepo) GetList(ctx context.Context, req *models.GetListCategoryRequest) (*models.ListCategoryResponse, error) {
	resp := &models.ListCategoryResponse{}
	var (
		query  string
		filter = " WHERE TRUE "
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query = `
		SELECT
			COUNT(*) OVER(),
			id, status
		FROM categories
	`
	if len(req.Search) > 0 {
		filter += " AND name ILIKE '%' || '" + req.Search + "' || '%' "
	}

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	query += filter + offset + limit

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf(logPath1, " categoryRepo.GetList: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var category models.Category
		err := rows.Scan(&resp.Count, &category.ID, &category.Status)
		if err != nil {
			return nil, fmt.Errorf(logPath1, " categoryRepo.GetList: %w", err)
		}
		resp.List = append(resp.List, &category)
	}

	return resp, nil

}
