package main

import (
	"gotgbot/tg"
	"log"
	"time"
)

func main() {
	updateId := int64(0)

	for {
		log.Println(time.Now().Unix())
		u, err := tg.GetUpdates(updateId)
		if err != nil {
			log.Fatal(err)
		}
		if !u.Ok {
			log.Fatal(*u.Description, ": ", *u.ErrorCode)
		}

		if len(*u.Result) != 0 {
			for i := range *u.Result {
				update := (*u.Result)[i]
				updateId = update.Id + 1

				if update.Message != nil {
					tg.HandleMessage(update.Message)
				}
			}
		} else {
			log.Println("Updates is nil")
		}
	}
}
