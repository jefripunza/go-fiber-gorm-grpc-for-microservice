package response

import (
	"time"

	"gorm.io/gorm"
)

type ReadProducts struct {
	ID        uint           `json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-"`
	Code      string         `json:"code"`
	Price     uint           `json:"price"`
}

type ProductAdd200 struct {
	Message string `json:"message"`
}

type ProductAdd400 struct {
	Message string `json:"message"`
}
