package http

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"

	"github.com/pucsd2020-pp/rest-api/handler"
	"github.com/pucsd2020-pp/rest-api/model"
	"github.com/pucsd2020-pp/rest-api/repository"
	"github.com/pucsd2020-pp/rest-api/repository/user"
)

type File struct {
	handler.HTTPHandler
	repo repository.IRepository
}

func NewFileHandler(conn *sql.DB) *File {
	return &File{
		repo: user.NewFileRepository(conn),
	}
}

func (file *File) GetHTTPHandler() []*handler.HTTPHandler {
	return []*handler.HTTPHandler{
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodGet, Path: "file/{id}", Func: file.GetByID},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodPost, Path: "file", Func: file.Create},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodPut, Path: "file/{id}", Func: file.Update},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodDelete, Path: "file/{id}", Func: file.Delete},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodGet, Path: "file", Func: file.GetAll},
	}
}

func (file *File) GetByID(w http.ResponseWriter, r *http.Request) {
	var fl interface{}
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	for {
		if nil != err {
			break
		}

		fl, err = file.repo.GetByID(r.Context(), id)
		break
	}

	handler.WriteJSONResponse(w, r, fl, http.StatusOK, err)
}

func (file *File) Create(w http.ResponseWriter, r *http.Request) {
	var fl model.File
	err := json.NewDecoder(r.Body).Decode(&fl)
	for {
		if nil != err {
			break
		}

		_, err = file.repo.Create(r.Context(), fl)
		break
	}
	handler.WriteJSONResponse(w, r, fl, http.StatusOK, err)
}

func (file *File) Update(w http.ResponseWriter, r *http.Request) {
	var iUsr interface{}
	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	fl := model.File{}
	err := json.NewDecoder(r.Body).Decode(&fl)
	for {
		if nil != err {
			break
		}
		fl.Id = id
		if nil != err {
			break
		}

		// set logged in file id for tracking update
		fl.UpdatedBy = 0

		iUsr, err = file.repo.Update(r.Context(), fl)
		if nil != err {
			break
		}
		fl = iUsr.(model.File)
		break
	}

	handler.WriteJSONResponse(w, r, fl, http.StatusOK, err)
}

func (file *File) Delete(w http.ResponseWriter, r *http.Request) {
	var payload string
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	for {
		if nil != err {
			break
		}

		err = file.repo.Delete(r.Context(), id)
		if nil != err {
			break
		}
		payload = "File deleted successfully"
		break
	}

	handler.WriteJSONResponse(w, r, payload, http.StatusOK, err)
}

func (file *File) GetAll(w http.ResponseWriter, r *http.Request) {
	fls, err := file.repo.GetAll(r.Context())
	handler.WriteJSONResponse(w, r, fls, http.StatusOK, err)
}
