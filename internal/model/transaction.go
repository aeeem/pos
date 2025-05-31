package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type JSON json.RawMessage

// Scan scan value into Jsonb, implements sql.Scanner interface
func (j *JSON) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	result := json.RawMessage{}
	err := json.Unmarshal(bytes, &result)
	*j = JSON(result)
	return err
}

// Value return json value, implement driver.Valuer interface
func (j JSON) Value() (driver.Value, error) {
	if len(j) == 0 {
		return nil, nil
	}
	return json.RawMessage(j).MarshalJSON()
}

type Status string

const (
	Active   Status = "completed"
	Inactive Status = "cancelled"
	Pending  Status = "pending"
)

type Transaction struct {
	gorm.Model
	CustomerName string  `json:"customer_name"`
	Status       Status  `gorm:"type:status" json:"status" `
	TotalPrice   float64 `json:"total_price"`
	Cart         []Cart  `json:"cart"`
	CartJson     *[]byte `json:"-"`
}
