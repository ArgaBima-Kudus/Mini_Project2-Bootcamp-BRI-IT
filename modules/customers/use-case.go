package customers

type Usecase struct {
	repo *Repository
}

func NewUsecase(repo *Repository) *Usecase {
	return &Usecase{
		repo: repo,
	}
}

func (u Usecase) Create(customers *Customers) error {
	return u.repo.Save(customers)
}

func (u Usecase) Read() ([]Customers, error) {
	return u.repo.FindAll()
}

func (u Usecase) ReadByID(id any) (Customers, error) {
	return u.repo.FindById(id)
}

func (u Usecase) Save(customers *Customers) error {
	return u.repo.Save(customers)
}

func (u Usecase) FindAll() ([]Customers, error) {
	return u.repo.FindAll()
}

// FindById is a function to find a customer by id
/*func (u Usecase) FindById(id int) (Customers, error) {
	return u.repo.FindById(strconv.Itoa(id))
}*/

func (u Usecase) UpdateByID(customers *Customers) error {
	return u.repo.UpdateByID(customers)
}

func (u Usecase) DeleteByID(customers *Customers) error {
	return u.repo.DeleteByID(customers)
}

func (u Usecase) getByEmail(email string) (Customers, error) {
	return u.repo.FindByEmail(email)
}
