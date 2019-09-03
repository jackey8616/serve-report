package fhdl

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"
	"strings"
)

// FileHandler expose
type FileHandler struct {
	dataPath     *string
	repoPath     *string
	repoHTMLPath *string
	repoTarPath  *string
}

// NewFileHandler : Construcutor
func NewFileHandler(dataPath *string) *FileHandler {
	fh := new(FileHandler)
	fh.dataPath = dataPath
	tmp := fmt.Sprintf("%s/repo", *dataPath)
	fh.repoPath = &tmp

	fh.createFolder(*fh.repoPath)
	return fh
}

func (fh *FileHandler) createFolder(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return os.MkdirAll(path, 0755)
	}
	return nil
}

func (fh *FileHandler) createFile(name string) (*os.File, error) {
	if err := fh.createFolder(string([]rune(name)[0:strings.LastIndex(name, "/")])); err != nil {
		return nil, err
	}
	return os.Create(name)
}

// SaveUploadedFile : Write uploaded file into specifiec location.
func (fh *FileHandler) SaveUploadedFile(repoName, branchName, tagName, fileName *string, file *multipart.File) (*string, error) {
	tarFilePath := ""
	if *tagName == "" {
		tarFilePath = fmt.Sprintf("%s/%s/%s/tar/%s", *fh.repoPath, *repoName, *branchName, *fileName)
	} else {
		tarFilePath = fmt.Sprintf("%s/%s/%s/tar/%s/%s", *fh.repoPath, *repoName, *branchName, *tagName, *fileName)
	}
	f, err := fh.createFile(tarFilePath)
	if err != nil {
		return nil, err
	}
	if _, err = io.Copy(f, *file); err != nil {
		return nil, err
	}
	return &tarFilePath, nil
}

// UnTarGzipHTML : Unzip a tar.gz format compresed file which contain HTML.
func (fh *FileHandler) UnTarGzipHTML(tarFilePath, repoName, branchName, tagName, commit *string) error {
	htmlPath := ""
	if *tagName == "" {
		htmlPath = fmt.Sprintf("%s/%s/%s/html/%s/", *fh.repoPath, *repoName, *branchName, *commit)
	} else {
		htmlPath = fmt.Sprintf("%s/%s/%s/html/%s/%s/", *fh.repoPath, *repoName, *branchName, *tagName, *commit)
	}
	file, err := os.Open(*tarFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	gr, err := gzip.NewReader(file)
	if err != nil {
		return err
	}
	defer gr.Close()

	tr := tar.NewReader(gr)
	for {
		header, err := tr.Next()

		if err == io.EOF {
			break
		}

		if err != nil {

			return err
		}
		fileName := htmlPath + header.Name
		switch header.Typeflag {
		case tar.TypeDir:
			continue
		case tar.TypeReg:
			outFile, err := fh.createFile(fileName)
			if err != nil {
				log.Println(err)
			}
			defer outFile.Close()
			if _, err := io.Copy(outFile, tr); err != nil {
				log.Println(err)
			}
		default:
			log.Printf(
				"ExtractTarGz: uknown type: %v in %s",
				header.Typeflag,
				header.Name)
		}
	}
	return nil
}
