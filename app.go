package serve_report

import (
	"net/http"

	"github.com/gorilla/mux"

	"serve_report/api"
)

func main() {

	api := &api.API{}

	r := mux.NewRouter()
	api.RegisterRoute(r)
	http.Handle("/", r)
	http.ListenAndServe(":8000", nil)
}
