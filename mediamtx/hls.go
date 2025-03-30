package mediamtx

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// HLSMuxer represents an HLS muxer instance
type HLSMuxer struct {
	Path        string `json:"path"`
	Created     string `json:"created"`
	LastRequest string `json:"lastRequest"`
	BytesSent   int64  `json:"bytesSent"`
}

// HLSMuxerList represents a paginated list of HLS muxers
type HLSMuxerList struct {
	PageCount int        `json:"pageCount"`
	ItemCount int        `json:"itemCount"`
	Items     []HLSMuxer `json:"items"`
}

// ListHLSMuxers returns all HLS muxers with pagination
func (mtx *mediamtx) ListHLSMuxers(page, itemsPerPage int) (*HLSMuxerList, error) {
	url := fmt.Sprintf("%s/v3/hlsmuxers/list?page=%d&itemsPerPage=%d",
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

	var muxers HLSMuxerList
	err = json.Unmarshal(body, &muxers)
	return &muxers, err
}

// GetHLSMuxer returns a specific HLS muxer by path name
func (mtx *mediamtx) GetHLSMuxer(name string) (*HLSMuxer, error) {
	if !strings.HasPrefix(name, "/") {
		name = "/" + name
	}
	url := fmt.Sprintf("%s/v3/hlsmuxers/get%s", mtx.baseAddress, name)

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

	var muxer HLSMuxer
	err = json.Unmarshal(body, &muxer)
	return &muxer, err
}
