package repo

import (
	"net/http"
	"sereport/sereport/usecase"

	"github.com/gorilla/mux"
)

// APIRepo endpoint
type APIRepo struct {
	repoUsecase *usecase.RepoUsecase
}

// NewAPIRepo : Constructor
func NewAPIRepo(repoUsecase *usecase.RepoUsecase) *APIRepo {
	apiRepo := new(APIRepo)
	apiRepo.repoUsecase = repoUsecase
	return apiRepo
}

// RegisterRoute : For Http route registering
func (ar *APIRepo) RegisterRoute(r *mux.Router) *mux.Router {
	sr := r.PathPrefix("/repo").Subrouter()
	sr.HandleFunc("/list", ar.list).Methods("GET")
	sr.HandleFunc("/register", ar.register).Methods("POST")
	sr.HandleFunc("/html", ar.html).Methods("POST")
	return sr
}

func (ar *APIRepo) list(w http.ResponseWriter, req *http.Request) {

}

func (ar *APIRepo) register(w http.ResponseWriter, req *http.Request) {

}

func (ar *APIRepo) html(w http.ResponseWriter, req *http.Request) {
	req.ParseMultipartForm(32 << 20) // 32KB

	repoName := req.Form.Get("name")
	branchName := req.Form.Get("branch")
	commit := req.Form.Get("commit")
	tag := req.Form.Get("tag")

	file, _, err := req.FormFile("tar")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer file.Close()

	if err := ar.repoUsecase.UploadTarGzip(&repoName, &branchName, &commit, &tag, &file); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
