package utils

import (
	"encoding/json"
	"log"
	"net/http"
	"reflect"
)

func RespondErrorJson(w http.ResponseWriter, code int, err error) {
	log.Printf("code=%d, err=%+v", code, err)
	if e, ok := err.(*HTTPError); ok {
		RespondJSON(w, code, e)
	} else if err != nil {
		he := HTTPError{
			Message: err.Error(),
		}
		RespondJSON(w, code, he)
	}
}

func RespondJSON(w http.ResponseWriter, status int, payload interface{}) {
	res, err := marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, writeErr := w.Write([]byte(err.Error()))
		if writeErr != nil {
			log.Print(writeErr)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, writeErr := w.Write(res)
	if writeErr != nil {
		log.Print(writeErr)
	}
}

func marshal(payload interface{}) ([]byte, error) {
	if isNil(payload) {
		return []byte(`{}`), nil
	}
	return json.Marshal(payload)
}

func isNil(p interface{}) bool {
	if p == nil {
		return true
	}
	switch reflect.TypeOf(p).Kind() {
	case reflect.Ptr, reflect.Map, reflect.Slice, reflect.Array:
		return reflect.ValueOf(p).IsNil()
	}
	return false
}
