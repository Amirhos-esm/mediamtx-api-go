package mediamtx

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"gopkg.in/yaml.v2"
)

// SRTConn represents an SRT connection with all statistics
type SRTConn struct {
	ID         string `json:"id"`
	Created    string `json:"created"`
	RemoteAddr string `json:"remoteAddr"`
	State      string `json:"state"` // "idle", "read", or "publish"
	Path       string `json:"path"`
	Query      string `json:"query"`

	// Packet statistics
	PacketsSent              int64 `json:"packetsSent"`
	PacketsReceived          int64 `json:"packetsReceived"`
	PacketsReceivedBelated   int64 `json:"packetsReceivedBelated"`
	PacketsSentUnique        int64 `json:"packetsSentUnique"`
	PacketsReceivedUnique    int64 `json:"packetsReceivedUnique"`
	PacketsSendLoss          int64 `json:"packetsSendLoss"`
	PacketsReceivedLoss      int64 `json:"packetsReceivedLoss"`
	PacketsRetrans           int64 `json:"packetsRetrans"`
	PacketsReceivedRetrans   int64 `json:"packetsReceivedRetrans"`
	PacketsSentACK           int64 `json:"packetsSentACK"`
	PacketsReceivedACK       int64 `json:"packetsReceivedACK"`
	PacketsSentNAK           int64 `json:"packetsSentNAK"`
	PacketsReceivedNAK       int64 `json:"packetsReceivedNAK"`
	PacketsSentKM            int64 `json:"packetsSentKM"`
	PacketsReceivedKM        int64 `json:"packetsReceivedKM"`
	PacketsSendDrop          int64 `json:"packetsSendDrop"`
	PacketsReceivedDrop      int64 `json:"packetsReceivedDrop"`
	PacketsReceivedUndecrypt int64 `json:"packetsReceivedUndecrypt"`

	// Byte statistics
	BytesSent              int64 `json:"bytesSent"`
	BytesReceived          int64 `json:"bytesReceived"`
	BytesReceivedBelated   int64 `json:"bytesReceivedBelated"`
	BytesSentUnique        int64 `json:"bytesSentUnique"`
	BytesReceivedUnique    int64 `json:"bytesReceivedUnique"`
	BytesReceivedLoss      int64 `json:"bytesReceivedLoss"`
	BytesRetrans           int64 `json:"bytesRetrans"`
	BytesReceivedRetrans   int64 `json:"bytesReceivedRetrans"`
	BytesSendDrop          int64 `json:"bytesSendDrop"`
	BytesReceivedDrop      int64 `json:"bytesReceivedDrop"`
	BytesReceivedUndecrypt int64 `json:"bytesReceivedUndecrypt"`

	// Performance metrics
	UsPacketsSendPeriod           float64 `json:"usPacketsSendPeriod"`
	PacketsFlowWindow             int64   `json:"packetsFlowWindow"`
	PacketsFlightSize             int64   `json:"packetsFlightSize"`
	MsRTT                         float64 `json:"msRTT"`
	MbpsSendRate                  float64 `json:"mbpsSendRate"`
	MbpsReceiveRate               float64 `json:"mbpsReceiveRate"`
	MbpsLinkCapacity              float64 `json:"mbpsLinkCapacity"`
	BytesAvailSendBuf             int64   `json:"bytesAvailSendBuf"`
	BytesAvailReceiveBuf          int64   `json:"bytesAvailReceiveBuf"`
	MbpsMaxBW                     float64 `json:"mbpsMaxBW"`
	ByteMSS                       int64   `json:"byteMSS"`
	PacketsSendBuf                int64   `json:"packetsSendBuf"`
	BytesSendBuf                  int64   `json:"bytesSendBuf"`
	MsSendBuf                     int64   `json:"msSendBuf"`
	MsSendTsbPdDelay              int64   `json:"msSendTsbPdDelay"`
	PacketsReceiveBuf             int64   `json:"packetsReceiveBuf"`
	BytesReceiveBuf               int64   `json:"bytesReceiveBuf"`
	MsReceiveBuf                  int64   `json:"msReceiveBuf"`
	MsReceiveTsbPdDelay           int64   `json:"msReceiveTsbPdDelay"`
	PacketsReorderTolerance       int64   `json:"packetsReorderTolerance"`
	PacketsReceivedAvgBelatedTime int64   `json:"packetsReceivedAvgBelatedTime"`
	PacketsSendLossRate           float64 `json:"packetsSendLossRate"`
	PacketsReceivedLossRate       float64 `json:"packetsReceivedLossRate"`
}

// SRTConnList represents a paginated list of SRT connections
type SRTConnList struct {
	PageCount int       `json:"pageCount"`
	ItemCount int       `json:"itemCount"`
	Items     []SRTConn `json:"items"`
}

// String returns a YAML representation of the AllPath struct.
func (ap SRTConnList) String() string {
	yamlData, err := yaml.Marshal(ap)
	if err != nil {
		return fmt.Sprintf("error marshaling AllPath to YAML: %v", err)
	}
	return string(yamlData)
}

// String returns a YAML representation of the Path struct.
func (p SRTConn) String() string {
	yamlData, err := yaml.Marshal(p)
	if err != nil {
		return fmt.Sprintf("error marshaling Path to YAML: %v", err)
	}
	return string(yamlData)
}

// ListSRTConns returns all SRT connections with pagination
func (mtx *Mediamtx) ListSRTConns(page, itemsPerPage int) (*SRTConnList, error) {
	url := fmt.Sprintf("%s/v3/srtconns/list?page=%d&itemsPerPage=%d",
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

	var conns SRTConnList
	err = json.Unmarshal(body, &conns)
	return &conns, err
}

// GetSRTConn returns a specific SRT connection by ID
func (mtx *Mediamtx) GetSRTConn(id string) (*SRTConn, error) {
	url := fmt.Sprintf("%s/v3/srtconns/get/%s", mtx.baseAddress, id)

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

	var conn SRTConn
	err = json.Unmarshal(body, &conn)
	return &conn, err
}

// KickSRTConn kicks out an SRT connection from the server
func (mtx *Mediamtx) KickSRTConn(id string) error {
	url := fmt.Sprintf("%s/v3/srtconns/kick/%s", mtx.baseAddress, id)

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
