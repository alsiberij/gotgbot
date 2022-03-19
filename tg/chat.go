package tg

type (
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
