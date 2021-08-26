package zerotier

import (
	"net"
	"net/http"
	"strings"

	"github.com/X2OX/zerotier-central-bot/telegram"
	_ "go.x2ox.com/utils/timezone"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	if !checkIP(r) {
		w.WriteHeader(http.StatusTeapot)
		return
	}
	telegram.Handel(r)
}

func checkIP(r *http.Request) bool {
	_, in1, _ := net.ParseCIDR("149.154.160.0/20")
	_, in2, _ := net.ParseCIDR("91.108.4.0/22")
	ip := net.ParseIP(getRemoteAddr(r))
	return in1.Contains(ip) || in2.Contains(ip)
}

func getRemoteAddr(r *http.Request) string {
	if ip := strings.TrimSpace(strings.Split(r.Header.Get("X-Forwarded-For"), ",")[0]); ip != "" {
		return ip
	}
	if ip := strings.TrimSpace(r.Header.Get("X-Real-Ip")); ip != "" {
		return ip
	}
	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
	}
	return ""
}
