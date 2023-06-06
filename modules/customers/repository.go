package customers

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

func (r Repository) Save(customers *Customers) error {
	return r.db.Create(customers).Error
}

func (r Repository) FindAll() ([]Customers, error) {
	var customers []Customers
	err := r.db.Find(&customers).Error
	if err != nil {
		return customers, err
	}
	return customers, nil
}

// FindById is a function to find a customer by id
func (r Repository) FindById(id any) (Customers, error) {
	var customers Customers
	err := r.db.Find(&customers, "id = ?", id).Error
	if err != nil {
		return customers, err
	}
	return customers, nil
}

func (r Repository) UpdateByID(customers *Customers) error {
	return r.db.Save(customers).Error
}

func (r Repository) DeleteByID(customers *Customers) error {
	return r.db.Delete(customers).Error
}

func (r Repository) FindByEmail(email string) (Customers, error) {
	var customers Customers
	err := r.db.Find(&customers, "email = ?", email).Error
	if err != nil {
		return customers, err
	}
	return customers, nil
}
