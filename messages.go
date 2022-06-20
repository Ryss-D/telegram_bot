package main

var bonoDebutShort string = "Esta es la info sobre bono debut"

var bonoDebut string = "El bono debut aplica para usuarios nuevos donde la primera recarga debe ser igual o mayor a $10.000.\vPara mayor información te invitamos a ingresar a los términos y condiciones en 👉🏼 https://www.wplay.co/mas/ayuda/200k-bono-debut-wplay/ 👍🏼"

var comoDescargarShort string = "Como descargar"

var comoDescargar string = "¿Cómo puedo descargar la App?\vPara descargar la App desde iOS ingresa aquí: https://apps.apple.com/co/app/id1490941885\vPara descargar la App desde Android ingresa aquí: https://play.google.com/store/apps/details?id=com.wplay.sports"

var tutorialesLink string = "Link a tutoriales"

var comoRetirar string = "Resumen sobre como retirar"

type doubt struct {
	question    string
	fullAnswer  string
	shortAnswer string
	link        string
}

var comoDeposistar = doubt{
	question:    "¿Cómo depositar",
	fullAnswer:  "Tenemos muchos métodos de pago en línea o pago en efectivo muy cerca de ti ver nuestros métodos de recarga:\v Link video \v Link video",
	shortAnswer: "Tenemos muchos métodos de pago \v Link video \v Link video",
	link:        "link",
}

var comoApostar = doubt{
	question:    "¿Cómo apostar",
	fullAnswer:  "Acá te enseñamos como apostar mira el video",
	shortAnswer: "Acá te enseñamos como apostar mira el video",
	link:        "link",
}
