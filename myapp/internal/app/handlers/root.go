package handlers

import (
	"fmt"
	"net/http"

	"github.com/Sun1tAr/MIREA-TIP-Practice-2/myapp/utils"
)

func Root(w http.ResponseWriter, r *http.Request) {
	utils.LogRequest(r)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintln(w, "Hello, Go project structure!")
}
