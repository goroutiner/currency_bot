package currency

import (
	"currency_bot/internal/client"
	"currency_bot/internal/entities"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/bojanz/currency"
)
// Service структура методанных для сервиса currency
type Service struct {
	Apikey       string            // Apikey ключ для получения данных курса валют из сервиса FreecurrencyAPI
	CountryСodes map[string]string // CountryСodes хранит коды локального располажения валют
}

// NewCurrencyService создает новый сервис для работы с валютой
func NewCurrencyService(apikey string) *Service {
	return &Service{
		Apikey:       apikey,
		CountryСodes: entities.CountryСodes,
	}
}

// GetSymbol метод для получения валютного символа
func (c *Service) GetSymbol(rateCode string) (string, error) {
	locCode := c.CountryСodes[rateCode]
	locale := currency.NewLocale(locCode)
	symbol, ok := currency.GetSymbol(rateCode, locale)
	if !ok {
		return "", errors.New("symbol for a currency code is not found")
	}

	return symbol, nil
}

// GetRate метод для получения значения курса валюты (текущий курс и исторический).
// status может принимать значения "latest" или "historical",
// если передан status = "latest" и rateID = "", то результатом будет актуальный курс.
func (c *Service) GetRate(status, rateCode, date string) (string, error) {
	jsonData, err := client.GetData(c.Apikey, status, rateCode, date)
	if err != nil {
		return "", err
	}

	var symbol string

	symbol, err = c.GetSymbol(rateCode)
	if err != nil {
		return "", err
	}

	var (
		ltsValue  entities.LtsData
		histValue entities.HistData
		rateValue float64
	)

	if status == "historical" {
		err = json.Unmarshal(jsonData, &histValue)
		rateValue = histValue.Data[date][rateCode]
	} else {
		err = json.Unmarshal(jsonData, &ltsValue)
		rateValue = ltsValue.Data[rateCode]
	}
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%.2f %s\n", rateValue, symbol), nil
}

// GetPlotData получает методанные, необходимые для создания графика 
func (c *Service) GetPlotData(rateCode, startDate, endDate string) (*entities.CurrencyPlot, error) {
	var (
		histValue entities.HistData
		parts     = 9
		status    = "historical"
	)

	leftEdge, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		return nil, err
	}

	rightEdge, err := time.Parse("2006-01-02", endDate)
	if err != nil {
		return nil, err
	}

	if leftEdge.After(rightEdge) {
		leftEdge, rightEdge = rightEdge, leftEdge
	}

	if rightEdge.Sub(leftEdge).Hours()/24 < 9 {
		return nil, errors.New("интервал между датами меньше 9 дней")
	}

	interval := rightEdge.Sub(leftEdge) / time.Duration(parts)

	dateList := make([]string, parts+1)
	dateList[0] = startDate
	dateList[parts-1] = endDate

	for i := 0; i <= parts; i++ {
		dateList[i] = leftEdge.Format("2006-01-02")
		leftEdge = leftEdge.Add(interval)
	}

	ratesList := make([]float64, parts+1)
	for i, date := range dateList {
		jsonData, _ := client.GetData(c.Apikey, status, rateCode, date)
		_ = json.Unmarshal(jsonData, &histValue)
		ratesList[i] = histValue.Data[date][rateCode]
	}

	return &entities.CurrencyPlot{DateList: dateList, RatesList: ratesList, RateCode: rateCode}, nil
}
