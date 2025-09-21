package app

import (
	"net/http"

	"github.com/Sun1tAr/MIREA-TIP-Practice-2/myapp/internal/app/handlers"
	"github.com/Sun1tAr/MIREA-TIP-Practice-2/myapp/utils"
)

type pingResp struct {
	Status string `json:"status"`
	Time   string `json:"time"`
}

func Run() {
	mux := http.NewServeMux()

	// Корневой маршрут
	mux.HandleFunc("/", handlers.Root)
	// /ping
	mux.HandleFunc("/ping", handlers.Ping)
	// badRequest
	mux.HandleFunc("/fail", handlers.Fail)

	handler := withRequestID(mux)

	if handler != nil {
		utils.LogInfo("Server is starting on :8080")
		if err := http.ListenAndServe(":8080", handler); err != nil {
			utils.LogError("server error: " + err.Error())
		}
	} else {
		utils.LogInfo("Server is starting on :8080")
		if err := http.ListenAndServe(":8080", mux); err != nil {
			utils.LogError("server error: " + err.Error())
		}
	}
}

func withRequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.Header.Get("X-Request-Id")
		if id == "" {
			id = utils.NewID16()
		}
		w.Header().Set("X-Request-Id", id)
		next.ServeHTTP(w, r)
	})
}
