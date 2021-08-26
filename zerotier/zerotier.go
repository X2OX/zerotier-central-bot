package zerotier

import (
	"encoding/json"
	"net"
	"os"
	"time"
)

var (
	token     = os.Getenv("ZEROTIER_CENTRAL_TOKEN")
	NetworkID = os.Getenv("ZEROTIER_NETWORK_ID")
)

type Node struct {
	NodeID          string    `json:"node_id"`          // 节点 ID
	NetworkID       string    `json:"network_id"`       // network id
	Name            string    `json:"name"`             // 名称
	Description     string    `json:"description"`      // 描述
	Version         string    `json:"version"`          // 客户端版本
	Online          bool      `json:"online"`           // 是否在线
	CreationTime    time.Time `json:"creation_time"`    // 首次连接时间
	LastOnline      time.Time `json:"last_online"`      // 最后在线时间
	IPAssignments   []net.IP  `json:"ip_assignments"`   // 客户端 IP
	Authorized      bool      `json:"authorized"`       // 是否验证
	PhysicalAddress net.IP    `json:"physical_address"` // 物理 IP
}

func List() []*Node {
	bts, err := request("GET", apiGetMembers(), nil)
	if err != nil {
		return nil
	}
	var arg []*Member
	if err = json.Unmarshal(bts, &arg); err != nil {
		return nil
	}
	return convMember(arg)
}

func Get(nodeID string) *Node {
	bts, err := request("GET", apiGetMember(nodeID), nil)
	if err != nil {
		return nil
	}
	var arg Member
	if err = json.Unmarshal(bts, &arg); err != nil {
		return nil
	}
	return convMember([]*Member{&arg})[0]
}

func SetName(nodeID, name string) error {
	return postMember(nodeID, map[string]string{
		"name": name,
	})
}

func SetDescription(nodeID, description string) error {
	return postMember(nodeID, map[string]string{
		"description": description,
	})
}

func SetIP(nodeID string, ips []net.IP) error {
	return postMember(nodeID, map[string]interface{}{
		"config": map[string]interface{}{
			"ipAssignments": ips,
		},
	})
}

func Auth(nodeID string, ok bool) error {
	return postMember(nodeID, map[string]interface{}{
		"config": map[string]interface{}{
			"authorized": ok,
		},
	})
}

func Delete(nodeID string) error {
	_, err := request("DELETE", apiDeleteMembers(nodeID), nil)
	return err
}

func Status() bool {
	bts, err := request("GET", apiStatus, nil)
	if err != nil {
		return false
	}
	var arg RespStatus
	if err = json.Unmarshal(bts, &arg); err != nil {
		return false
	}
	return arg.User != nil
}
