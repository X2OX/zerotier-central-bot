package main

import (
	"log"
	"net/http"

	"github.com/X2OX/zerotier-central-bot/api/timer"
	"github.com/X2OX/zerotier-central-bot/api/zerotier"
	_ "go.x2ox.com/utils/timezone"
)

func main() {
	http.HandleFunc("/api/timer", timer.Handler)
	http.HandleFunc("/api/zerotier", zerotier.Handler)
	if err := http.ListenAndServe(":8888", nil); err != nil {
		log.Fatalf("http listen failed: %s", err)
	}
}
