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

type FilePermission struct {
	handler.HTTPHandler
	repo repository.IRepository
}

func NewFilePermissionHandler(conn *sql.DB) *FilePermission {
	return &FilePermission{
		repo: user.NewFilePermissionRepository(conn),
	}
}

func (filepermission *FilePermission) GetHTTPHandler() []*handler.HTTPHandler {
	return []*handler.HTTPHandler{
		&handler.HTTPHandler{Authenticated: false, Method: http.MethodGet, Path: "filepermission/{id}", Func: filepermission.GetByID},
		&handler.HTTPHandler{Authenticated: false, Method: http.MethodPost, Path: "filepermission", Func: filepermission.Create},
		&handler.HTTPHandler{Authenticated: false, Method: http.MethodPut, Path: "filepermission/{id}", Func: filepermission.Update},
		&handler.HTTPHandler{Authenticated: false, Method: http.MethodDelete, Path: "filepermission/{id}", Func: filepermission.Delete},
		&handler.HTTPHandler{Authenticated: false, Method: http.MethodGet, Path: "filepermission", Func: filepermission.GetAll},
		&handler.HTTPHandler{Authenticated: false, Method: http.MethodPatch, Path: "filepermission/{id}", Func: filepermission.Update},
	}
}
func (filepermission *FilePermission) GetByID(w http.ResponseWriter, r *http.Request) {
	var grp interface{}
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	for {
		if nil != err {
			break
		}
		
		grp, err = filepermission.repo.GetByID(r.Context(), id)
		break
	}
	handler.WriteJSONResponse(w, r, grp, http.StatusOK, err)
}

func (filepermission *FilePermission) Create(w http.ResponseWriter, r *http.Request) {
	var grp model.FilePermission	
	err := json.NewDecoder(r.Body).Decode(&grp)
	for {
		if nil != err {
			break
		}
		_, err = filepermission.repo.Create(r.Context(), grp)
		break
	}
	handler.WriteJSONResponse(w, r, grp, http.StatusOK, err)
}

func (filepermission *FilePermission) Update(w http.ResponseWriter, r *http.Request) {
	var iUsr interface{}
	id,_ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	grp := model.FilePermission{}
	err := json.NewDecoder(r.Body).Decode(&grp)
	for {
		if nil != err {
			break
		}
		grp.Id = id
		if nil != err {
			break
		}
		iUsr, err = filepermission.repo.Update(r.Context(), grp)
			if nil != err {
				break
			}
			grp = iUsr.(model.FilePermission)
			break
		}
		handler.WriteJSONResponse(w, r, grp, http.StatusOK, err)
}

func (filepermission *FilePermission) Delete(w http.ResponseWriter, r *http.Request) {
	var payload string
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	for {
		if nil != err {
			break
		}
		err = filepermission.repo.Delete(r.Context(), id)
		if nil != err {
			break
		}
		payload = "FilePermission deleted successfully"
		break
	}
	handler.WriteJSONResponse(w, r, payload, http.StatusOK, err)
}

func (filepermission *FilePermission) GetAll(w http.ResponseWriter, r *http.Request) {
	grps, err := filepermission.repo.GetAll(r.Context())
	handler.WriteJSONResponse(w, r, grps, http.StatusOK, err)
}