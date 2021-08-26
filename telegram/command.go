package telegram

import (
	"fmt"

	"github.com/FlowerLab/dandelion"
	"github.com/X2OX/zerotier-central-bot/database"
	"github.com/X2OX/zerotier-central-bot/zerotier"
)

type (
	Command      struct{}
	CommandStart struct{}
	CommandInfo  struct{}
)

func (Command) Adapter() dandelion.Adapters { return nil }
func (Command) IsMatch(c *dandelion.Context) bool {
	return c.Message.Message != nil && c.Message.Message.IsCommand()
}
func (Command) Handle(c *dandelion.Context) bool {
	switch c.Message.Message.Command() {
	case "start":
		cmdStart(c, c.Message.Message.CommandArguments())
	case "info":
		cmdInfo(c)
	case "x2ox":
		cmdSetCommand(c)
	default:
		c.SendText("Please enter the correct command")
	}
	return true
}

func cmdStart(c *dandelion.Context, arg string) {
	if arg == "" { // 只看未经验证，只看离线，查看全部
		_, _ = c.Send(dandelion.MessageConfig{
			BaseChat: dandelion.BaseChat{
				ChatID: c.Message.Message.From.ID,
			},
			Text:                  renderNodes(zerotier.List()),
			ParseMode:             dandelion.ModeHTML,
			DisableWebPagePreview: true,
		})
		return
	}

	node := zerotier.Get(arg)
	_, _ = c.Send(dandelion.MessageConfig{
		BaseChat: dandelion.BaseChat{
			ChatID: c.Message.Message.From.ID,
			ReplyMarkup: &dandelion.InlineKeyboardMarkup{InlineKeyboard: [][]dandelion.InlineKeyboardButton{
				{
					dandelion.InlineKeyboardButton{Text: "ModifyName",
						CallbackData: NewCallbackData(CallbackTypeNodeInfoName, node.NodeID)},
					dandelion.InlineKeyboardButton{Text: "ModifyDescription",
						CallbackData: NewCallbackData(CallbackTypeNodeInfoDescription, node.NodeID)},
				},
				{
					dandelion.InlineKeyboardButton{Text: "AddIP",
						CallbackData: NewCallbackData(CallbackTypeNodeInfoAddIP, node.NodeID)},
					dandelion.InlineKeyboardButton{Text: "SetIP",
						CallbackData: NewCallbackData(CallbackTypeNodeInfoSetIP, node.NodeID)},
				},
				{
					dandelion.InlineKeyboardButton{Text: "Auth",
						CallbackData: NewCallbackData(CallbackTypeNodeInfoAuth, node.NodeID)},
					dandelion.InlineKeyboardButton{Text: "DeAuth",
						CallbackData: NewCallbackData(CallbackTypeNodeInfoDeAuth, node.NodeID)},
				},
				{
					dandelion.InlineKeyboardButton{Text: "Delete",
						CallbackData: NewCallbackData(CallbackTypeNodeInfoDelete, node.NodeID)},
				},
			},
			}},
		Text:                  renderNode(node),
		ParseMode:             dandelion.ModeHTML,
		DisableWebPagePreview: true,
	})
}

func cmdInfo(c *dandelion.Context) {
	_, _ = c.Send(dandelion.MessageConfig{
		BaseChat: dandelion.BaseChat{
			ChatID: c.Message.Message.From.ID,
		},
		Text: fmt.Sprintf(`<b>Servre Info</b>

TelegramID: <code>%d</code>
ZeroTier: %t
Firebase: %t`, c.Message.Message.From.ID, zerotier.Status(), database.Status()),
		ParseMode:             dandelion.ModeHTML,
		DisableWebPagePreview: true,
	})
}

func cmdSetCommand(c *dandelion.Context) {
	_, _ = c.Send(dandelion.SetMyCommandsConfig{}.Set([]dandelion.BotCommand{
		{Command: "start", Description: "「 Node info 」"},
		{Command: "info", Description: "「 Service info 」"},
	}))
}
