package model

type File struct {
    Id            int64   `json:"id,omitempty" key:"primary" column:"fid"`
    FileName      string   `json:"file_name" column:"FileName"`
	Owner         string  `json:"owner_name" column:"Owner"`

}

func (file *File) Table() string {
	return "Files"
}

func (file *File) String() string {
	return Stringify(file)
}