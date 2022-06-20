package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var comoDepositarShort string = "Tenemos muchos métodos de pago \v Link video \v Link video"

var comoDepositar string = "Tenemos muchos métodos de pago en línea o pago en efectivo muy cerca de ti ver nuestros métodos de recarga:\v Link video \v Link video"

var comoApostar string = "Acá te enseaños como apostar mira el video:\vLink video"

var bonoDebutShort string = "Esta es la info sobre bono debut"

var bonoDebut string = "El bono debut aplica para usuarios nuevos donde la primera recarga debe ser igual o mayor a $10.000.\vPara mayor información te invitamos a ingresar a los términos y condiciones en 👉🏼 https://www.wplay.co/mas/ayuda/200k-bono-debut-wplay/ 👍🏼"

var comoDescargarShort string = "Como descargar"

var comoDescargar string = "¿Cómo puedo descargar la App?\vPara descargar la App desde iOS ingresa aquí: https://apps.apple.com/co/app/id1490941885\vPara descargar la App desde Android ingresa aquí: https://play.google.com/store/apps/details?id=com.wplay.sports"

var tutorialesLink string = "Link a tutoriales"

var comoRetirar string = "Resumen sobre como retirar"

var numericInlineKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("¿Como depositar?", comoDepositarShort),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("¿Como apostar", comoApostar),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Bono Debut", bonoDebutShort),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("¿Como descargar la app?", comoDescargarShort),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Tutoriales", tutorialesLink),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("¿Como retirar mi saldo?", comoRetirar),
	),
)

var numericKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("¿Como depositar?"),
		tgbotapi.NewKeyboardButton("¿Como apostar"),
		tgbotapi.NewKeyboardButton("Bono Debut"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("¿Como descargar la app?"),
		tgbotapi.NewKeyboardButton("Tutoriales"),
		tgbotapi.NewKeyboardButton("¿Como retirar mi saldo?"),
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
			case "¿Como depositar?":
				msg.Text = comoDepositar
			case "¿Como apostar":
				msg.Text = comoApostar
			case "Bono Debut":
				msg.Text = bonoDebut
			case "¿Como descargar la app?":
				msg.Text = comoDescargar
			case "Tutoriales":
				msg.Text = tutorialesLink
			case "¿Como retirar mi saldo?":
				msg.Text = comoRetirar
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
