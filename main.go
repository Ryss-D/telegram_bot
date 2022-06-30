package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var genericMessage = "¡Bienvenidos al Canal más POWEROSO de Telegram! 🔥⚡\n Wplay.co te estará acompañando de formas interactivas para que te conviertas en un crack de las apuestas ⚽🃏 \n ¿Qué deseas aprender?"

type doubt struct {
	question    string
	fullAnswer  string
	shortAnswer string
	link        string
}

var comoDepositar = doubt{
	question:    "¿Cómo depositar",
	fullAnswer:  "Tenemos muchos métodos de pago en línea o pago en efectivo muy cerca de ti ver nuestros métodos de recarga: \n https://www.youtube.com/watch?v=cwVw3B0fh2Y&list=PLyxCM60Pf1ztnA9sf7AHZrcQQBrE6Zcfl&index=2 \n https://www.wplay.co/mas/ayuda/depositos-y-retiros/",
	shortAnswer: "Tenemos muchos métodos de pago \v Link video \v Link video",
	link:        "link",
}

var comoApostar = doubt{
	question:    "¿Cómo apostar",
	fullAnswer:  "Acá te enseaños como apostar mira el video: \n https://www.youtube.com/watch?v=cnYbcq8Q7Bk&list=PLyxCM60Pf1ztnA9sf7AHZrcQQBrE6Zcfl \n https://www.wplay.co/mas/ayuda/como-apostar-en-wplay-co/",
	shortAnswer: "Acá te enseñamos como apostar mira el video",
	link:        "link",
}

var bonoDebut = doubt{
	question:    "Bono Debut",
	fullAnswer:  "El bono debut aplica para usuarios nuevos donde la primera recarga debe ser igual o mayor a $10.000. \n Para mayor información te invitamos a ingresar a los términos y condiciones en 👉🏼 https://www.wplay.co/mas/ayuda/200k-bono-debut-wplay/ 👍🏼",
	shortAnswer: "Esta es la info sobre bono debut",
	link:        "link",
}

var comoDescargar = doubt{
	question:    "¿Cómo descargar",
	fullAnswer:  "¿Cómo puedo descargar la App? \n Para descargar la App desde iOS ingresa aquí: https://apps.apple.com/co/app/id1490941885 \n Para descargar la App desde Android ingresa aquí: https://play.google.com/store/apps/details?id=com.wplay.sports",
	shortAnswer: "Como descargar",
	link:        "link",
}

var tutoriales = doubt{
	question:    "Tutoriales",
	fullAnswer:  "En wplay.co te brindamos varios tutoriales para el manejo de la plataforma \n Te invitamos a ingresar a https://www.wplay.co/mas/tutoriales-wplay/ donde podrás visualizar la información.",
	shortAnswer: "Tutoriales",
	link:        "link",
}

var comoRetirar = doubt{
	question:    "¿Cómo retirar",
	fullAnswer:  "En Wplay.co contamos gran variedad de medios de retiro \n Te invitamos a ingresar https://blog.wplay.co/retiro-de-tus-ganancias-en-wplay-co/ donde podrás visualizar los puntos autorizados para retirar y los respectivos pasos en cada uno de ellos.",
	shortAnswer: "Resumen de como retirar",
	link:        "link",
}
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
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, genericMessage)
			msg.ReplyMarkup = numericKeyboard

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
