package main

import (
	"support-bot/cmd"
	"support-bot/internal/bot"
	"support-bot/internal/db"
	"support-bot/internal/log"
	"support-bot/internal/translations"
)

// Загрузка конфигурации
func main() {
	cfg, err := cmd.LoadConfig(".", ".env", "env")
	if err != nil {
		log.Error().Err(err).Send()
		panic(err)
	}

	// Инициализация переводчика
	translator, err := translations.NewTranslator("./locales", cfg.DefaultLocale, cfg.BotPrefix)
	if err != nil {
		log.Error().Err(err).Send()
		panic(err)
	}
	// Инициализация базы данных
	database, err := db.InitDB()
	if err != nil {
		log.Error().Err(err).Send()
		panic(err)
	}
	// Закрытие базы данных при завершении программы
	defer func(database *db.DB) {
		err := database.Close()
		if err != nil {
			log.Error().Err(err).Send()
			panic(err)
		}
	}(database)
	// Инициализация бота
	err = bot.InitBot(cfg.TelegramLoggerBotToken, cfg.TelegramAdminUserID, translator, database)
	if err != nil {
		log.Error().Err(err).Send()
		panic(err)
	}
}
