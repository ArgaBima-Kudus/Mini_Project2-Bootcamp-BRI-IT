package account

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type accountController struct {
	usecase *accountUsecase
}

func NewAccountController(usecase *accountUsecase) *accountController {
	return &accountController{
		usecase: usecase,
	}
}

type AccountItemResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role_id  int    `json:"role_id"`
	Verified string
	Active   string
}

type JwtResponse struct {
	Data string `json:"token"`
}

type CreateResponse struct {
	Message string              `json:"message"`
	Data    AccountItemResponse `json:"data"`
}

func EncryptPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func (c accountController) Create(req *CreateRequest) (*AccountItemResponse, error) {
	password, err := EncryptPassword(req.Password)
	if err != nil {
		return nil, err
	}

	fmt.Println(password)

	account := Actors{
		Username: req.Username,
		Password: password,
		Role_id:  req.Role_id,
		Verified: req.Verified,
		Active:   req.Active,
	}

	err = c.usecase.Create(&account)
	if err != nil {
		return nil, err
	}

	res := &CreateResponse{
		Message: "Success",
		Data: AccountItemResponse{
			ID:       int(account.ID),
			Username: account.Username,
			Password: account.Password,
			Role_id:  int(account.Role_id),
			Verified: account.Verified,
			Active:   account.Active,
		},
	}
	return &res.Data, nil
}

func ComparePassword(hashedPassword string, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return err
	}

	return nil
}

type LoginResponse struct {
	Message string      `json:"message"`
	Data    string `json:"data"`
}

type ReadByUsernameResponse struct {
	Message string              `json:"message"`
	Data    AccountItemResponse `json:"data"`
}

func (c accountController) ReadByUsername(username string) (*ReadByUsernameResponse, error) {
	fmt.Println(username)
	account, err := c.usecase.getByUsername(username)
	if err != nil {
		return nil, err
	}

	res := &ReadByUsernameResponse{
		Message: "Success",
		Data: AccountItemResponse{
			ID:       int(account.ID),
			Username: account.Username,
			Password: account.Password,
			Role_id:  int(account.Role_id),
			Verified: account.Verified,
			Active:   account.Active,
		},
	}
	return res, nil
}

func (c accountController) Login(req *LoginRequest) (*LoginResponse, error) {
	account, err := c.usecase.getByUsername(req.Username)
	if err != nil {
		return nil, err
	}

	err = ComparePassword(account.Password, req.Password)
	if err != nil {
		return nil, err
	}

	claims := jwt.MapClaims{
		"sub":  account.Role_id,
		"name": account.Username,
		"iat":  time.Now().Unix(),
		"exp":  time.Now().Add(time.Hour * 1).Unix(),
	}

	// Tandatangani token dengan kunci rahasia
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte("secret-key"))
	if err != nil {
		return nil, err
	}

	// Gunakan signedToken seperti yang Anda butuhkan

	res := &LoginResponse{
		Message: "Success",
		Data: signedToken,
	}
	return res, nil
}
