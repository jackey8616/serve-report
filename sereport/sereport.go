package sereport

import (
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"

	"sereport/sereport/api"
	api_repo_pkg "sereport/sereport/api/repo"
	"sereport/sereport/entity"
	"sereport/sereport/usecase"
	"sereport/sereport/usecase/fhdl"
)

// Sereport service entity
type Sereport struct {
	dataPath    *string
	repoUsecase *usecase.RepoUsecase

	api *api.API
}

// NewSereport : Constructor
func NewSereport(db *mgo.Database, dataPath *string) *Sereport {
	sr := new(Sereport)
	sr.dataPath = dataPath

	repoInf := entity.NewMongoInf(db.C("Repo"))
	fileHandler := fhdl.NewFileHandler(sr.dataPath)

	sr.repoUsecase = usecase.NewRepoUsecase(repoInf, fileHandler)
	sr.api = api.NewAPI(api_repo_pkg.NewAPIRepo(sr.repoUsecase))
	return sr
}

// RegisterRoute : For http module to register routes.
func (sr *Sereport) RegisterRoute() *mux.Router {
	r := mux.NewRouter()
	sr.api.RegisterRoute(r)
	return r
}
