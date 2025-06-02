package delivery

type SaveOrUpdate struct {
	CustomerName string `json:"customer_name" validate:"required"`
	PhoneNumber  string `json:"phone_number" validate:"required"`
}
