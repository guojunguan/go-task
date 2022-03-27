package client

import (
	"encoding/json"
	"net/http"
	"time"
)

func Get(url string, body string, headers map[string]string, target interface{}) error {
	c := http.Client{Timeout: time.Duration(3) * time.Second}
	req, requestError := http.NewRequest("GET", url, nil)

	if requestError != nil {
		return requestError
	}

	for key, value := range headers {
		req.Header.Add(key, value)
	}

	response, responseError := c.Do(req)
	if responseError != nil {
		return responseError
	}

	defer response.Body.Close()
	return json.NewDecoder(response.Body).Decode(&target)
}
