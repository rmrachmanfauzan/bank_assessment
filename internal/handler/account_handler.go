package handler

import (
	"net/http"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/rmrachmanfauzan/bank_assessment/internal/repository"
	util "github.com/rmrachmanfauzan/bank_assessment/internal/utilities"
)

type AccountHandler struct {
	repo repository.AccountRepository
}



func NewAccountHandler(r repository.AccountRepository) *AccountHandler {
	return &AccountHandler{r}
}

// POST /users
func (h *AccountHandler) TopupAccount(c echo.Context) error {
	type TopupAccountDTO struct {
		NoRekening string  `json:"no_rekening" validate:"required"`
		Nominal    float64 `json:"nominal" validate:"required,gt=0"`
	}

	var req TopupAccountDTO
	
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, util.ResponseMessage(nil, "Invalid Request Body", nil))
		
	}

	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, util.ResponseMessage(err, "Invalid Request Body", nil))
	}

	account, err := h.repo.TopupAccount(req.NoRekening,req.Nominal)
	if err != nil{
		if err.Error() == "rekening not found" {			
			return echo.NewHTTPError(http.StatusBadRequest, util.ResponseMessage(nil, err.Error(), nil))
		}
		
		return echo.NewHTTPError(http.StatusInternalServerError,util.ResponseMessage(err, "Error top up account",nil))
	}

	return c.JSON(http.StatusCreated, util.ResponseMessage(nil,"Top Up Account successfully", account))
}

func (h *AccountHandler) WithdrawAccount(c echo.Context) error {
	type WithdrawAccountDTO struct {
		NoRekening string  `json:"no_rekening" validate:"required"`
		Nominal    float64 `json:"nominal" validate:"required,gt=0"`
	}

	var req WithdrawAccountDTO
	
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, util.ResponseMessage(nil, "Invalid Request Body", nil))
	}

	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, util.ResponseMessage(err, "Invalid Request Body", nil))
	}

	account, err := h.repo.WithdrawAccount(req.NoRekening,req.Nominal)
	if err != nil{
		if err.Error() == "rekening not found" || err.Error() == "insufficient balance" {		
			return echo.NewHTTPError(http.StatusBadRequest, util.ResponseMessage(nil, err.Error(), nil))
		}
		
		return echo.NewHTTPError(http.StatusInternalServerError,util.ResponseMessage(err, "Error top up account",nil))
	}

	return c.JSON(http.StatusCreated, util.ResponseMessage(nil,"Withdraw Account successfully", account))
}


// GET /users/:id
func (h *AccountHandler) GetSaldo(c echo.Context) error {
	no_rekening := c.Param("no_rekening")
	
	if no_rekening == ""{
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid No Rekening")
	}

	user, err := h.repo.GetSaldo(string(no_rekening))
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Rekening not found")
	}

	return c.JSON(http.StatusOK,  util.ResponseMessage(nil, "Rekening Found", user))
}
