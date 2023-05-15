package client

import (
	// "bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"strconv"

	"github.com/Niiazgulov/vk-internship-go.git/internal/config"
)

var (
	tlgAPI = "api.telegram.org"
	token  = "bot" + config.GetTokenFile()
	urlReq = url.URL{Scheme: "https", Host: tlgAPI}
	client = &http.Client{}
)

func makeRequest(method string, query url.Values) []byte {
	urlReq.Path = path.Join(token, method)
	request, err := http.NewRequest(http.MethodGet, urlReq.String(), nil)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	request.URL.RawQuery = query.Encode()
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return body
}

func GetUpdates(offset int) ([]Update, error) {
	q := url.Values{}
	q.Add("offset", strconv.Itoa(offset))
	q.Add("limit", strconv.Itoa(50))
	data := makeRequest("getUpdates", q)
	var res UpdatesResp
	if err := json.Unmarshal(data, &res); err != nil {
		return nil, err
	}

	return res.Result, nil
}

func SendMessages(upd Update) error {
	var msg ResponseMessage
	msg.ChatID = upd.Message.Chat.ChatID
	urlReq.Path = path.Join(token, "sendMessage")

	switch upd.Message.Text {
	case cmdStart:
		msg.Text = msgStart
	case cmdHumour:
		msg.Text = msgSure
	case cmdYes:
		msg.Text = msgReadySad
	case cmdNo:
		msg.Text = msgReadyPositive
	case cmdAsk:
		msg.Text = msgAsk
	case cmdOtvet1:
		msg.Text = msgAnsw
	case cmdOtvet2:
		msg.Text = msgAnsw2
	case cmdHelp:
		msg.Text = msgHelp
	case cmdCallMe:
		msg.Text = msgHelpPos
	case cmdForgetMe:
		msg.Text = msgHelpNeg
	case cmdAnekdot:
		msg.Text = msgAnekdot
	case cmdAnekdot2:
		msg.Text = msgAnekdot2
	case cmdAnekdot3:
		msg.Text = msgAnekdot3
	case cmdAnekdot4:
		msg.Text = msgAnekdot4
	case "K":
		msg.Text = "No probLama"
		msg.Reply.Keyboard = [][]KeyboardButton{{{Text: "Кнопка1"}, {Text: "Кнопка2"}}, {{Text: "2Кнопка3"}, {Text: "2Кнопка4"}}}
	default:
		msg.Text = msgDefault
	}

	// buffMsg, err := json.Marshal(msg)
	// if err != nil {
	// 	return err
	// }
	// _, err = http.Post(urlReq.String(), "application/json", bytes.NewBuffer(buffMsg))
	// if err != nil {
	// 	return err
	// }

	q := url.Values{}
	q.Add("chat_id", strconv.Itoa(msg.ChatID))
	q.Add("text", msg.Text)
	makeRequest("sendMessage", q)

	return nil
}
