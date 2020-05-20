package model

type FilePermission struct {
	Id int64 `json:"user_id" column:"id"`
	FId int64 `json:"file_id" column:"fid"`
	PId int64 `json:"permission_id" column:"pid"`
}

func (filepermission *FilePermission) Table() string {
	return "FilePermission"
}

func (filepermission *FilePermission) String() string {
	return Stringify(filepermission)
}