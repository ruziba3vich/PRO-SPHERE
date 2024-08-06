package vidgets

type (
	CurrencyData struct {
		ID       int     `json:"id"`
		Code     string  `json:"Code"`
		Ccy      string  `json:"Ccy"`
		CcyNmRU  string  `json:"CcyNm_RU"`
		CcyNmUZ  string  `json:"CcyNm_UZ"`
		CcyNmUZC string  `json:"CcyNm_UZC"`
		CcyNmEN  string  `json:"CcyNm_EN"`
		Nominal  string  `json:"Nominal"`
		Rate     float64 `json:"Rate,string"`
		Diff     float64 `json:"Diff,string"`
		Date     string  `json:"Date"`
	}

	GetCurrencyByCcyRequest struct {
		Ccy string `json:"ccy"`
	}

	GetAllCurrenciesResponse struct {
		Response []CurrencyData `json:"response"`
	}
)

/*
	GetCurrencyByCcy()
	GetAllCurrencies()
*/
