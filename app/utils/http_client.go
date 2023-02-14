package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/pkg/errors"
)

func MakeRequest(method, url string, headers map[string]string, body interface{}) ([]byte, error) {
	req, err := http.NewRequest(method, url, parseBody(body))
	if err != nil {
		return nil, err
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if res.StatusCode >= http.StatusBadRequest {
		return nil, errors.New(res.Status)
	}
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

func parseBody(b interface{}) io.Reader {
	switch v := b.(type) {
	case []byte:
		return bytes.NewBuffer(v)
	case string:
		return bytes.NewBufferString(v)
	default:
		return nil
	}
}

func ParseRequestBody(r *http.Request, v interface{}) error {
	// リクエストボディを読み込む
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return errors.Wrap(err, "failed to read request body")
	}

	// リクエストボディを指定した型にマッピングする
	if err := json.Unmarshal(body, v); err != nil {
		return errors.Wrap(err, "failed to unmarshal request body")
	}

	return nil
}
