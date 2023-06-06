package customers

import (
	"time"

	"gorm.io/gorm"
)

type Customers struct {
	gorm.Model
	ID        int
	Firstname string
	Lastname  string
	Email     string
	Avatar    string
	CreatedAt time.Time
}
