package monobanksdk

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func doPostRequest(url, token string, body []byte) ([]byte, int, error) {
	return doRequest(http.MethodPost, url, token, body)
}
func doGetRequest(url, token string) ([]byte, int, error) {
	return doRequest(http.MethodGet, url, token, nil)
}

func doRequest(method, url, token string, requestBody []byte) ([]byte, int, error) {
	req, _ := http.NewRequest(method, fmt.Sprintf(MonoURL, url), bytes.NewBuffer(requestBody))

	req.Header.Add("content-type", "application/json")

	if token != "" {
		req.Header.Add(TokenHeaderName, token)
	}

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	return body, res.StatusCode, nil
}
