package telegram

import (
	"errors"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"github.com/FlowerLab/dandelion"
	"github.com/X2OX/zerotier-central-bot/zerotier"
)

var (
	telegramID int64
	bot        *dandelion.Engine
	adapters   dandelion.Adapters
)

func init() {
	var err error
	if bot, err = dandelion.New(os.Getenv("TELEGRAM_BOT_TOKEN")); err != nil {
		panic(err)
	}
	telegramID, _ = strconv.ParseInt(os.Getenv("TELEGRAM_ID"), 10, 64)
	adapters = dandelion.Adapters{&Auth{}}
	bot.SetAdapter(&Auth{})
}

func Run() {
	bot.Run()
}

func SendOffline(nodeID string) { _ = sendNotify(nodeID, "Offline") }
func SendOnline(nodeID string)  { _ = sendNotify(nodeID, "Online") }
func SendAuth(nodeID string)    { _ = sendNotify(nodeID, "Need Verification") }
func sendNotify(nodeID, content string) error {
	if telegramID == 0 {
		return errors.New("telegram user is not specified")
	}
	_, err := bot.Send(dandelion.MessageConfig{
		BaseChat: dandelion.BaseChat{
			ChatID: telegramID,
		},
		Text:      renderNodeNotify(zerotier.Get(nodeID), content),
		ParseMode: dandelion.ModeHTML,
	})
	return err
}

func Handel(r *http.Request) {
	update, err := bot.HandleUpdate(r)
	if err != nil {
		return
	}

	adapters.Match(&dandelion.Context{
		Engine:  bot,
		Message: *update,
	})
}

type Auth struct{}

func (Auth) Adapter() dandelion.Adapters {
	return dandelion.Adapters{&Command{}, &Message{}, &Callback{}}
}
func (Auth) IsMatch(_ *dandelion.Context) bool { return true }
func (Auth) Handle(c *dandelion.Context) bool {
	if c == nil {
		return true
	}
	return !(telegramID == 0 ||
		(c.Message.Message != nil && c.Message.Message.From.ID == telegramID) ||
		(c.Message.CallbackQuery != nil && c.Message.CallbackQuery.From.ID == telegramID))
}

func Webhook(_url string) error {
	u, err := url.Parse(_url)
	if err != nil {
		return err
	}
	_, err = bot.Request(dandelion.WebhookConfig{URL: u})
	return err
}
