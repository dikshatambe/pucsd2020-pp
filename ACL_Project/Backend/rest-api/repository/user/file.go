package user

import (	
	"os"	
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
	/*lines to get file content from server*/
	return driver.GetById(file.conn, obj, id)
}

/*func (file *fileRepository) Create(cntx context.Context, obj interface{}) (interface{}, error) {
	usr := obj.(model.File)
	//lines to create file on server
	result, err := driver.Create(file.conn, &usr)
	if nil != err {
		return 0, err
	}

	id, _ := result.LastInsertId()
	usr.Id = id
	return id, nil
}*/

func (file*fileRepository) Create(cntx context.Context, obj interface{}) (interface{}, error) {
	usr := obj.(model.File)
	if usr.FileTypeId == 1 {
	basepath := usr.FilePath
	filename := usr.FileName
	dst,_ := os.Create(filepath.Join(basepath, filename, "/"))
	defer dst.Close()
	}
	if usr.FileTypeId == 2 {
	basepath := usr.FilePath
	foldername := usr.FileName
	_, err := os.Stat(filepath.Join(basepath, foldername, "/"))
	if os.IsNotExist(err) {
	errDir := os.MkdirAll(usr.FilePath+"/"+usr.FileName, 0755)
	if errDir != nil {
	return 0, err
	}
	}
	
	}
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

func (file *fileRepository) Delete(cntx context.Context, id int64) error {
	obj := &model.File{Id: id}
	/*Lines to delete file over the server*/
	return driver.SoftDeleteById(file.conn, obj, id)
}

func (file *fileRepository) GetAll(cntx context.Context) ([]interface{}, error) {
	obj := &model.File{}
	return driver.GetAll(file.conn, obj, 0, 0)
}
