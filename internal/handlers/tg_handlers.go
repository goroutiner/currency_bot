package handlers

import (
	"currency_bot/internal/currency"
	"currency_bot/internal/entities"
	"currency_bot/internal/plotter"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	tele "gopkg.in/telebot.v4"
)

// SendStart отправляет пользователю список валют
func SendStart(c tele.Context) error {
	entities.ModeSelector.Inline(entities.ModeSelector.Row(entities.ExchangeBtn, entities.PlotterBtn))
	return c.Send("Давай начнем ▶️ Для начала выберите интересующий режим:", entities.ModeSelector)
}

// ChangeCurrency отправляет пользователю список валют для выбора или изменения
func ChangeCurrency(c tele.Context) error {
	entities.RatesSelector.Inline(entities.RatesSelector.Split(2, entities.BtnsRates)...)
	return c.Send("Выбери интересующую валюту для конвертации:", entities.RatesSelector)
}

// ChangeMode отправляет пользователю список режимов для выбора или изменения
func ChangeMode(c tele.Context) error {
	entities.ModeSelector.Inline(entities.ModeSelector.Row(entities.ExchangeBtn, entities.PlotterBtn))
	return c.Send("Выберите интересующий режим:", entities.ModeSelector)
}

// SendCurrencies отправляет пользователю список валют для конвертации
func SendCurrencies(c tele.Context) error {
	// сохраняем выбранный режим
	entities.CurrencyMode = c.Callback().Unique

	return ChangeCurrency(c)
}

// CheckingSelect проверяет сделан ли ранее выбор режима и валюты
func CheckingSelect(c tele.Context) bool {
	ok := true

	if entities.CurrencyMode == "" {
		ChangeMode(c)
		ok = false
	} else if entities.СurrencyCode == "" {
		ChangeCurrency(c)
		ok = false
	}

	return ok
}

// SendAsk просит пользователя выбрать критерий критерий курса по актуальности
// или ввести две даты в зависимости от режима
func SendAsk(c tele.Context) error {
	// сохранеяем код выбранной валюты
	entities.СurrencyCode = strings.TrimPrefix(c.Data(), "\f")

	switch entities.CurrencyMode {
	case entities.ExchangeBtn.Text:
		entities.StatusesSelector.Inline(entities.StatusesSelector.Row(entities.LatestBtn, entities.HistorBtn))
		return c.Send("Выберите критерий курса по актуальности:", entities.StatusesSelector)
	case entities.PlotterBtn.Text:
		return c.Send("Введите 2 разные даты, имеющие временной отрезок между собой не менее 9 дней, начиная с 1999 года в формате YYYY-MM-DD через пробел:")
	default:
		return ChangeMode(c)
	}
}

// SendLatestRate отправляет пользователю актуальный курс
func SendLatestRate(currency *currency.Service) tele.HandlerFunc {
	return func(c tele.Context) error {
		if ok := CheckingSelect(c); !ok {
			return nil
		}

		status := entities.LatestBtn.Text
		rateCode := entities.СurrencyCode

		res, err := currency.GetRate(status, rateCode, "")
		if err != nil {
			log.Println(err.Error())
			return nil
		}

		return c.Send(fmt.Sprintf("✅ %s", res))
	}
}

// SendAskDate отправляет пользователю сообщение, чтобы тот ввел дату
func SendAskDate(c tele.Context) error {
	if ok := CheckingSelect(c); !ok {
		return nil
	}

	return c.Send("Введите дату начиная с 1999 года в формате YYYY-MM-DD:")
}

// SendHistoricalRate отправляет пользователю исторический курс
func SendHistoricalRate(c tele.Context, currency *currency.Service) error {
	if ok := CheckingSelect(c); !ok {
		return nil
	}

	date := c.Text()
	status := entities.HistorBtn.Text
	rateCode := entities.СurrencyCode

	res, err := currency.GetRate(status, rateCode, date)
	if err != nil {
		log.Println(err.Error())
		return c.Send("Некорректно указана дата❗ Повторите ввод в формате YYYY-MM-DD:")
	}

	return c.Send(fmt.Sprintf("✅ %s", res))
}

// SendPlot отправляет пользователю график изменения курса в указанном отрезке времени
func SendPlot(c tele.Context, currency *currency.Service) error {
	if ok := CheckingSelect(c); !ok {
		return nil
	}

	warning := "Даты введены некорректно или временной отрезок между ними менее 9 дней❗\nВведите 2 разные даты, имеющие временной отрезок между собой не менее 9 дней, начиная с 1999 года в формате YYYY-MM-DD через пробел:"

	twoDate := strings.Split(c.Text(), " ")
	if len(twoDate) != 2 {
		return c.Send(warning)
	}

	if entities.Delay != nil {
		// отправляем пользователю уведомление об оставшемся времени до готовности графика
		func(c tele.Context) error {
			if time.Until(entities.RestartTime).Seconds() <= 0 {
				return nil
			}

			notify := fmt.Sprintf("До готовности графика осталось %.0f секунд(ы)... Пожалуйста подождите 😁", time.Until(entities.RestartTime).Seconds())
			return c.Send(notify)
		}(c)

		// ожидаем пока истечет 1 min
		<-entities.Delay.C
	}

	data, err := currency.GetPlotData(entities.СurrencyCode, twoDate[0], twoDate[1])
	if err != nil {
		log.Println(err.Error())
		return c.Send(warning)
	}

	plotFile, err := plotter.GetPlot(data)
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	defer os.Remove(plotFile)

	entities.Delay = time.NewTimer(time.Minute)
	entities.RestartTime = time.Now().Add(time.Minute)

	return c.Send(&tele.Photo{File: tele.FromDisk(plotFile)})
}

// SendResult обрабатывает текст в зависимости от выбранного режима и отправляет результат
func SendResult(currency *currency.Service) tele.HandlerFunc {
	return func(c tele.Context) error {
		switch entities.CurrencyMode {
		case entities.ExchangeBtn.Text:
			return SendHistoricalRate(c, currency)
		case entities.PlotterBtn.Text:
			return SendPlot(c, currency)
		default:
			return ChangeMode(c)
		}
	}
}
