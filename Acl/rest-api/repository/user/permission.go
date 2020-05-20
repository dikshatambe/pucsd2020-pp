package user

import (
	"context"
	"database/sql"

	"github.com/pucsd2020-pp/rest-api/driver"
	"github.com/pucsd2020-pp/rest-api/model"
)

type permissionRepository struct {
	conn *sql.DB
}

func NewPermissionRepository(conn *sql.DB) *permissionRepository {
	return &permissionRepository{conn: conn}
}

func (permission *permissionRepository) GetByID(cntx context.Context, id int64) (interface{}, error) {
	obj := new(model.Permission)
	return driver.GetById(permission.conn, obj, id)
}

func (permission *permissionRepository) Create(cntx context.Context, obj interface{}) (interface{}, error) {
	perm := obj.(model.Permission)
	result, err := driver.Create(permission.conn, &perm)
	if nil != err {
		return 0, err
	}

	id, _ := result.LastInsertId()
	perm.Id = id
	return id, nil
}

func (permission *permissionRepository) Update(cntx context.Context, obj interface{}) (interface{}, error) {
	perm := obj.(model.Permission)
	err := driver.UpdateById(permission.conn, &perm)
	return obj, err
}

func (permission *permissionRepository) Delete(cntx context.Context, id int64) error {
	obj := &model.Permission{Id: id}
	return driver.SoftDeleteById(permission.conn, obj, id)
}

func (permission *permissionRepository) GetAll(cntx context.Context) ([]interface{}, error) {
	obj := &model.Permission{}
	return driver.GetAll(permission.conn, obj, 0, 0)
}
