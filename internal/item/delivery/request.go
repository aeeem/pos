package delivery

type SaveOrUpdate struct {
	ID       int64  `json:"id" validate:"-"`
	ItemName string `json:"name" validate:"required"`
	MaxPrice int64  `json:"max_price" validate:"required"`
}

type DeleteRequest struct {
	ID int64 `json:"id" validate:"required"`
}
