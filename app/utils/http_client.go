package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func MakeRequest(method, url string, headers map[string]string, body []byte) ([]byte, error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	responseBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return responseBody, nil
}

func ParseRequestBody(r *http.Request, v interface{}) error {
	// リクエストボディを読み込む
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}

	// リクエストボディを指定した型にマッピングする
	if err := json.Unmarshal(body, v); err != nil {
		return err
	}

	return nil
}
