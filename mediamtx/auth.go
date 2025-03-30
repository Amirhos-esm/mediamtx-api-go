package mediamtx

import (
	"encoding/json"
	"fmt"
)

type Action int

const (
	ActionPublish Action = iota + 1
	ActionRead
	ActionPlayback
	ActionAPI
	ActionMetrics
	ActionPprof
)

var ActionStateStringHashMAP = map[Action]string{
	ActionPublish:  "publish",
	ActionRead:     "read",
	ActionPlayback: "playback",
	ActionAPI:      "api",
	ActionMetrics:  "metrics",
	ActionPprof:    "pprof",
}

func (a Action) String() string {
	return ActionStateStringHashMAP[a]
}
func (a *Action) FromString(str string) error {
	for k, v := range ActionStateStringHashMAP {
		if v == str {
			*a = k // Correctly set the value using the pointer receiver
			return nil
		}
	}
	return fmt.Errorf("invalid Action '%s'", str)
}

/* Custom Unmarshaller for Action */
func (a *Action) UnmarshalJSON(data []byte) error {
	var statusStr string
	if err := json.Unmarshal(data, &statusStr); err != nil {
		return err
	}
	if err := a.FromString(statusStr); err != nil {
		return err
	}
	return nil
}

/* Custom Marshaller for Status */
func (a Action) MarshalJSON() ([]byte, error) {
	ret := a.String()
	if ret == "" {
		return nil, fmt.Errorf("bad Action data `%v`", int(a))
	}
	return json.Marshal(ret)
}

type Protocol int

const (
	ProtocolRTSP Protocol = iota + 1
	ProtocolRTMP
	ProtocolHLS
	ProtocolWebRTC
	ProtocolSRT
)

var ProtocolStringHashMAP = map[Protocol]string{
	ProtocolRTSP:   "rtsp",
	ProtocolRTMP:   "rtmp",
	ProtocolHLS:    "hls",
	ProtocolWebRTC: "webrtc",
	ProtocolSRT:    "srt",
}

func (a Protocol) String() string {
	return ProtocolStringHashMAP[a]
}

func (a *Protocol) FromString(str string) error {
	for k, v := range ProtocolStringHashMAP {
		if v == str {
			*a = k // Correctly set the value using the pointer receiver
			return nil
		}
	}
	return fmt.Errorf("invalid Protocol '%s'", str)
}

/* Custom Unmarshaller for Action */
func (a *Protocol) UnmarshalJSON(data []byte) error {
	var statusStr string
	if err := json.Unmarshal(data, &statusStr); err != nil {
		return err
	}
	if err := a.FromString(statusStr); err != nil {
		return err
	}
	return nil
}

/* Custom Marshaller for Status */
func (a Protocol) MarshalJSON() ([]byte, error) {
	ret := a.String()
	if ret == "" {
		return nil, fmt.Errorf("bad protocol data `%v`", int(a))
	}
	return json.Marshal(ret)
}

type AuthenticationData struct {
	User     string   `json:"user"`
	Password string   `json:"password"`
	IP       string   `json:"ip" binding:"required"`
	Action   Action   `json:"action" binding:"required"`
	Path     string   `json:"path"`
	Protocol Protocol `json:"protocol" binding:"required"`
	ID       string   `json:"id"`
	Query    string   `json:"query"`
}

func (h AuthenticationData) String() string {
	return fmt.Sprintf("User: %s, Password: %s, IP: %s, Action: %s, Path: %s, Protocol: %s, ID: %s, Query: %s",
		h.User, h.Password, h.IP, h.Action.String(), h.Path, h.Protocol.String(), h.ID, h.Query)
}

type AuthenticationCallback func(*AuthenticationData) (allow bool)

func (mtx *mediamtx) AddAuthenticationCallback(callback AuthenticationCallback) {

	mtx.authMutex.Lock()
	defer mtx.authMutex.Unlock()
	mtx.authCb = callback

}
