package lty_models

import (
	"time"
)

type InspirationCategory struct {
	ID            int       `json:"id"`
	Name          string    `json:"name"`
	SubCategories string    `json:"sub_categories"` // Stored as JSON string in DB
	SortOrder     int       `json:"sort_order"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
