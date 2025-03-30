package mediamtx

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"gopkg.in/yaml.v2"
)

type Recording struct {
	Name     string             `json:"name"`
	Segments []RecordingSegment `json:"segments"`
}

type RecordingSegment struct {
	Start string `json:"start"`
}

type RecordingList struct {
	PageCount int         `json:"pageCount"`
	ItemCount int         `json:"itemCount"`
	Items     []Recording `json:"items"`
}

// String returns a YAML representation of the AllPath struct.
func (ap Recording) String() string {
	yamlData, err := yaml.Marshal(ap)
	if err != nil {
		return fmt.Sprintf("error marshaling AllPath to YAML: %v", err)
	}
	return string(yamlData)
}

// String returns a YAML representation of the AllPath struct.
func (ap RecordingList) String() string {
	yamlData, err := yaml.Marshal(ap)
	if err != nil {
		return fmt.Sprintf("error marshaling AllPath to YAML: %v", err)
	}
	return string(yamlData)
}

func (mtx *mediamtx) ListRecordings(page, itemsPerPage int) (*RecordingList, error) {
	url := fmt.Sprintf("%s/v3/recordings/list?page=%d&itemsPerPage=%d", mtx.baseAddress, page, itemsPerPage)

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

	var recordings RecordingList
	err = json.Unmarshal(body, &recordings)
	return &recordings, err
}

func (mtx *mediamtx) GetRecordings(pathName string) (*Recording, error) {
	if !strings.HasPrefix(pathName, "/") {
		pathName = "/" + pathName
	}
	url := mtx.baseAddress + "/v3/recordings/get" + pathName

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

	var recording Recording
	err = json.Unmarshal(body, &recording)
	return &recording, err
}

func (mtx *mediamtx) DeleteRecordingSegment(pathName, startTime string) error {
	url := fmt.Sprintf("%s/v3/recordings/deletesegment?path=%s&start=%s",
		mtx.baseAddress,
		url.QueryEscape(pathName),
		url.QueryEscape(startTime))

	req, err := http.NewRequest("DELETE", url, nil)
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
		return fmt.Errorf("server response with %d code and response: %v", resp.StatusCode, string(body))
	}

	return nil
}
