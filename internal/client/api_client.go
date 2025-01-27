package client

import (
	"errors"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

var baseURL = "https://api.freecurrencyapi.com/v1" // baseURL ссылка для получения данных через сервис Freecurrencyapi

// GetData получает значение курса валюты в формате json
func GetData(apikey, status, rateCode, date string) ([]byte, error) {
	baseURL, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	baseURL.Path, err = url.JoinPath(baseURL.Path, status)
	if err != nil {
		return nil, err
	}

	params := url.Values{}
	params.Add("apikey", apikey)

	if rateCode != "" {
		params.Add("currencies", rateCode)
	}

	if status == "historical" {
		params.Add("date", date)
	}

	baseURL.RawQuery = params.Encode()

	req, err := http.NewRequest("GET", baseURL.String(), nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{Timeout: 10 * time.Second}

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("data is not valid")
	}

	defer resp.Body.Close()

	var body []byte

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return body, nil
}
