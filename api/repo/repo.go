package repo

import (
	"net/http"

	"github.com/gorilla/mux"
)

type RepoAPI struct{}

func NewRepoAPI() *RepoAPI {
	return &RepoAPI{}
}

func (rAPI *RepoAPI) RegisterRoute(r *mux.Router) *mux.Router {
	sr := r.PathPrefix("/repo").Subrouter()
	sr.HandleFunc("/list", rAPI.list).Methods("GET")
	sr.HandleFunc("/register", rAPI.register).Methods("POST")
	return sr
}

func (rAPI *RepoAPI) list(w http.ResponseWriter, req *http.Request) {

}

func (rAPI *RepoAPI) register(w http.ResponseWriter, req *http.Request) {

}
