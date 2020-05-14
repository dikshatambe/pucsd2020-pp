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

type UserGroup struct {
	handler.HTTPHandler
	repo repository.IRepository
}

func NewUserGroupHandler(conn *sql.DB) *UserGroup {
	return &UserGroup{
		repo: user.NewUserGroupRepository(conn),
	}
}

func (usergroup *UserGroup) GetHTTPHandler() []*handler.HTTPHandler {
	return []*handler.HTTPHandler{
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodGet, Path: "usergroup/{id}", Func: usergroup.GetByID},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodPost, Path: "usergroup", Func: usergroup.Create},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodPut, Path: "usergroup/{id}", Func: usergroup.Update},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodDelete, Path: "usergroup/{id}", Func: usergroup.Delete},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodGet, Path: "usergroup", Func: usergroup.GetAll},
	}
}

func (usergroup *UserGroup) GetByID(w http.ResponseWriter, r *http.Request) {
	var usr interface{}
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	for {
		if nil != err {
			break
		}

		usr, err = usergroup.repo.GetByID(r.Context(), id)
		break
	}

	handler.WriteJSONResponse(w, r, usr, http.StatusOK, err)
}

func (usergroup *UserGroup) Create(w http.ResponseWriter, r *http.Request) {
	var usr model.UserGroup
	err := json.NewDecoder(r.Body).Decode(&usr)
	for {
		if nil != err {
			break
		}

		_, err = usergroup.repo.Create(r.Context(), usr)
		break
	}
	handler.WriteJSONResponse(w, r, usr, http.StatusOK, err)
}

func (usergroup *UserGroup) Update(w http.ResponseWriter, r *http.Request) {
	var iUsr interface{}
	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	usr := model.UserGroup{}
	err := json.NewDecoder(r.Body).Decode(&usr)
	for {
		if nil != err {
			break
		}
		usr.Id = id
		if nil != err {
			break
		}

		// set logged in usergroup id for tracking update
		usr.UpdatedBy = 0

		iUsr, err = usergroup.repo.Update(r.Context(), usr)
		if nil != err {
			break
		}
		usr = iUsr.(model.UserGroup)
		break
	}

	handler.WriteJSONResponse(w, r, usr, http.StatusOK, err)
}

func (usergroup *UserGroup) Delete(w http.ResponseWriter, r *http.Request) {
	var payload string
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	for {
		if nil != err {
			break
		}

		err = usergroup.repo.Delete(r.Context(), id)
		if nil != err {
			break
		}
		payload = "UserGroup deleted successfully"
		break
	}

	handler.WriteJSONResponse(w, r, payload, http.StatusOK, err)
}

func (usergroup *UserGroup) GetAll(w http.ResponseWriter, r *http.Request) {
	usrs, err := usergroup.repo.GetAll(r.Context())
	handler.WriteJSONResponse(w, r, usrs, http.StatusOK, err)
}
