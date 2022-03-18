package tg

import "log"

//Greet
func Greet(chatId int64, username string) {
	_, err := SendMessage(chatId, "Привет, "+username+"!")
	if err != nil {
		log.Println(err)
	}
}
