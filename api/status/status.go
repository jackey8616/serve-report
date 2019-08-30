package status

import (
	"encoding/json"
	"net/http"
	"serve_report/entity"
)

func StatusAPI(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "Application/Json")
	json, _ := json.Marshal(entity.StatusResponse{
		Hello: "World",
	})
	w.Write(json)
}
