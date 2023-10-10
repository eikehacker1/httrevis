// httpclient.go
package httpclient

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetURL(url string) (*http.Response, error) {
	resp, err := http.Get(url)
	return resp, err
}

func GetStatusAndLength(url string) (int, int, error) {
	resp, err := GetURL(url)
	if err != nil {
		return 0, 0, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, 0, err
	}

	return resp.StatusCode, len(body), nil
}
