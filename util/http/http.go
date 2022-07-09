package http

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"github.com/ratel-online/core/util/json"
	"github.com/ratel-online/core/util/strings"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

var httpClient = &http.Client{}
var emtpy = struct{}{}

func init() {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	httpClient = &http.Client{
		Transport: tr,
		Timeout:   time.Duration(3) * time.Second,
	}
}

func Get(url string, params map[string]interface{}) ([]byte, error) {
	resp, err := httpClient.Get(url + encodeQueries(params))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func Post(url string, params map[string]interface{}, body interface{}) ([]byte, error) {
	if body == nil {
		body = emtpy
	}
	resp, err := httpClient.Post(url+encodeQueries(params), "application/json", bytes.NewReader(json.Marshal(body)))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func encodeQueries(params map[string]interface{}) string {
	if len(params) == 0 {
		return ""
	}
	var buffer bytes.Buffer
	buffer.WriteString("?")
	for k, v := range params {
		if v == nil {
			buffer.WriteString(fmt.Sprintf("%s=", url.QueryEscape(k)))
			continue
		}
		buffer.WriteString(fmt.Sprintf("%s=%v&", url.QueryEscape(k), url.QueryEscape(strings.String(v))))
	}
	buffer.Truncate(buffer.Len() - 1)
	return buffer.String()
}
