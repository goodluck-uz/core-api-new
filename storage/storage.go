package storage

import (
	"context"

	"github.com/goodluck-uz/core-api/api/models"
)

type StorageI interface {
	CloseDB()
	Category() CategoryRepoI
}

type CategoryRepoI interface {
	Create(ctx context.Context, req *models.CreateCategoryRequest) (*models.Category, error)
	Update(ctx context.Context, req *models.UpdateCategoryRequest) (*models.Category, error)
	GetByID(ctx context.Context, id int) (*models.Category, error)
	Delete(ctx context.Context, id int) error
	GetList(ctx context.Context, req *models.GetListCategoryRequest) (*models.ListCategoryResponse, error)
}
