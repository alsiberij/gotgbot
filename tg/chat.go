package tg

type (
	Message struct {
		Id       int64
		From     *UserShort       `json:"from"`
		Chat     *Chat            `json:"chat"`
		Date     int64            `json:"date"`
		Photo    *[]Photo         `json:"photo"`
		Text     string           `json:"text"`
		Entities *[]MessageEntity `json:"entities"`
	}
	UserShort struct {
		Id           int64  `json:"id"`
		IsBot        bool   `json:"is_bot"`
		FirstName    string `json:"first_name"`
		Username     string `json:"username"`
		LanguageCode string `json:"language_code"`
	}
	Chat struct {
		Id        int64  `json:"id"`
		FirstName string `json:"first_name"`
		Username  string `json:"username"`
		Type      string `json:"type"`
	}
	MessageEntity struct {
		Offset int64  `json:"offset"`
		Length int64  `json:"length"`
		Type   string `json:"type"`
	}

	MyChatMember struct {
		Chat          Chat       `json:"chat"`
		From          UserShort  `json:"from"`
		Date          int64      `json:"date"`
		OldChatMember ChatMember `json:"old_chat_member"`
		NewChatMember ChatMember `json:"new_chat_member"`
	}
	ChatMember struct {
		User      UserShort `json:"user"`
		Status    string    `json:"status"`
		UntilDate int64     `json:"untilDate"`
	}
)
