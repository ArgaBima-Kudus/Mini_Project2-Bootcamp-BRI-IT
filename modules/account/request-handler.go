package account

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CreateRequest struct {
	ID       uint
	Username string `json:"username"`
	Password string `json:"password"`
	Role_id  uint
	Verified string `json:"verified"`
	Active   string `json:"active"`
}

type RequestHandler struct {
	ctrl *accountController
}

func NewAccountRequestHandler(ctrl *accountController) *RequestHandler {
	return &RequestHandler{
		ctrl: ctrl,
	}
}

func DefaultRequestHandler(db *gorm.DB) *RequestHandler {
	return NewAccountRequestHandler(
		NewAccountController(
			NewAccountUsecase(
				NewAccountRepository(db),
			),
		),
	)
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func (h RequestHandler) Create(c *gin.Context) {
	var req CreateRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	res, err := h.ctrl.Create(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h RequestHandler) Login(c *gin.Context) {
	var req LoginRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	res, err := h.ctrl.Login(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}
	c.Header("Authorization", fmt.Sprint(res.Data))
	c.JSON(http.StatusOK, res)
}

func (h RequestHandler) ReadByUsername(c *gin.Context) {

	username := c.Query("username")

	res, err := h.ctrl.ReadByUsername(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
