package zerotier

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

func postMember(memberID string, data interface{}) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(data); err != nil {
		return err
	}
	_, err := request("POST", apiUpdateMembers(memberID), &buf)
	return err
}

func request(method, url string, body io.Reader) ([]byte, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "bearer "+token)
	var resp *http.Response
	if resp, err = http.DefaultClient.Do(req); err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

type Member struct {
	ID                  string        `json:"id"`
	Type                string        `json:"type"`
	Clock               int64         `json:"clock"`
	NetworkID           string        `json:"networkId"`
	NodeID              string        `json:"nodeId"`
	ControllerID        string        `json:"controllerId"`
	Hidden              bool          `json:"hidden"`
	Name                string        `json:"name"`
	Online              bool          `json:"online"`
	Description         string        `json:"description"`
	Config              *MemberConfig `json:"config"`
	LastOnline          int64         `json:"lastOnline"`
	PhysicalAddress     *string       `json:"physicalAddress"`
	PhysicalLocation    interface{}   `json:"physicalLocation"`
	ClientVersion       string        `json:"clientVersion"`
	ProtocolVersion     int           `json:"protocolVersion"`
	SupportsRulesEngine bool          `json:"supportsRulesEngine"`
}

type MemberConfig struct {
	ActiveBridge         bool          `json:"activeBridge"`
	Address              string        `json:"address"`
	Authorized           bool          `json:"authorized"`
	Capabilities         []interface{} `json:"capabilities"`
	CreationTime         int64         `json:"creationTime"`
	Id                   string        `json:"id"`
	Identity             string        `json:"identity"`
	IpAssignments        []string      `json:"ipAssignments"`
	LastAuthorizedTime   int64         `json:"lastAuthorizedTime"`
	LastDeauthorizedTime int64         `json:"lastDeauthorizedTime"`
	NoAutoAssignIps      bool          `json:"noAutoAssignIps"`
	Nwid                 string        `json:"nwid"`
	Objtype              string        `json:"objtype"`
	RemoteTraceLevel     int           `json:"remoteTraceLevel"`
	RemoteTraceTarget    string        `json:"remoteTraceTarget"`
	Revision             int           `json:"revision"`
	Tags                 []interface{} `json:"tags"`
	VMajor               int           `json:"vMajor"`
	VMinor               int           `json:"vMinor"`
	VRev                 int           `json:"vRev"`
	VProto               int           `json:"vProto"`
	SsoExempt            bool          `json:"ssoExempt"`
}

type RespStatus struct {
	ID          string      `json:"id"`
	Type        string      `json:"type"`
	Online      bool        `json:"online"`
	Clock       int64       `json:"clock"`
	Version     interface{} `json:"version"`
	ApiVersion  string      `json:"apiVersion"`
	Uptime      int64       `json:"uptime"`
	ClusterNode string      `json:"clusterNode"`
	User        *User       `json:"user"`
}

type User struct {
	ID           string `json:"id"`
	Type         string `json:"type"`
	CreationTime int64  `json:"creationTime"`
	DisplayName  string `json:"displayName"`
	Email        string `json:"email"`
}
