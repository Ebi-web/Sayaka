package api

import (
	"net/http"

	"Sayaka/utils"
)

type AppHandler struct {
	h func(http.ResponseWriter, *http.Request) (int, interface{}, error)
}

func (a AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	status, res, err := a.h(w, r)
	if err != nil {
		utils.RespondErrorJson(w, status, err)
		return
	}
	utils.RespondJSON(w, status, res)
	return
}
