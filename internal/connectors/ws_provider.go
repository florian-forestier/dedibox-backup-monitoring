package connectors

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
)

func GetURL[T any](url string, headers map[string]string, expectedCode int) (bodyResponse T, err error) {
	client := &http.Client{}

	request, _ := http.NewRequest("GET", url, nil)
	for key, value := range headers {
		request.Header.Add(key, value)
	}

	response, err := client.Do(request)

	if err != nil {
		return
	}

	if response.StatusCode != expectedCode {
		err = errors.New("expected status code " + strconv.Itoa(expectedCode) + " but got " + strconv.Itoa(response.StatusCode))
		return
	}

	body, _ := io.ReadAll(response.Body)
	_ = json.Unmarshal(body, &bodyResponse)
	return
}
