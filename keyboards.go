package main

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

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
