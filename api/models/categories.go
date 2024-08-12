package models

type Category struct {
	ID     int  `json:"id"`
	Status bool `json:"status"`
}

type CreateCategoryRequest struct {
	Status bool `json:"status"`
}

type GetCategoryResponse struct {
	ID     int  `json:"id"`
	Status bool `json:"status"`
}

type UpdateCategoryRequest struct {
	ID     int  `json:"id"`
	Status bool `json:"status"`
}

type PATCHCategoryRequest struct {
	Status bool `json:"status"`
}

type ListCategoryResponse struct {
	List  []*Category `json:"list"`
	Count int         `json:"count"`
}

type GetListCategoryRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}
