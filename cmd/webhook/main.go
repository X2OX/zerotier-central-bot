package main

import (
	"github.com/X2OX/zerotier-central-bot/telegram"
)

func main() {
	if err := telegram.Webhook("your url"); err != nil {
		panic(err)
	}
}
