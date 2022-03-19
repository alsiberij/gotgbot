package tg

import (
	"errors"
	"log"
	"strings"
	"sync"
)

type (
	UserChannels struct {
		mutex sync.Mutex
		data  map[int64]chan *Message
	}
)

var (
	channels = UserChannels{
		mutex: sync.Mutex{},
		data:  make(map[int64]chan *Message),
	}
)

func (c *UserChannels) createChan(chatId int64) chan *Message {
	c.mutex.Lock()
	v, ok := c.data[chatId]
	if ok {
		close(v)
	}
	ch := make(chan *Message)
	c.data[chatId] = ch
	c.mutex.Unlock()
	return ch
}

func (c *UserChannels) deleteChan(chatId int64) {
	c.mutex.Lock()
	v, ok := c.data[chatId]
	if ok {
		close(v)
		delete(c.data, chatId)
	}
	c.mutex.Unlock()
}

func (c *UserChannels) getChan(chatId int64) (chan *Message, bool) {
	c.mutex.Lock()
	result, ok := c.data[chatId]
	c.mutex.Unlock()
	return result, ok
}

func HandleMessage(msg *Message) {
	if msg == nil {
		return
	}

	entities := msg.Entities
	commandIndex := -1
	if entities != nil {
		for i := range *entities {
			if (*entities)[i].Type == "bot_command" {
				commandIndex = i
				break
			}
		}
	}

	if commandIndex != -1 {
		commandOffset := (*entities)[commandIndex].Offset
		commandLength := (*entities)[commandIndex].Length
		command := msg.Text[commandOffset : commandOffset+commandLength]

		commandArgs := strings.Split(command, "@")
		if len(commandArgs) == 2 {
			if commandArgs[1] != BotName {
				return
			}
			command = commandArgs[0]
		}

		switch command {
		case "/start":
			go helpCommand(msg)
		case "/help":
			go helpCommand(msg)
		case "/back":
			go backCommand(msg)
		case "/greet":
			go greetCommand(msg)
		}
	} else {
		v, ok := channels.getChan(msg.Chat.Id)
		if ok {
			v <- msg
			return
		}

		log.Println("MESSAGE: ", msg.Text)
	}
}

func greetCommand(start *Message) {
	ch := channels.createChan(start.Chat.Id)
	closed, err := GreetCommand(start, ch)
	if err != nil {
		log.Println("ERROR HANDLING /GREET: " + err.Error())
	}
	if !closed {
		channels.deleteChan(start.Chat.Id)
	}
}
func GreetCommand(start *Message, ch chan *Message) (bool, error) {
	if start.Chat.Type != "private" {
		_, err := SendMessage(start.Chat.Id, "Всем привет, меня зовут Максим Сергеевич!")
		if err != nil {
			return false, err
		}
		return false, nil
	}

	_, err := SendMessage(start.Chat.Id, "Как тебя зовут?")
	if err != nil {
		return false, err
	}

	var msg *Message
	var isOpened bool

	for {
		msg, isOpened = <-ch
		if !isOpened {
			return true, nil
		}
		if msg == nil {
			return false, errors.New("passed message is nil")
		}

		if msg.From.Id != start.From.Id {
			continue
		}

		if msg.Text != "Максим" {
			break
		} else {
			_, err = SendMessage(start.Chat.Id, "Нет, так зовут меня, а тебя как?")
			if err != nil {
				return false, err
			}
		}
	}

	_, err = SendMessage(start.Chat.Id, "Привет, "+msg.Text)
	if err != nil {
		return false, err
	}

	return false, nil
}

func helpCommand(start *Message) {
	_, err := HelpCommand(start, nil)
	if err != nil {
		log.Println("ERROR HANDLING /HELP or /START: " + err.Error())
	}
}
func HelpCommand(start *Message, ch chan *Message) (bool, error) {
	msgText := "Привет. Меня зовут Максим. Я много что умею! Например\n" +
		"/help - Расскажу о своих способностях\n" +
		"/greet - Поздороваюсь\n" +
		"/back - Перестану выполнять команду"
	_, err := SendMessage(start.Chat.Id, msgText)
	return false, err
}

func backCommand(start *Message) {
	channels.deleteChan(start.Chat.Id)
	_, err := BackCommand(start, nil)
	if err != nil {
		log.Println("ERROR HANDLING /HELP or /START: " + err.Error())
	}
}
func BackCommand(start *Message, ch chan *Message) (bool, error) {
	_, err := SendMessage(start.Chat.Id, "Окей, забыли")
	return false, err
}
