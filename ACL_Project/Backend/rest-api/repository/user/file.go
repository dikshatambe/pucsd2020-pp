package user

import (
	"context"
	"database/sql"
	"github.com/pucsd2020-pp/rest-api/driver"
	"github.com/pucsd2020-pp/rest-api/model"
	"os"
	"path/filepath"
)

type fileRepository struct {
	conn *sql.DB
}

func NewFileRepository(conn *sql.DB) *fileRepository {
	return &fileRepository{conn: conn}
}

func (file *fileRepository) GetByID(cntx context.Context, id int64) (interface{}, error) {
	obj := new(model.File)
	/*lines to get file content from server*/
	return driver.GetById(file.conn, obj, id)
}

func (file *fileRepository) Create(cntx context.Context, obj interface{}) (interface{}, error) {
	usr := obj.(model.File)

	basepath := usr.FilePath
	filename := usr.FileName
	dst, _ := os.Create(filepath.Join(basepath, filename, "/"))
	defer dst.Close()

	result, err := driver.Create(file.conn, &usr)
	if nil != err {
		return 0, err
	}

	id, _ := result.LastInsertId()
	usr.Id = id
	return id, nil
}

func (file *fileRepository) Update(cntx context.Context, obj interface{}) (interface{}, error) {
	usr := obj.(model.File)
	/*lines to add lines in file on server*/
	err := driver.UpdateById(file.conn, &usr)
	return obj, err
}

func (file *fileRepository) Delete(cntx context.Context, id int64) (interface{}, error) {
	usr := &model.File{Id: id}
	/*Lines to delete file over the server*/
	obj := new(model.File)
	driver.GetById(file.conn, obj, id)

	basepath := obj.FilePath
	filename := obj.FileName
	os.Remove(filepath.Join(basepath, filename, "/"))

	return driver.DeleteById(file.conn, usr, id)
}

func (file *fileRepository) GetAll(cntx context.Context) ([]interface{}, error) {
	obj := &model.File{}
	return driver.GetAll(file.conn, obj, 0, 0)
}
