package currencyService

import (
	"github.com/ruziba3vich/prosphere/internal/items/models/vidgets"
	"github.com/ruziba3vich/prosphere/internal/items/repo"
)

type (
	CurrencyService struct {
		storage repo.CurrencyRepository
	}
)

func New(storage repo.CurrencyRepository) repo.CurrencyRepository {
	return &CurrencyService{
		storage: storage,
	}
}

func (c *CurrencyService) GetAllCurrencies() (*vidgets.GetAllCurrenciesResponse, error) {
	return c.storage.GetAllCurrencies()
}

func (c *CurrencyService) GetCurrencyByCcy(req *vidgets.GetCurrencyByCcyRequest) (*vidgets.CurrencyData, error) {
	return c.storage.GetCurrencyByCcy(req)
}
