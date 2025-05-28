package delivery

type GetRequest struct {
	Page   int    `json:"page" validate:"required"`
	Limit  int    `json:"limit" validate:"required"`
	Search string `json:"search" validate:"-"`
}

type SaveOrUpdate struct {
	ID       int64  `json:"id" validate:"-"`
	ItemName string `json:"name" validate:"required"`
	MaxPrice int64  `json:"max_price" validate:"required"`
}

type DeleteRequest struct {
	ID int64 `json:"id" validate:"required"`
}
