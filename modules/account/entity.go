package account

import "gorm.io/gorm"

type Actors struct {
	gorm.Model
	ID       uint
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password"`
	Role_id  uint
	Verified string `json:"verified"`
	Active   string `json:"active"`
}
