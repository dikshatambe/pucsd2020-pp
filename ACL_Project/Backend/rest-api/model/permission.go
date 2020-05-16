package model

type Permission struct {
    Id		  int64   `json:"p_id" column:"pid"`
    PName 	  string  `json:"p_name" column:"PName"`
}

func (permission *Permission) Table() string {
	return "permissions"
}

func (permission *Permission) String() string {
	return Stringify(permission)
}