package mediamtx

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"gopkg.in/yaml.v2"
)

// WebRTCSession represents a WebRTC session
type WebRTCSession struct {
	ID                        string `json:"id"`
	Created                   string `json:"created"`
	RemoteAddr                string `json:"remoteAddr"`
	PeerConnectionEstablished bool   `json:"peerConnectionEstablished"`
	LocalCandidate            string `json:"localCandidate"`
	RemoteCandidate           string `json:"remoteCandidate"`
	State                     string `json:"state"` // "read" or "publish"
	Path                      string `json:"path"`
	Query                     string `json:"query"`
	BytesReceived             int64  `json:"bytesReceived"`
	BytesSent                 int64  `json:"bytesSent"`
}

// WebRTCSessionList represents a paginated list of WebRTC sessions
type WebRTCSessionList struct {
	PageCount int             `json:"pageCount"`
	ItemCount int             `json:"itemCount"`
	Items     []WebRTCSession `json:"items"`
}

// String returns a YAML representation of the AllPath struct.
func (ap WebRTCSession) String() string {
	yamlData, err := yaml.Marshal(ap)
	if err != nil {
		return fmt.Sprintf("error marshaling AllPath to YAML: %v", err)
	}
	return string(yamlData)
}

// String returns a YAML representation of the Path struct.
func (p WebRTCSessionList) String() string {
	yamlData, err := yaml.Marshal(p)
	if err != nil {
		return fmt.Sprintf("error marshaling Path to YAML: %v", err)
	}
	return string(yamlData)
}

// ListWebRTCSessions returns all WebRTC sessions with pagination
func (mtx *Mediamtx) ListWebRTCSessions(page, itemsPerPage int) (*WebRTCSessionList, error) {
	url := fmt.Sprintf("%s/v3/webrtcsessions/list?page=%d&itemsPerPage=%d",
		mtx.baseAddress, page, itemsPerPage)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("server responded with %d code and response: %v",
			resp.StatusCode, string(body))
	}

	var sessions WebRTCSessionList
	err = json.Unmarshal(body, &sessions)
	return &sessions, err
}

// GetWebRTCSession returns a specific WebRTC session by ID
func (mtx *Mediamtx) GetWebRTCSession(id string) (*WebRTCSession, error) {
	url := fmt.Sprintf("%s/v3/webrtcsessions/get/%s", mtx.baseAddress, id)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("server responded with %d code and response: %v",
			resp.StatusCode, string(body))
	}

	var session WebRTCSession
	err = json.Unmarshal(body, &session)
	return &session, err
}

// KickWebRTCSession kicks out a WebRTC session from the server
func (mtx *Mediamtx) KickWebRTCSession(id string) error {
	url := fmt.Sprintf("%s/v3/webrtcsessions/kick/%s", mtx.baseAddress, id)

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("server responded with %d code and response: %v",
			resp.StatusCode, string(body))
	}

	return nil
}
