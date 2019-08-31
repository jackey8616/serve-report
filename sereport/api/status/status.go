package status

import (
	"encoding/json"
	"net/http"

	"sereport/sereport/entity"
)

// APIStatus : For Http route registering
func APIStatus(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "Application/Json")
	json, _ := json.Marshal(entity.StatusResponse{
		Hello: "World",
	})
	w.Write(json)
}
