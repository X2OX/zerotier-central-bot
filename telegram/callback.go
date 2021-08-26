package telegram

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/FlowerLab/dandelion"
	"github.com/X2OX/zerotier-central-bot/zerotier"
)

type CallbackDataType uint8

const (
	CallbackNone CallbackDataType = iota
	CallbackTypeNodeInfoName
	CallbackTypeNodeInfoDescription
	CallbackTypeNodeInfoAddIP
	CallbackTypeNodeInfoSetIP
	CallbackTypeNodeInfoAuth
	CallbackTypeNodeInfoDeAuth
	CallbackTypeNodeInfoDelete
)

var (
	callbackDataTypes = []string{
		"None", "ModifyName", "ModifyDescription", "AddIP", "SetIP", "Auth", "DeAuth", "Delete",
	}
)

func (c CallbackDataType) String() string { return callbackDataTypes[c] }

type callbackData struct {
	Type  CallbackDataType `json:"t"`
	Param []string         `json:"p,omitempty"`
}

func (c callbackData) IsSetInfo() bool {
	return c.Type == CallbackTypeNodeInfoName || c.Type == CallbackTypeNodeInfoDescription ||
		c.Type == CallbackTypeNodeInfoAddIP || c.Type == CallbackTypeNodeInfoSetIP ||
		c.Type == CallbackTypeNodeInfoAuth || c.Type == CallbackTypeNodeInfoDeAuth ||
		c.Type == CallbackTypeNodeInfoDelete
}

func (c callbackData) ParamIs(i int) bool { return len(c.Param) >= i }

func parseCallbackData(s string) callbackData {
	var arg callbackData
	_ = json.Unmarshal([]byte(s), &arg)
	return arg
}

func NewCallbackData(t CallbackDataType, param ...string) *string {
	if b, err := json.Marshal(&callbackData{t, param}); err == nil && len(b) > 0 {
		s := string(b)
		return &s
	}
	return nil
}

type (
	Callback struct{}
)

func (Callback) Adapter() dandelion.Adapters { return nil }
func (Callback) IsMatch(c *dandelion.Context) bool {
	return c.Message.CallbackQuery != nil && parseCallbackData(c.Message.CallbackQuery.Data).Type != CallbackNone
}

func (Callback) Handle(c *dandelion.Context) bool {
	data := parseCallbackData(c.Message.CallbackQuery.Data)
	if data.IsSetInfo() {
		_ = setInfo(c, data)
	}

	return true
}

func setInfo(c *dandelion.Context, data callbackData) (err error) {
	if !data.ParamIs(1) {
		return errors.New("callback data missing parameters")
	}

	switch data.Type {
	case CallbackTypeNodeInfoName, CallbackTypeNodeInfoDescription, CallbackTypeNodeInfoAddIP, CallbackTypeNodeInfoSetIP:
		_, _ = c.Send(dandelion.MessageConfig{
			BaseChat: dandelion.BaseChat{
				ChatID:      c.Message.CallbackQuery.From.ID,
				ReplyMarkup: dandelion.ForceReply{ForceReply: true},
			},
			Text: fmt.Sprintf(`#%s #%s

Please reply parameters.

If you need to set the IP, please reply the IP address separated by spaces.
Example: <code>10.0.0.1 10.0.1.1 10.0.2.1</code>`, data.Param[0], data.Type.String()),
			ParseMode: dandelion.ModeHTML,
		})
	case CallbackTypeNodeInfoAuth:
		err = zerotier.Auth(data.Param[0], true)
		answerCallback(c.Message.CallbackQuery.ID, "Finish")
	case CallbackTypeNodeInfoDeAuth:
		err = zerotier.Auth(data.Param[0], false)
		answerCallback(c.Message.CallbackQuery.ID, "Finish")
	case CallbackTypeNodeInfoDelete:
		err = zerotier.Delete(data.Param[0])
		answerCallback(c.Message.CallbackQuery.ID, "Finish")
	}

	return
}

func answerCallback(callbackQueryID, text string) {
	_, _ = bot.Send(dandelion.CallbackConfig{
		CallbackQueryID: callbackQueryID,
		Text:            text,
	})
}
