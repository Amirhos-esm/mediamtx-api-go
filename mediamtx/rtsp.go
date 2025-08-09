package mediamtx

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"gopkg.in/yaml.v2"
)

// RTSPConn represents an RTSP connection
type RTSPConn struct {
	ID            string `json:"id"`
	Created       string `json:"created"`
	RemoteAddr    string `json:"remoteAddr"`
	BytesReceived int64  `json:"bytesReceived"`
	BytesSent     int64  `json:"bytesSent"`
	Session       string `json:"session,omitempty"` // nullable
}

// RTSPConnList represents a paginated list of RTSP connections
type RTSPConnList struct {
	PageCount int        `json:"pageCount"`
	ItemCount int        `json:"itemCount"`
	Items     []RTSPConn `json:"items"`
}

// RTSPSession represents an RTSP session
type RTSPSession struct {
	ID                  string  `json:"id"`
	Created             string  `json:"created"`
	RemoteAddr          string  `json:"remoteAddr"`
	State               string  `json:"state"` // "idle", "read", or "publish"
	Path                string  `json:"path"`
	Query               string  `json:"query"`
	Transport           string  `json:"transport,omitempty"` // nullable
	BytesReceived       int64   `json:"bytesReceived"`
	BytesSent           int64   `json:"bytesSent"`
	RtpPacketsReceived  int64   `json:"rtpPacketsReceived"`
	RtpPacketsSent      int64   `json:"rtpPacketsSent"`
	RtpPacketsLost      int64   `json:"rtpPacketsLost"`
	RtpPacketsInError   int64   `json:"rtpPacketsInError"`
	RtpPacketsJitter    float64 `json:"rtpPacketsJitter"`
	RtcpPacketsReceived int64   `json:"rtcpPacketsReceived"`
	RtcpPacketsSent     int64   `json:"rtcpPacketsSent"`
	RtcpPacketsInError  int64   `json:"rtcpPacketsInError"`
}

// RTSPSessionList represents a paginated list of RTSP sessions
type RTSPSessionList struct {
	PageCount int           `json:"pageCount"`
	ItemCount int           `json:"itemCount"`
	Items     []RTSPSession `json:"items"`
}

// String returns a YAML representation of the AllPath struct.
func (ap RTSPConnList) String() string {
	yamlData, err := yaml.Marshal(ap)
	if err != nil {
		return fmt.Sprintf("error marshaling AllPath to YAML: %v", err)
	}
	return string(yamlData)
}

// String returns a YAML representation of the Path struct.
func (p RTSPConn) String() string {
	yamlData, err := yaml.Marshal(p)
	if err != nil {
		return fmt.Sprintf("error marshaling Path to YAML: %v", err)
	}
	return string(yamlData)
}

// String returns a YAML representation of the AllPath struct.
func (ap RTSPSession) String() string {
	yamlData, err := yaml.Marshal(ap)
	if err != nil {
		return fmt.Sprintf("error marshaling AllPath to YAML: %v", err)
	}
	return string(yamlData)
}

// String returns a YAML representation of the Path struct.
func (p RTSPSessionList) String() string {
	yamlData, err := yaml.Marshal(p)
	if err != nil {
		return fmt.Sprintf("error marshaling Path to YAML: %v", err)
	}
	return string(yamlData)
}

// ListRTSPConns returns all RTSP connections with pagination
func (mtx *Mediamtx) ListRTSPConns(page, itemsPerPage int) (*RTSPConnList, error) {
	url := fmt.Sprintf("%s/v3/rtspconns/list?page=%d&itemsPerPage=%d",
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

	var conns RTSPConnList
	err = json.Unmarshal(body, &conns)
	return &conns, err
}

// GetRTSPConn returns a specific RTSP connection by ID
func (mtx *Mediamtx) GetRTSPConn(id string) (*RTSPConn, error) {
	url := fmt.Sprintf("%s/v3/rtspconns/get/%s", mtx.baseAddress, id)

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

	var conn RTSPConn
	err = json.Unmarshal(body, &conn)
	return &conn, err
}

// ListRTSPSessions returns all RTSP sessions with pagination
func (mtx *Mediamtx) ListRTSPSessions(page, itemsPerPage int) (*RTSPSessionList, error) {
	url := fmt.Sprintf("%s/v3/rtspsessions/list?page=%d&itemsPerPage=%d",
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

	var sessions RTSPSessionList
	err = json.Unmarshal(body, &sessions)
	return &sessions, err
}

// GetRTSPSession returns a specific RTSP session by ID
func (mtx *Mediamtx) GetRTSPSession(id string) (*RTSPSession, error) {
	url := fmt.Sprintf("%s/v3/rtspsessions/get/%s", mtx.baseAddress, id)

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

	var session RTSPSession
	err = json.Unmarshal(body, &session)
	return &session, err
}

// KickRTSPSession kicks out an RTSP session from the server
func (mtx *Mediamtx) KickRTSPSession(id string) error {
	url := fmt.Sprintf("%s/v3/rtspsessions/kick/%s", mtx.baseAddress, id)

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

// ListRTSPConnsSecure returns all secure RTSPS connections with pagination
func (mtx *Mediamtx) ListRTSPConnsSecure(page, itemsPerPage int) (*RTSPConnList, error) {
	url := fmt.Sprintf("%s/v3/rtspsconns/list?page=%d&itemsPerPage=%d",
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

	var conns RTSPConnList
	err = json.Unmarshal(body, &conns)
	return &conns, err
}

// ListRTSPSessionsSecure returns all secure RTSPS sessions with pagination
func (mtx *Mediamtx) ListRTSPSessionsSecure(page, itemsPerPage int) (*RTSPSessionList, error) {
	url := fmt.Sprintf("%s/v3/rtspssessions/list?page=%d&itemsPerPage=%d",
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

	var sessions RTSPSessionList
	err = json.Unmarshal(body, &sessions)
	return &sessions, err
}

// GetRTSPConn returns a specific RTSPS connection by ID
func (mtx *Mediamtx) GetRTSPConnSecure(id string) (*RTSPConn, error) {
	url := fmt.Sprintf("%s/v3/rtspsconns/get/%s", mtx.baseAddress, id)

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

	var conn RTSPConn
	err = json.Unmarshal(body, &conn)
	return &conn, err
}

// GetRTSPSession returns a specific RTSPS session by ID
func (mtx *Mediamtx) GetRTSPSessionSecure(id string) (*RTSPSession, error) {
	url := fmt.Sprintf("%s/v3/rtspssessions/get/%s", mtx.baseAddress, id)

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

	var session RTSPSession
	err = json.Unmarshal(body, &session)
	return &session, err
}

// KickRTSPSession kicks out an RTSPS session from the server
func (mtx *Mediamtx) KickRTSPSessionSecure(id string) error {
	url := fmt.Sprintf("%s/v3/rtspssessions/kick/%s", mtx.baseAddress, id)

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
