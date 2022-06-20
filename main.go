package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var numericInlineKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData(comoDepositar.question, comoDepositar.shortAnswer),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData(comoApostar.question, comoApostar.shortAnswer),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData(bonoDebut.question, bonoDebut.shortAnswer),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData(comoDescargar.question, comoDescargar.shortAnswer),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData(tutoriales.question, tutoriales.shortAnswer),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData(comoRetirar.question, comoRetirar.shortAnswer),
	),
)

var numericKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(comoDepositar.question),
		tgbotapi.NewKeyboardButton(comoApostar.question),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(bonoDebut.question),
		tgbotapi.NewKeyboardButton(comoDescargar.question),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(tutoriales.question),
		tgbotapi.NewKeyboardButton(comoRetirar.question),
	),
)

func main() {
	bot, err := tgbotapi.NewBotAPI("5428271492:AAGCAeJPxZlJW6p6jv9htw8MIWZvHD4leHM")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			wplay_generic := "Canal POWEROSO"
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, wplay_generic)
			//msg.ReplyToMessageID = update.Message.MessageID
			//bot.Send(msg)

			switch update.Message.Text {
			case "open":
				msg.ReplyMarkup = numericInlineKeyboard
			case "other":
				msg.ReplyMarkup = numericKeyboard
			case comoDepositar.question:
				msg.Text = comoDepositar.fullAnswer
			case comoApostar.question:
				msg.Text = comoApostar.fullAnswer
			case bonoDebut.question:
				msg.Text = bonoDebut.fullAnswer
			case comoDescargar.question:
				msg.Text = comoDescargar.fullAnswer
			case tutoriales.question:
				msg.Text = tutoriales.fullAnswer
			case comoRetirar.question:
				msg.Text = comoRetirar.fullAnswer
			}
			// Send the message.
			if _, err = bot.Send(msg); err != nil {
				panic(err)
			}
		} else if update.CallbackQuery != nil {
			callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
			if _, err := bot.Request(callback); err != nil {
				panic(err)
			}

			msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Data)
			if _, err := bot.Send(msg); err != nil {
				panic(err)
			}
		}
	}
}
