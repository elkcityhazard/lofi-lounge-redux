package main

import (
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/elkcityhazard/lofi-lounge-redux/internal/models"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func (app *application) writeJSON(w http.ResponseWriter, status int, data interface{}, wrap string) error {

	wrapper := make(map[string]interface{})

	wrapper[wrap] = data

	js, err := json.Marshal(wrapper)

	if err != nil {
		return err
	}

	fmt.Println("working")
	_, err = w.Write(js)

	if err != nil {
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return err
	}

	return nil
}

func (app *application) errorJSON(w http.ResponseWriter, err error) {
	type jsonError struct {
		Message string `json:"message"`
	}

	theError := jsonError{
		Message: err.Error(),
	}

	app.writeJSON(w, http.StatusBadRequest, theError, "error")
}

func (app *application) GenerateErrorJSON(method string, status int, errMsg string, w http.ResponseWriter, r *http.Request) {
	fmt.Println("Fired")
	// This needs to be a post request
	if r.Method != method {

		msg := errors.New(errMsg)

		errStruct := models.Error{
			Status:  status,
			Message: msg,
		}

		err := app.writeJSON(w, http.StatusMethodNotAllowed, errStruct, "error")

		if err != nil {
			http.Error(w, "error processing request", http.StatusInternalServerError)
			return
		}
	}
}

func (app *application) RandomString(n int) string {
	const randomStringSource = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_+_+"
	s, r := make([]rune, n), []rune(randomStringSource)

	for i := range s {
		p, _ := rand.Prime(rand.Reader, len(r))
		x, y := p.Uint64(), uint64(len(r))
		s[i] = r[x%y]
	}

	return string(s)

}

func (app *application) UploadSingleFile(r *http.Request, uploadDir string, rename ...bool) (*models.UploadedFile, error) {
	// handle whether or not the file is being renamed
	renameFile := true

	//	since rename is a variatic bool it stored in a slice of bool

	if len(rename) > 0 {
		renameFile = rename[0]
	}

	files, err := app.UploadFiles(r, uploadDir, renameFile)

	if err != nil {
		return nil, err
	}

	if len(files) > 1 {
		err = errors.New("error: trying to upload too many files")
		return nil, err
	}

	return files[0], nil
}

func (app *application) UploadFiles(r *http.Request, uploadDir string, rename ...bool) ([]*models.UploadedFile, error) {

	// handle whether or not the file is being renamed
	renameFile := true

	//	since rename is a variatic bool it stored in a slice of bool

	if len(rename) > 0 {
		renameFile = rename[0]
	}

	//	create a slice to hold our uploaded files

	var uploadedFiles []*models.UploadedFile

	//	 set a default max file size if there is none

	if app.maxFileSize == 0 {
		app.maxFileSize = 1024 * 1024 * 1024
	}

	err := app.CreateDirIfNotExists(uploadDir)

	if err != nil {
		return nil, err
	}

	//	parse the multipart form which contains the file

	err = r.ParseMultipartForm(int64(app.maxFileSize))

	// check the error from parsing the multipart form

	if err != nil {
		return nil, errors.New("the uploaded file is too big")
	}

	//	here we are going to be checking the headers of each uploaded file
	//	from the multipart form
	//	the outer loop is ranging through each uploaded file
	//	 the inner loop is ranging through the headers of the file
	//	*any time you are deferring things in a loop you need to wrap the contents of the loop in a function

	for _, fHeaders := range r.MultipartForm.File {
		for _, hdr := range fHeaders {
			uploadedFiles, err = func(uploadedFiles []*models.UploadedFile) ([]*models.UploadedFile, error) {

				var uploadedFile models.UploadedFile

				infile, err := hdr.Open()

				if err != nil {
					return nil, err
				}

				defer infile.Close()

				//	create a buffer of 512 bytes
				//	so we can look at the file
				//	and determine the mime/type

				buff := make([]byte, 512)

				//	read the file into the buff so we can read
				//	 the first 512 bytes

				_, err = infile.Read(buff)

				if err != nil {
					return nil, err
				}

				//	TODO:  Check to see if the file type is permitted

				allowed := false

				fileType := http.DetectContentType(buff)

				fmt.Println("FileType", fileType)

				//	 if there are allowed types, use strings.EqualFold to check to make sure
				//	 that  the filetype is included in the allowed filetypes

				if len(app.allowedFileTypes) > 0 {
					for _, x := range app.allowedFileTypes {
						if strings.EqualFold(fileType, x) {
							allowed = true
						}
					}
				} else {
					allowed = true
				}

				if !allowed {
					return nil, errors.New("the uploaded file type is not permitted")
				}

				//	since we read the first 512 bytes to get the mime type
				//	 we have to go back to the start of the file or else
				// it will fail saving it correctly

				_, err = infile.Seek(0, 0)

				if err != nil {
					return nil, err
				}

				// handle renaming the file

				if renameFile {
					uploadedFile.OriginalFilename = fmt.Sprintf("%s", filepath.Base(hdr.Filename))
					uploadedFile.NewFileName = fmt.Sprintf("%s%s", app.RandomString(25), filepath.Ext(hdr.Filename))

				} else {
					uploadedFile.NewFileName = hdr.Filename
				}

				// save the file to the disk

				var outfile *os.File
				defer outfile.Close()

				outfile, err = os.Create(filepath.Join(uploadDir, uploadedFile.NewFileName))
				if err != nil {
					return nil, err
				} else {
					fileSize, err := io.Copy(outfile, infile)

					if err != nil {
						return nil, err
					}
					uploadedFile.FileSize = fileSize
				}

				uploadedFiles = append(uploadedFiles, &uploadedFile)
				return uploadedFiles, nil
			}(uploadedFiles)
			if err != nil {
				return uploadedFiles, err
			}
		}
	}
	fmt.Println(uploadedFiles)
	return uploadedFiles, nil
}

// CreateDirIfNotExists checks to see if a directory exists and creates it if it does not exist
func (app *application) CreateDirIfNotExists(path string) error {
	const mode = 0755

	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, mode)

		if err != nil {
			return err
		}
	}
	return nil
}
