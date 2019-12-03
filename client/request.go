package client 

import (
	"io"
	"net/http"
	"net/url"
	"strings"
)

const endPoint = "https://www.city.aizuwakamatsu.fukushima.jp/index_php/gomical/index_i.php?typ=p"

func createGetRequest() (*http.Request, error) {
	req, err := createRequest("GET", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.122 Safari/537.36 Vivaldi/2.3.1440.61")

	return req, nil
}

func createPostRequest(id string) (*http.Request, error) {
	values := &url.Values{}
	values.Add("m", id)
	values.Add("d", "1")

	req, err := createRequest("POST", values)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.122 Safari/537.36 Vivaldi/2.3.1440.61")

	return req, nil
}

func createRequest(method string, values *url.Values) (*http.Request, error) {
	var body io.Reader
	if values != nil {
		body = strings.NewReader(values.Encode())
	}

	return http.NewRequest(method, endPoint, body)
}
