package mediamtx

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func (mtx *mediamtx) GetGlobalConfiguration() (*GlobalConf, error) {
	url := mtx.baseAddress + "/v3/config/global/get"

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
	var conf GlobalConf
	err = json.Unmarshal(body, &conf)
	return &conf, err
}
func (mtx *mediamtx) PatchGlobalConfiguration(patches map[string]any) error {
	if patches == nil {
		return errors.New("patchs is empty")
	}
	url := mtx.baseAddress + "/v3/config/global/patch"
	payload, err := json.Marshal(patches)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPatch, url, bytes.NewBuffer(payload))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
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

func (mtx *mediamtx) GetDefaultPathConfiguration() (*PathConf, error) {
	url := mtx.baseAddress + "/v3/config/pathdefaults/get"

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
	var conf PathConf
	err = json.Unmarshal(body, &conf)
	return &conf, err
}

func (mtx *mediamtx) PatchDefaultPathConfiguration(patches map[string]any) error {
	if patches == nil {
		return errors.New("patchs is empty")
	}
	url := mtx.baseAddress + "/v3/config/pathdefaults/patch"
	payload, err := json.Marshal(patches)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPatch, url, bytes.NewBuffer(payload))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
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

func (mtx *mediamtx) GetAlltPathConfiguration() (*AllPathConfiguration, error) {
	url := mtx.baseAddress + "/v3/config/paths/list"

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
	var conf AllPathConfiguration
	err = json.Unmarshal(body, &conf)
	return &conf, err
}

func (mtx *mediamtx) GetPathConfiguration(path string) (*PathConf, error) {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	url := mtx.baseAddress + "/v3/config/path/get" + path

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
	var conf PathConf
	err = json.Unmarshal(body, &conf)
	return &conf, err
}

func (mtx *mediamtx) AddPathConfiguration(path string, conf map[string]any) error {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	url := mtx.baseAddress + "/v3/config/paths/add" + path

	payload, err := json.Marshal(conf)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(payload))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
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

func (mtx *mediamtx) PatchPathConfiguration(path string, patches map[string]any) error {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	url := mtx.baseAddress + "/v3/config/paths/patch" + path

	payload, err := json.Marshal(patches)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPatch, url, bytes.NewBuffer(payload))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
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

func (mtx *mediamtx) ReplacePathConfiguration(path string, conf *PathConf) error {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	url := mtx.baseAddress + "/v3/config/paths/add" + path

	payload, err := json.Marshal(conf)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(payload))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
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

func (mtx *mediamtx) DeletePathConfiguration(path string) error {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	url := mtx.baseAddress + "/v3/config/paths/delete" + path
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
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
