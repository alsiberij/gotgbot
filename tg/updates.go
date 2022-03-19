package tg

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
	"os"
	"strconv"
	"time"
)

type (
	Updates struct {
		Ok          bool      `json:"ok"`
		Result      *[]Update `json:"result"`
		ErrorCode   *int      `json:"error_code"`
		Description *string   `json:"description"`
	}
	Update struct {
		Id           int64         `json:"update_id"`
		Message      *Message      `json:"message"`
		MyChatMember *MyChatMember `json:"my_chat_member"`
	}
)

func GetUpdates(sinceUpdateId int64) (Updates, error) {
	var updates Updates

	rq := fasthttp.AcquireRequest()
	rs := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(rq)
	defer fasthttp.ReleaseResponse(rs)

	rq.SetRequestURI(DefaultUri + "/getUpdates")
	rq.Header.SetContentType("application/application/x-www-form-urlencoded")
	rq.Header.SetMethod("POST")

	if sinceUpdateId != 0 {
		rq.PostArgs().Add("offset", strconv.FormatInt(sinceUpdateId, 10))
	}

	rq.PostArgs().Add("timeout", "60")

	err := ApiClient.Do(rq, rs)
	if err != nil {
		return updates, err
	}

	f, _ := os.Create("tg/logs/getUpdates-log" + time.Now().Format("2006-01-02---15-04") + ".json")
	_, _ = f.Write(rs.Body())
	_ = f.Close()

	err = json.Unmarshal(rs.Body(), &updates)
	if err != nil {
		return updates, err
	}

	return updates, nil
}
