package telegram

import (
	"net"
	"strings"

	"github.com/FlowerLab/dandelion"
	"github.com/X2OX/zerotier-central-bot/zerotier"
)

type Message struct{}

func (i *Message) Adapter() dandelion.Adapters { return nil }
func (i *Message) IsMatch(c *dandelion.Context) bool {
	return c.Message.Message != nil && !c.Message.Message.IsCommand()
}
func (i *Message) Handle(c *dandelion.Context) bool {
	if c.Message.Message.ReplyToMessage == nil {
		return false
	}
	setNodeInfo(c)
	return true
}

func setNodeInfo(c *dandelion.Context) {
	arr := strings.SplitN(c.Message.Message.ReplyToMessage.Text, "\n", 2)
	if len(arr) < 1 {
		return
	}

	arr = strings.Split(strings.ReplaceAll(arr[0], " ", ""), "#")
	if len(arr) != 3 {
		return
	}
	switch arr[2] {
	case CallbackTypeNodeInfoName.String():
		_ = zerotier.SetName(arr[1], c.Message.Message.Text)
	case CallbackTypeNodeInfoDescription.String():
		_ = zerotier.SetDescription(arr[1], c.Message.Message.Text)
	case CallbackTypeNodeInfoAddIP.String():
		ip := net.ParseIP(c.Message.Message.Text)
		if len(ip) == 0 {
			return
		}
		node := zerotier.Get(arr[1])
		node.IPAssignments = append(node.IPAssignments, ip)
		_ = zerotier.SetIP(arr[1], node.IPAssignments)
	case CallbackTypeNodeInfoSetIP.String():
		var ips []net.IP
		for _, v := range strings.Split(c.Message.Message.Text, " ") {
			if ip := net.ParseIP(v); len(ip) != 0 {
				ips = append(ips, ip)
			}
		}
		_ = zerotier.SetIP(arr[1], ips)
	}
	_, _ = bot.Send(dandelion.DeleteMessageConfig{
		ChatID:    c.Message.Message.ReplyToMessage.From.ID,
		MessageID: c.Message.Message.ReplyToMessage.MessageID,
	})
	_, _ = bot.Send(dandelion.MessageConfig{
		BaseChat: dandelion.BaseChat{
			ChatID:           c.Message.Message.From.ID,
			ReplyToMessageID: c.Message.Message.MessageID,
		},
		Text:      renderNode(zerotier.Get(arr[1])),
		ParseMode: dandelion.ModeHTML,
	})
}
