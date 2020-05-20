package model

type File struct {
	Id       int64  `json:"id,omitempty" key:"primary" column:"fid"`
	FileName string `json:"file_name" column:"FileName"`
	FilePath string `json:"file_path" column:" FilePath"`
}

func (file *File) Table() string {
	return "Files"
}

func (file *File) String() string {
	return Stringify(file)
}
