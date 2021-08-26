package timer

import (
	"net/http"
	"os"
	"sync"

	"github.com/X2OX/zerotier-central-bot/database"
	"github.com/X2OX/zerotier-central-bot/telegram"
	"github.com/X2OX/zerotier-central-bot/zerotier"
	_ "go.x2ox.com/utils/timezone"
)

var (
	token = os.Getenv("TIMER_TOKEN")
	wg    sync.WaitGroup
)

func Handler(w http.ResponseWriter, r *http.Request) {
	if token == "" || r.URL.Query().Get("token") != token || !database.IsEnable() {
		w.WriteHeader(http.StatusTeapot)
		return
	}

	for _, v := range zerotier.List() {
		node := v
		wg.Add(1)
		go func() { // avoid serverless timeouts
			dbNode := database.GetNode(node.NodeID)

			switch {
			case !node.Authorized && dbNode != nil && !dbNode.Authorized: // not verified, is notified
			case node.Authorized && dbNode != nil && !dbNode.Authorized: // verification complete
				database.DeleteNode(node.NodeID)
			case !node.Authorized: // not verified, add to database and notify
				telegram.SendAuth(node.NodeID)
				database.AddNode(node.NodeID, node)
			case node.Online && dbNode != nil && !dbNode.Online: // offline device just went online, go to notify
				telegram.SendOnline(node.NodeID)
				database.DeleteNode(node.NodeID)
			case !node.Online && dbNode != nil && !dbNode.Online: // offline device is still offline
			case !node.Online: // device is offline, go to notify
				telegram.SendOffline(node.NodeID)
				database.AddNode(node.NodeID, node)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
