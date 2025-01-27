package main

import (
	"currency_bot/internal/currency"
	"currency_bot/internal/telegram"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	currencyService := currency.NewCurrencyService(os.Getenv("BOT_TOKEN"))
	telegramService := telegram.NewTelegramService(os.Getenv("API_KEY"), currencyService)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		fmt.Println("Telegram bot starting...")	
		telegramService.Start()
	}()

	<-done

	fmt.Println("Telegram bot stopping...")
	telegramService.Stop()
	fmt.Println("Telegram bot stopped.")
}
