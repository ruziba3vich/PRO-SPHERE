package storage

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/ruziba3vich/prosphere/internal/items/config"
	"github.com/ruziba3vich/prosphere/internal/items/models/vidgets"
)

type (
	CurrencyStorage struct {
		config *config.Config
		logger *log.Logger
	}
)

func NewCurrencyStorage(cfg *config.Config, logger *log.Logger) *CurrencyStorage {
	return &CurrencyStorage{
		config: cfg,
		logger: logger,
	}
}

func (c *CurrencyStorage) GetCurrencyByCcy(req *vidgets.GetCurrencyByCcyRequest) (*vidgets.CurrencyData, error) {
	data, err := c.getAllCurrencyData()
	if err != nil {
		c.logger.Println(err.Error())
		return nil, err
	}
	for _, currency := range data {
		if currency.Ccy == req.Ccy {
			return currency, nil
		}
	}

	return nil, fmt.Errorf("currency not found: %s", req.Ccy)
}

func (c *CurrencyStorage) GetAllCurrencies() (*vidgets.GetAllCurrenciesResponse, error) {
	data, err := c.getAllCurrencyData()
	if err != nil {
		c.logger.Println(err.Error())
		return nil, err
	}
	var response vidgets.GetAllCurrenciesResponse

	response.Response = append(response.Response, data...)
	return &response, nil
}

func (c *CurrencyStorage) getAllCurrencyData() ([]*vidgets.CurrencyData, error) {
	resp, err := http.Get(c.config.Apis.CurrencyApi)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %s", err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %s", err.Error())
	}

	var data []*vidgets.CurrencyData
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, fmt.Errorf("failed to parse response: %s", err.Error())
	}
	return data, nil
}
