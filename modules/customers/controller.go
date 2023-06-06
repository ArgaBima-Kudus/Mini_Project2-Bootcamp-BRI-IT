package customers

type Controller struct {
	usecase *Usecase
}

func NewController(usecase *Usecase) *Controller {
	return &Controller{
		usecase: usecase,
	}
}

type CreateResponse struct {
	Message string           `json:"message"`
	Data    UserItemResponse `json:"data"`
}

type UserItemResponse struct {
	ID        int    `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Avatar    string `json:"avatar"`
}

func (c Controller) Create(req *CreateRequest) (*CreateResponse, error) {
	user := Customers{
		Firstname: req.Firstname,
		Lastname:  req.Lastname,
		Email:     req.Email,
		Avatar:    req.Avatar,
	}

	err := c.usecase.Save(&user)
	if err != nil {
		return nil, err
	}

	res := &CreateResponse{
		Message: "Success",
		Data: UserItemResponse{
			ID:        int(user.ID),
			Firstname: user.Firstname,
			Lastname:  user.Lastname,
			Email:     user.Email,
			Avatar:    user.Avatar,
		},
	}
	return res, nil
}

type ReadResponse struct {
	Data []UserItemResponse `json:"data"`
}

func (c Controller) Read() (*ReadResponse, error) {
	users, err := c.usecase.Read()
	if err != nil {
		return nil, err
	}

	res := &ReadResponse{}
	for _, user := range users {
		res.Data = append(res.Data, UserItemResponse{
			ID:        int(user.ID),
			Firstname: user.Firstname,
			Lastname:  user.Lastname,
			Email:     user.Email,
			Avatar:    user.Avatar,
		})
	}

	return res, nil
}

/* ReadByID */
type ReadByIDResponse struct {
	Data UserItemResponse `json:"data"`
}

func (c Controller) ReadByID(id any) (*ReadByIDResponse, error) {
	user, err := c.usecase.ReadByID(id)
	if err != nil {
		return nil, err
	}

	res := &ReadByIDResponse{
		Data: UserItemResponse{
			ID:        int(user.ID),
			Firstname: user.Firstname,
			Lastname:  user.Lastname,
			Email:     user.Email,
			Avatar:    user.Avatar,
		},
	}

	return res, nil
}

type UpdateResponse struct {
	Message string           `json:"message"`
	Data    UserItemResponse `json:"data"`
}

func (c Controller) UpdateByID(req *CreateRequest, id any) (*CreateResponse, error) {
	user, err := c.usecase.ReadByID(id)
	if err != nil {
		return nil, err
	}

	customer := &Customers{
		ID:        user.ID,
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Email:     user.Email,
		Avatar:    user.Avatar,
	}

	customer.Firstname = req.Firstname
	customer.Lastname = req.Lastname
	customer.Email = req.Email
	customer.Avatar = req.Avatar

	createdAt := user.CreatedAt
	customer.CreatedAt = createdAt

	err = c.usecase.UpdateByID(customer)
	if err != nil {
		return nil, err
	}
	response := &CreateResponse{
		Message: "Success",
		Data: UserItemResponse{
			ID:        customer.ID,
			Firstname: customer.Firstname,
			Lastname:  customer.Lastname,
			Email:     customer.Email,
			Avatar:    customer.Avatar,
		},
	}

	return response, nil
}

type DeleteResponse struct {
	Message string `json:"message"`
}

func (c Controller) DeleteByID(id any) (*DeleteResponse, error) {
	user, err := c.usecase.ReadByID(id)
	if err != nil {
		return nil, err
	}

	customer := &Customers{
		ID:        user.ID,
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Email:     user.Email,
		Avatar:    user.Avatar,
	}

	if err := c.usecase.DeleteByID(customer); err != nil {
		return nil, err
	}

	res := &DeleteResponse{
		Message: "Success",
	}

	return res, nil
}

type getByEmailResponse struct {
	Status string
	Data   []UserItemResponse
}

func (c Controller) getByEmail(email string) (*getByEmailResponse, error) {
	user, err := c.usecase.getByEmail(email)
	if err != nil {
		return nil, err
	}

	res := &getByEmailResponse{
		Status: "Success",
		Data: []UserItemResponse{
			{
				ID:        int(user.ID),
				Firstname: user.Firstname,
				Lastname:  user.Lastname,
				Email:     user.Email,
				Avatar:    user.Avatar,
			},
		},
	}

	return res, nil
}
