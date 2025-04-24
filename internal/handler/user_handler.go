package handler

import (

	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/rmrachmanfauzan/bank_assessment/internal/model"
	"github.com/rmrachmanfauzan/bank_assessment/internal/repository"
	util "github.com/rmrachmanfauzan/bank_assessment/internal/utilities"
)

type UserHandler struct {
	repo repository.UserRepository
}



func NewUserHandler(r repository.UserRepository) *UserHandler {
	return &UserHandler{r}
}

// POST /users
func (h *UserHandler) RegisterUser(c echo.Context) error {
	var req model.User
	
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, util.ResponseMessage(nil, "Invalid Request Body", nil))
	}

	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, util.ResponseMessage(err, "Invalid Request Body", nil))
	}

	if err := h.repo.RegisterUser(&req); err != nil {
		if err.Error() == "NIK already exists" || err.Error() == "phone already exists" {
			
			return echo.NewHTTPError(http.StatusBadRequest, util.ResponseMessage(nil, err.Error(), nil))
		}
		
		return echo.NewHTTPError(http.StatusInternalServerError,util.ResponseMessage(err, "Error creating user",nil))
	}

	return c.JSON(http.StatusCreated, util.ResponseMessage(nil,"User created successfully", req))
}

// GET /users/:id
func (h *UserHandler) FindUser(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid user ID")
	}

	user, err := h.repo.Find(uint(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "User not found")
	}

	return c.JSON(http.StatusOK,  util.ResponseMessage(nil, "User Found", user))
}
