package mediamtx

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"gopkg.in/yaml.v2"
)

// RTMPConn represents an RTMP connection
type RTMPConn struct {
	ID            string `json:"id"`
	Created       string `json:"created"`
	RemoteAddr    string `json:"remoteAddr"`
	State         string `json:"state"` // "idle", "read", or "publish"
	Path          string `json:"path"`
	Query         string `json:"query"`
	BytesReceived int64  `json:"bytesReceived"`
	BytesSent     int64  `json:"bytesSent"`
}

// RTMPConnList represents a paginated list of RTMP connections
type RTMPConnList struct {
	PageCount int        `json:"pageCount"`
	ItemCount int        `json:"itemCount"`
	Items     []RTMPConn `json:"items"`
}

// String returns a YAML representation of the AllPath struct.
func (ap RTMPConn) String() string {
	yamlData, err := yaml.Marshal(ap)
	if err != nil {
		return fmt.Sprintf("error marshaling AllPath to YAML: %v", err)
	}
	return string(yamlData)
}

// String returns a YAML representation of the AllPath struct.
func (ap RTMPConnList) String() string {
	yamlData, err := yaml.Marshal(ap)
	if err != nil {
		return fmt.Sprintf("error marshaling AllPath to YAML: %v", err)
	}
	return string(yamlData)
}

// ListRTMPConns returns all RTMP connections with pagination
func (mtx *Mediamtx) ListRTMPConns(page, itemsPerPage int) (*RTMPConnList, error) {
	url := fmt.Sprintf("%s/v3/rtmpconns/list?page=%d&itemsPerPage=%d",
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

	var conns RTMPConnList
	err = json.Unmarshal(body, &conns)
	return &conns, err
}

// GetRTMPConn returns a specific RTMP connection by ID
func (mtx *Mediamtx) GetRTMPConn(id string) (*RTMPConn, error) {
	url := fmt.Sprintf("%s/v3/rtmpconns/get/%s", mtx.baseAddress, id)

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

	var conn RTMPConn
	err = json.Unmarshal(body, &conn)
	return &conn, err
}

// KickRTMPConn kicks out an RTMP connection from the server
func (mtx *Mediamtx) KickRTMPConn(id string) error {
	url := fmt.Sprintf("%s/v3/rtmpconns/kick/%s", mtx.baseAddress, id)

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

// ListRTMPSConns returns all secure RTMPS connections with pagination
func (mtx *Mediamtx) ListRTMPSConns(page, itemsPerPage int) (*RTMPConnList, error) {
	url := fmt.Sprintf("%s/v3/rtmpsconns/list?page=%d&itemsPerPage=%d",
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

	var conns RTMPConnList
	err = json.Unmarshal(body, &conns)
	return &conns, err
}

// GetRTMPSConn returns a specific RTMPS connection by ID
func (mtx *Mediamtx) GetRTMPSConn(id string) (*RTMPConn, error) {
	url := fmt.Sprintf("%s/v3/rtmpsconns/get/%s", mtx.baseAddress, id)

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

	var conn RTMPConn
	err = json.Unmarshal(body, &conn)
	return &conn, err
}

// KickRTMPSConn kicks out an RTMPS connection from the server
func (mtx *Mediamtx) KickRTMPSConn(id string) error {
	url := fmt.Sprintf("%s/v3/rtmpsconns/kick/%s", mtx.baseAddress, id)

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
