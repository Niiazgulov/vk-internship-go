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
	ChatID int      `json:"chat_id"`
	Text   string   `json:"text"`
	Reply  RKMarkup `json:"reply_markup"`
}

type RKMarkup struct {
	Keyboard [][]Button `json:"keyboard"`
}

type Button struct {
	Text string `json:"text"`
}
