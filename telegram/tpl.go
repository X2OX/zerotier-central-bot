package telegram

import (
	"bytes"
	"html/template"

	"github.com/X2OX/zerotier-central-bot/zerotier"
)

func renderNodes(list []*zerotier.Node) string {
	if len(list) == 0 {
		return ""
	}
	var buf bytes.Buffer
	if err := template.Must(template.New("tplNodes").Funcs(map[string]interface{}{
		"NetworkID":   func() string { return list[0].NetworkID },
		"BotUsername": func() string { return bot.Username() },
	}).Parse(tplNodes)).Execute(&buf, list); err != nil {
		return ""
	}
	return buf.String()
}

func renderNode(n *zerotier.Node) string {
	var buf bytes.Buffer
	if err := template.Must(template.New("tplNode").Funcs(map[string]interface{}{
		"BotUsername": func() string { return bot.Username() },
	}).Parse(tplNode)).Execute(&buf, n); err != nil {
		return ""
	}
	return buf.String()
}
func renderNodeNotify(n *zerotier.Node, content string) string {
	var buf bytes.Buffer
	if err := template.Must(template.New("tplNodeNotify").Funcs(map[string]interface{}{
		"BotUsername": func() string { return bot.Username() },
		"Content":     func() string { return content },
	}).Parse(tplNodeNotify)).Execute(&buf, n); err != nil {
		return ""
	}
	return buf.String()
}

const (
	tplNodes = `<b>ZeroTier</b> #{{NetworkID}} 

{{range .}}#{{.NodeID}} <a href="https://t.me/{{BotUsername}}?start={{.NodeID}}">Click for details</a> {{if .Authorized}}
{{.Name}} {{if .Online}}🖇{{else}}📎{{end}} <code>{{.PhysicalAddress}}</code> {{if .Online}}{{else}}| {{.LastOnline}}{{end}}
IP: {{range .IPAssignments}}<code>{{.}}</code> {{end}}
{{else}}🔐
{{end}}
{{end}}`

	tplNode = `#{{.NodeID}}

{{.Name}} - {{if .Online}}Online{{else}}Offline{{end}}{{if .Authorized}}{{else}} | 🔐{{end}}

<b>Version</b>: {{.Version}}
<b>CreationTime</b>: {{.CreationTime}}
<b>LastOnline</b>: {{.LastOnline}}
<b>PhysicalAddress</b>: {{.PhysicalAddress}}
<b>IPAssignments</b>: {{range .IPAssignments}}<code>{{.}}</code> {{end}}

{{.Description}}`

	tplNodeNotify = `#Notify {{Content}}

#{{.NodeID}}

{{.Name}} - {{if .Online}}Online{{else}}Offline{{end}}{{if .Authorized}}{{else}} | 🔐{{end}}

<b>Version</b>: {{.Version}}
<b>CreationTime</b>: {{.CreationTime}}
<b>LastOnline</b>: {{.LastOnline}}
<b>PhysicalAddress</b>: {{.PhysicalAddress}}
<b>IPAssignments</b>: {{range .IPAssignments}}<code>{{.}}</code> {{end}}

{{.Description}}`
)
