package tg

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
	"os"
	"strconv"
	"time"
)

func SendMessage(chatId int64, text string) (Message, error) {
	var message Message

	rq := fasthttp.AcquireRequest()
	rs := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(rq)
	defer fasthttp.ReleaseResponse(rs)

	rq.SetRequestURI(DefaultUri + "/sendMessage")
	rq.Header.SetContentType("application/application/x-www-form-urlencoded")
	rq.Header.SetMethod("POST")

	rq.PostArgs().Add("chat_id", strconv.FormatInt(chatId, 10))
	rq.PostArgs().Add("text", text)

	err := ApiClient.Do(rq, rs)
	if err != nil {
		return message, err
	}

	f, _ := os.Create("tg/logs/sendMessage-log" + time.Now().Format("2006-01-02---15-04") + ".json")
	_, _ = f.Write(rs.Body())
	_ = f.Close()

	err = json.Unmarshal(rs.Body(), &message)
	if err != nil {
		return message, err
	}

	return message, nil
}
