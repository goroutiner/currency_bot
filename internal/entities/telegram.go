package entities

import (
	"time"

	tele "gopkg.in/telebot.v4"
)

var (
	Delay       *time.Timer // Delay таймер реализующий задержку между созданиями графиков (сервис Freecurrencyapi имеет ограничения - 10 req/min)
	RestartTime time.Time   // RestartTime необходим для вычисления остаточного времени до повторного создания графика

	Start          = "/start"           // Start команда для запуска бота
	ChangeMode     = "/change_mode"     // ChangeMode команда для смены режима
	ChangeCurrency = "/change_currency" // ChangeCurrency команда для смены валюты

	CurrencyMode string // CurrencyMode хранит в себе последней выбранный режим
	СurrencyCode string // СurrencyCode хранит в себе код последней выбранной валюты

	// ModeSelector реализует встроенные кнопки с выбором режима
	ModeSelector = &tele.ReplyMarkup{}
	PlotterBtn   = ModeSelector.Data("plotter", "plotter")
	ExchangeBtn  = ModeSelector.Data("exchange", "exchange")

	// RatesSelector реализует встроенные кнопки с выбором валюты для перевода
	RatesSelector = &tele.ReplyMarkup{}
	BtnsRates     = []tele.Btn{
		RatesSelector.Data("🇷🇺 RUB", "RUB"),
		RatesSelector.Data("🇺🇲 USD", "USD"),
		RatesSelector.Data("🇪🇺 EUR", "EUR"),
		RatesSelector.Data("🇨🇳 CNY", "CNY"),
		RatesSelector.Data("🇹🇷 TRY", "TRY"),
		RatesSelector.Data("🇬🇧 GBP", "GBP"),
		RatesSelector.Data("🇯🇵 JPY", "JPY"),
		RatesSelector.Data("🇮🇳 INR", "INR"),
	}

	// StatusesSelector реализует встроенные кнопки с выбором критерия актуальности
	StatusesSelector = &tele.ReplyMarkup{}
	LatestBtn        = StatusesSelector.Data("latest", "latest")
	HistorBtn        = StatusesSelector.Data("historical", "historical")
)
