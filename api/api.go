package api

import (
	"github.com/gorilla/mux"

	"serve_report/api/repo"
	"serve_report/api/status"
)

type API struct {
	repoAPI *repo.RepoAPI
}

func NewAPI() *API {
	api := &API{}
	api.repoAPI = repo.NewRepoAPI()
	return api
}

func (api *API) RegisterRoute(r *mux.Router) *mux.Router {
	sr := r.PathPrefix("/api").Subrouter()
	sr.HandleFunc("/status", status.StatusAPI)
	api.repoAPI.RegisterRoute(sr)
	return sr
}
