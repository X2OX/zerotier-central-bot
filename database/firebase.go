package database

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path"

	"github.com/X2OX/zerotier-central-bot/zerotier"
)

var (
	databaseURL string
	isEnable    bool
	client      = &http.Client{}
)

func init() {
	u, err := url.Parse(os.Getenv("FIREBASE_URL"))
	if err != nil {
		return
	}
	u.Path = path.Join(u.Path, "/node")
	databaseURL = u.String()
	isEnable = true
}

func IsEnable() bool { return isEnable }

func AddNode(nodeID string, node *zerotier.Node) {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(node); err != nil {
		return
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/%s.json", databaseURL, nodeID), &buf)
	if err != nil {
		return
	}

	var resp *http.Response
	if resp, err = client.Do(req); err != nil {
		return
	}
	_ = resp.Body.Close()
}

func ListNode() []*zerotier.Node {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s.json", databaseURL), nil)
	if err != nil {
		return nil
	}

	var resp *http.Response
	if resp, err = client.Do(req); err != nil || resp.StatusCode != 200 {
		return nil
	}
	var arg map[string]*zerotier.Node
	if err = json.NewDecoder(resp.Body).Decode(&arg); err != nil {
		return nil
	}

	arr := make([]*zerotier.Node, 0, len(arg))
	for _, v := range arg {
		arr = append(arr, v)
	}

	return arr
}

func GetNode(nodeID string) *zerotier.Node {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s.json", databaseURL, nodeID), nil)
	if err != nil {
		return nil
	}

	var resp *http.Response
	if resp, err = client.Do(req); err != nil || resp.StatusCode != 200 {
		return nil
	}
	var arg zerotier.Node
	if err = json.NewDecoder(resp.Body).Decode(&arg); err != nil || arg.NodeID == "" {
		return nil
	}
	return &arg
}

func DeleteNode(nodeID string) {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/%s.json", databaseURL, nodeID), nil)
	if err != nil {
		return
	}

	var resp *http.Response
	if resp, err = client.Do(req); err != nil || resp.StatusCode != 200 {
		return
	}
	_ = resp.Body.Close()
}

func testServer() error {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s.json", databaseURL), nil)
	if err != nil {
		return err
	}

	var resp *http.Response
	if resp, err = client.Do(req); err != nil {
		return err
	}
	_ = resp.Body.Close()
	if resp.StatusCode != 200 {
		return fmt.Errorf("%s | reference: https://firebase.google.com/docs/database/rest/save-data#section-rest-errors", resp.Status)
	}
	return nil
}

func Status() bool { return testServer() == nil }
