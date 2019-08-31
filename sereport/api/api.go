package api

import (
	"github.com/gorilla/mux"

	"sereport/sereport/api/repo"
	"sereport/sereport/api/status"
)

// API endpoint
type API struct {
	repoAPI *repo.APIRepo
}

// NewAPI : Constructor
func NewAPI(repoAPI *repo.APIRepo) *API {
	api := new(API)
	api.repoAPI = repoAPI
	return api
}

// RegisterRoute : For Http route registering
func (api *API) RegisterRoute(r *mux.Router) *mux.Router {
	sr := r.PathPrefix("/api").Subrouter()
	sr.HandleFunc("/status", status.APIStatus)
	api.repoAPI.RegisterRoute(sr)
	return sr
}
