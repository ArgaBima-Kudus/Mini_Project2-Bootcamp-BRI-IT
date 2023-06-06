package account

import (
	"gorm.io/gorm"
)

type accountRepo struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) *accountRepo {
	return &accountRepo{db}
}

func (r accountRepo) Save(account *Actors) error {
	return r.db.Create(account).Error
}

func (r accountRepo) FindByUsername(username string) (Actors, error) {
	var actor Actors

	err := r.db.Where("username = ?", string(username)).First(&actor).Error
	if err != nil {
		return actor, err
	}

	return actor, nil
}
