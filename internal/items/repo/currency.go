package repo

import "github.com/ruziba3vich/prosphere/internal/items/models/vidgets"

type (
	CurrencyRepository interface {
		GetCurrencyByCcy(*vidgets.GetCurrencyByCcyRequest) (*vidgets.CurrencyData, error)
		GetAllCurrencies() (*vidgets.GetAllCurrenciesResponse, error)
	}
)
