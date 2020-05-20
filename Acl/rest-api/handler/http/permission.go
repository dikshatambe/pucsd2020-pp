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

type Permission struct {
	handler.HTTPHandler
	repo repository.IRepository
}

func NewPermissionHandler(conn *sql.DB) *Permission {
	return &Permission{
		repo: user.NewPermissionRepository(conn),
	}
}

func (permission *Permission) GetHTTPHandler() []*handler.HTTPHandler {
	return []*handler.HTTPHandler{
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodGet, Path: "permission/{id}", Func: permission.GetByID},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodPost, Path: "permission", Func: permission.Create},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodPut, Path: "permission/{id}", Func: permission.Update},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodDelete, Path: "permission/{id}", Func: permission.Delete},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodGet, Path: "permission", Func: permission.GetAll},
	}
}

func (permission *Permission) GetByID(w http.ResponseWriter, r *http.Request) {
	var perm interface{}
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	for {
		if nil != err {
			break
		}

		perm, err = permission.repo.GetByID(r.Context(), id)
		break
	}

	handler.WriteJSONResponse(w, r, perm, http.StatusOK, err)
}

func (permission *Permission) Create(w http.ResponseWriter, r *http.Request) {
	var perm model.Permission
	err := json.NewDecoder(r.Body).Decode(&perm)
	for {
		if nil != err {
			break
		}

		_, err = permission.repo.Create(r.Context(), perm)
		break
	}
	handler.WriteJSONResponse(w, r, perm, http.StatusOK, err)
}

func (permission *Permission) Update(w http.ResponseWriter, r *http.Request) {
	var iUsr interface{}
	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	perm := model.Permission{}
	err := json.NewDecoder(r.Body).Decode(&perm)
	for {
		if nil != err {
			break
		}
		perm.Id = id
		if nil != err {
			break
		}

		// set logged in permission id for tracking update
		perm.UpdatedBy = 0

		iUsr, err = permission.repo.Update(r.Context(), perm)
		if nil != err {
			break
		}
		perm = iUsr.(model.Permission)
		break
	}

	handler.WriteJSONResponse(w, r, perm, http.StatusOK, err)
}

func (permission *Permission) Delete(w http.ResponseWriter, r *http.Request) {
	var payload string
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	for {
		if nil != err {
			break
		}

		err = permission.repo.Delete(r.Context(), id)
		if nil != err {
			break
		}
		payload = "Permission deleted successfully"
		break
	}

	handler.WriteJSONResponse(w, r, payload, http.StatusOK, err)
}

func (permission *Permission) GetAll(w http.ResponseWriter, r *http.Request) {
	perms, err := permission.repo.GetAll(r.Context())
	handler.WriteJSONResponse(w, r, perms, http.StatusOK, err)
}
