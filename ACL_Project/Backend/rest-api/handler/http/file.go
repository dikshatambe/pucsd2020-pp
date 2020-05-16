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
		&handler.HTTPHandler{Authenticated: false, Method: http.MethodGet, Path: "file/{id}", Func: file.GetByID},
		&handler.HTTPHandler{Authenticated: false, Method: http.MethodPost, Path: "file", Func: file.Create},
		&handler.HTTPHandler{Authenticated: false, Method: http.MethodPut, Path: "file/{id}", Func: file.Update},
		&handler.HTTPHandler{Authenticated: false, Method: http.MethodDelete, Path: "file/{id}", Func: file.Delete},
		&handler.HTTPHandler{Authenticated: false, Method: http.MethodGet, Path: "file", Func: file.GetAll},
		&handler.HTTPHandler{Authenticated: false, Method: http.MethodPatch, Path: "file/{id}", Func: file.Update},
	}
}
func (file *File) GetByID(w http.ResponseWriter, r *http.Request) {
	var grp interface{}
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	for {
		if nil != err {
			break
		}
		
		grp, err = file.repo.GetByID(r.Context(), id)
		break
	}
	handler.WriteJSONResponse(w, r, grp, http.StatusOK, err)
}

func (file *File) Create(w http.ResponseWriter, r *http.Request) {
	var grp model.File	
	err := json.NewDecoder(r.Body).Decode(&grp)
	for {
		if nil != err {
			break
		}
		_, err = file.repo.Create(r.Context(), grp)
		break
	}
	handler.WriteJSONResponse(w, r, grp, http.StatusOK, err)
}

func (file *File) Update(w http.ResponseWriter, r *http.Request) {
	var iUsr interface{}
	id,_ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	grp := model.File{}
	err := json.NewDecoder(r.Body).Decode(&grp)
	for {
		if nil != err {
			break
		}
		grp.Id = id
		if nil != err {
			break
		}
		iUsr, err = file.repo.Update(r.Context(), grp)
			if nil != err {
				break
			}
			grp = iUsr.(model.File)
			break
		}
		handler.WriteJSONResponse(w, r, grp, http.StatusOK, err)
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
	grps, err := file.repo.GetAll(r.Context())
	handler.WriteJSONResponse(w, r, grps, http.StatusOK, err)
}
