package zerotier

import (
	"fmt"
	"net"
	"time"
)

func convMember(members []*Member) []*Node {
	arr := make([]*Node, 0, len(members))
	for _, v := range members {
		arg := &Node{
			NodeID:      v.NodeID,
			NetworkID:   v.NetworkID,
			Name:        v.Name,
			Description: v.Description,
			Version:     v.ClientVersion,
			Online:      v.Online,
			LastOnline:  time.Unix(v.LastOnline/1000, 0),
		}

		if v.PhysicalAddress != nil {
			arg.PhysicalAddress = net.ParseIP(*v.PhysicalAddress)
		}
		if v.Config != nil {
			arg.CreationTime = time.Unix(v.Config.CreationTime/1000, 0)
			arg.Authorized = v.Config.Authorized

			if v.Config.IpAssignments != nil {
				arg.IPAssignments = make([]net.IP, 0, len(v.Config.IpAssignments))
				for _, ip := range v.Config.IpAssignments {
					arg.IPAssignments = append(arg.IPAssignments, net.ParseIP(ip))
				}
			}
		}
		arr = append(arr, arg)
	}

	return arr
}

const (
	apiNetworkMemberList = "https://my.zerotier.com/api/v1/network/%s/member"
	apiNetworkMember     = "https://my.zerotier.com/api/v1/network/%s/member/%s"
	apiStatus            = "https://my.zerotier.com/api/v1/status"
)

func apiGetMembers() string {
	return fmt.Sprintf(apiNetworkMemberList, NetworkID)
}
func apiGetMember(id string) string {
	return fmt.Sprintf(apiNetworkMember, NetworkID, id)
}
func apiUpdateMembers(id string) string {
	return fmt.Sprintf(apiNetworkMember, NetworkID, id)
}
func apiDeleteMembers(id string) string {
	return fmt.Sprintf(apiNetworkMember, NetworkID, id)
}
