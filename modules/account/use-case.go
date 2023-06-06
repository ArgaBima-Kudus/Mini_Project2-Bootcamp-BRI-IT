package account

type accountUsecase struct {
	repo *accountRepo
}

func NewAccountUsecase(repo *accountRepo) *accountUsecase {
	return &accountUsecase{
		repo: repo,
	}
}

func (u accountUsecase) Create(account *Actors) error {
	return u.repo.Save(account)
}

func (u accountUsecase) getByUsername(username string) (Actors, error) {
	return u.repo.FindByUsername(username)
}
