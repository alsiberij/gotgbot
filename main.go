package main

import (
	"gobottg/tg"
	"log"
)

func main() {
	u, err := tg.GetUpdates(0)
	if err != nil {
		log.Fatal(err)
	}
	if !u.Ok {
		log.Fatal(*u.Description, ": ", *u.ErrorCode)
	}

	for i := range *u.Result {
		msg := (*u.Result)[i].Message
		if msg != nil {
			log.Printf("Новое сообщение от %s: %s\n",
				msg.From.FirstName,
				msg.Text)

			continue
		}

		mcm := (*u.Result)[i].MyChatMember
		if mcm != nil {
			if mcm.NewChatMember.Status == "kicked" {
				log.Printf("%s вышел", mcm.NewChatMember.User.FirstName)
			}
			if mcm.NewChatMember.Status == "member" {
				log.Printf("%s вошел", mcm.NewChatMember.User.FirstName)
			}
		}
	}
}
