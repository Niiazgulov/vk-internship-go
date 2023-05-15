package client

type Update struct {
	UpdateID int     `json:"update_id"`
	Message  Message `json:"message"`
}

type Message struct {
	Text string `json:"text"`
	Chat Chat   `json:"chat"`
}

type Chat struct {
	ChatID int `json:"id"`
}

type UpdatesResp struct {
	Result []Update `json:"result"`
}

type ResponseMessage struct {
	ChatID int                 `json:"chat_id"`
	Text   string              `json:"text"`
	Reply  ReplyKeyboardMarkup `json:"reply_markup"`
}

type ReplyKeyboardMarkup struct {
	Keyboard [][]KeyboardButton `json:"keyboard"`
}

type KeyboardButton struct {
	Text string `json:"text"`
}
