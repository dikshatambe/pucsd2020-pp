package user

import (
	"context"
	"database/sql"

	"github.com/pucsd2020-pp/rest-api/driver"
	"github.com/pucsd2020-pp/rest-api/model"
)

type fileRepository struct {
	conn *sql.DB
}

func NewFileRepository(conn *sql.DB) *fileRepository {
	return &fileRepository{conn: conn}
}

func (file *fileRepository) GetByID(cntx context.Context, id int64) (interface{}, error) {
	obj := new(model.File)
	return driver.GetById(file.conn, obj, id)
}

func (file *fileRepository) Create(cntx context.Context, obj interface{}) (interface{}, error) {
	fl := obj.(model.File)
	result, err := driver.Create(file.conn, &fl)
	if nil != err {
		return 0, err
	}

	id, _ := result.LastInsertId()
	fl.Id = id
	return id, nil
}

func (file *fileRepository) Update(cntx context.Context, obj interface{}) (interface{}, error) {
	fl := obj.(model.File)
	err := driver.UpdateById(file.conn, &fl)
	return obj, err
}

func (file *fileRepository) Delete(cntx context.Context, id int64) error {
	obj := &model.File{Id: id}
	return driver.SoftDeleteById(file.conn, obj, id)
}

func (file *fileRepository) GetAll(cntx context.Context) ([]interface{}, error) {
	obj := &model.File{}
	return driver.GetAll(file.conn, obj, 0, 0)
}
