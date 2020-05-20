package model

type File struct {
	Id            int64  `json:"id,omitempty" key:"primary" autoincr:"1" column:"id"`
	FileName     string `json:"file_name" column:"FileName"`
}

func (file *File) Table() string {
	return "Files"
}

func (file *File) String() string {
	return Stringify(file)
}
