package telegram

import (
	"currency_bot/internal/currency"
	"currency_bot/internal/entities"
	"currency_bot/internal/handlers"
	"log"
	"time"

	tele "gopkg.in/telebot.v4"
)

// Service структура методанных для сервиса telegram
type Service struct {
	bot *tele.Bot
}

// NewTelegramService создает новый telegram-сервис, в котором описаны handlers,
// обрабатывающие поступающие запросы
func NewTelegramService(token string, currency *currency.Service) *Service {
	pref := tele.Settings{
		Token:  token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	bot, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err.Error())
	}

	// обработка запроса при команде "/start"
	bot.Handle(entities.Start, handlers.SendStart)

	// обработка запроса при команде "/change_mode"
	bot.Handle(entities.ChangeMode, handlers.ChangeMode)

	// обработка запроса при команде "/change_currency"
	bot.Handle(entities.ChangeCurrency, handlers.ChangeCurrency)

	// обработка запроса при выборе режима "exchange"
	bot.Handle(&entities.ExchangeBtn, handlers.SendCurrencies)

	// обработка запроса при выборе режима "plotter"
	bot.Handle(&entities.PlotterBtn, handlers.SendCurrencies)

	// обработка запроса при выборе валюты
	bot.Handle(tele.OnCallback, handlers.SendAsk)

	// обработка запроса при выборе статуса "latest"
	bot.Handle(&entities.LatestBtn, handlers.SendLatestRate(currency))

	// обработка запроса при выборе статуса "historical"
	bot.Handle(&entities.HistorBtn, handlers.SendAskDate)

	// обработка запроса при получении дат(ы) в заданном формате
	bot.Handle(tele.OnText, handlers.SendResult(currency))

	return &Service{bot: bot}
}

// Start запускает телеграмм бота
func (s *Service) Start() {
	s.bot.Start()
}

// Stop останавливает телеграмм бота
func (s *Service) Stop() {
	s.bot.Stop()
}
