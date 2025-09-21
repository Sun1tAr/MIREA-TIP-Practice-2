package handlers

import (
	"net/http"

	"github.com/Sun1tAr/MIREA-TIP-Practice-2/myapp/utils"
)

func Fail(w http.ResponseWriter, r *http.Request) {
	utils.LogRequest(r)
	utils.WriteErr(w, http.StatusBadRequest, "bad_request_example")
}
