package utils

import (
	"io/ioutil"
	"net/http"
    "time"
)

type HttpUtils struct {

}

var netClient = &http.Client{
	Timeout: time.Second * 10,
}

func NewHttpUtils() *HttpUtils {
    var n HttpUtils
    return &n
}

func (instance *HttpUtils) Get(url string) (string, error) {
    req, err := http.NewRequest("GET", url, nil)
	resp, err := netClient.Do(req)

	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	return string(body), nil
}
