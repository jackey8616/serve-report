package usecase

import (
	"fmt"
	"mime/multipart"
	"sereport/sereport/entity"
	"sereport/sereport/usecase/fhdl"

	"gopkg.in/mgo.v2/bson"
)

// RepoUsecase application class
type RepoUsecase struct {
	inf         *entity.MongoInf
	fileHandler *fhdl.FileHandler
}

// NewRepoUsecase : Constructor
func NewRepoUsecase(inf *entity.MongoInf, fileHandler *fhdl.FileHandler) *RepoUsecase {
	ru := new(RepoUsecase)
	ru.inf = inf
	ru.fileHandler = fileHandler
	return ru
}

// UploadTarGzip : Process uploaded file from API route.
func (ru *RepoUsecase) UploadTarGzip(repoName, branchName, commit, tag *string, uploadedFile *multipart.File) error {
	fileName := fmt.Sprintf("%s.tar.gz", *commit)
	tarFilePath, err := ru.fileHandler.SaveUploadedFile(repoName, branchName, tag, &fileName, uploadedFile)
	if err != nil {
		return err
	}
	if err := ru.fileHandler.UnTarGzipHTML(tarFilePath, repoName, branchName, tag, commit); err != nil {
		return err
	}
	return nil
}

// GetRepoByName : get repo by name.
func (ru *RepoUsecase) GetRepoByName(name string) (*entity.Repo, error) {
	result, err := ru.inf.FindOne(bson.M{"name": &name}, &entity.Repo{})
	if err != nil {
		return nil, err
	}
	return result.(*entity.Repo), nil
}
