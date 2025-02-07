package handlers

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ruziba3vich/prosphere/internal/items/models/errors"
	"github.com/ruziba3vich/prosphere/internal/items/models/vidgets"
	"github.com/ruziba3vich/prosphere/internal/items/repo"
)

type (
	CurrencyHandler struct {
		storage repo.CurrencyRepository
		logger  *log.Logger
	}
)

func NewCurrencyHandler(storage repo.CurrencyRepository, logger *log.Logger) *CurrencyHandler {
	return &CurrencyHandler{
		storage: storage,
		logger:  logger,
	}
}

func (c *CurrencyHandler) GetCurrencyByCcy(ctx *gin.Context) {
	var req vidgets.GetCurrencyByCcyRequest
	req.Ccy = ctx.Param("ccy")
	req.Ccy = strings.ToUpper(req.Ccy)
	response, err := c.storage.GetCurrencyByCcy(&req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errors.ProError{
			Message: "service returned an error while retrieving the data",
			Err:     err.Error(),
		})
		return
	}
	ctx.IndentedJSON(http.StatusOK, response)
}

func (c *CurrencyHandler) GetAllCurrenciesHandler(ctx *gin.Context) {
	response, err := c.storage.GetAllCurrencies()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errors.ProError{
			Message: "service returned an error while retrieving the data",
			Err:     err.Error(),
		})
		return
	}
	ctx.IndentedJSON(http.StatusOK, response)
}
