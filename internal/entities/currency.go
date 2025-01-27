package entities

// LtsData структура для парсинга json актуального курса
type LtsData struct {
	Data map[string]float64 `json:"data"`
}

// HistData структура для парсинга json исторического курса
type HistData struct {
	Data map[string]map[string]float64 `json:"data"`
}

// CurrencyPlot структура метаданных для графика
type CurrencyPlot struct {
	DateList  []string
	RatesList []float64
	RateCode  string
}

var (
	// CountryСodes словарь, содержащий коды стран для соответсвующих валют
	CountryСodes = map[string]string{
		"EUR": "EU",
		"RUB": "RU",
		"USD": "US",
		"CNY": "CN",
		"TRY": "TR",
		"GBP": "GB",
		"JPY": "JP",
		"INR": "IN",
	}
)
