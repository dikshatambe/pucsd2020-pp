package model

type Permission struct {
	Id            int64  `json:"id,omitempty" key:"primary" autoincr:"1" column:"id"`
	PermissionName     string `json:"first_name" column:"PName"`
}

func (permission *Permission) Table() string {
	return "permission_bit"
}

func (permission *Permission) String() string {
	return Stringify(permission)
}
