package helper

type GetRequest struct {
	Page   int    `json:"page" validate:"required"`
	Limit  int    `json:"limit" validate:"required"`
	Search string `json:"search" validate:"-"`
}

func PageToOffset(page, limit int64) int64 { return (page - 1) * limit } // PageToOffset

type ListResponse struct {
	Total int64 `json:"total"`
	Page  int64 `json:"page"`
	Limit int64 `json:"limit"`
}
