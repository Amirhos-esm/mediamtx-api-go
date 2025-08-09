package mediamtx

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"gopkg.in/yaml.v2"
)

// Path defines model for Path.
type Path struct {
	BytesReceived *int64        `json:"bytesReceived,omitempty"`
	BytesSent     *int64        `json:"bytesSent,omitempty"`
	ConfName      *string       `json:"confName,omitempty"`
	Name          *string       `json:"name,omitempty"`
	Readers       *[]PathReader `json:"readers,omitempty"`
	Ready         *bool         `json:"ready,omitempty"`
	ReadyTime     *string       `json:"readyTime"`
	Source        *PathSource   `json:"source,omitempty"`
	Tracks        *[]string     `json:"tracks,omitempty"`
}

type AllPath struct {
	PageCount int    `json:"pageCount"`
	ItemCount int    `json:"itemCount"`
	Items     []Path `json:"items"`
}

// String returns a YAML representation of the AllPath struct.
func (ap AllPath) String() string {
	yamlData, err := yaml.Marshal(ap)
	if err != nil {
		return fmt.Sprintf("error marshaling AllPath to YAML: %v", err)
	}
	return string(yamlData)
}

// String returns a YAML representation of the Path struct.
func (p Path) String() string {
	yamlData, err := yaml.Marshal(p)
	if err != nil {
		return fmt.Sprintf("error marshaling Path to YAML: %v", err)
	}
	return string(yamlData)
}

func (mtx *Mediamtx) GetAllPath(page, itemsPerPage int) (*AllPath, error) {
	url := fmt.Sprintf("%s/v3/paths/list?page=%d&itemsPerPage=%d",
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
		return nil, fmt.Errorf("server response with %d code and response: %v", resp.StatusCode, string(body))
	}
	var conf AllPath
	err = json.Unmarshal(body, &conf)
	return &conf, err
}

func (mtx *Mediamtx) GetPath(path string) (*Path, error) {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	url := mtx.baseAddress + "/v3/paths/get" + path

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
		return nil, fmt.Errorf("server response with %d code and response: %v", resp.StatusCode, string(body))
	}
	var conf Path
	err = json.Unmarshal(body, &conf)
	return &conf, err
}
