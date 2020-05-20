package model

type UserGroup struct {
	Id            int64  `json:"id,omitempty" key:"primary" autoincr:"1" column:"id"`
	UserId        int64  `json:"user_id" column:"UserId"`
}

func (usergroup *UserGroup) Table() string {
	return "usergroup_detail"
}

func (usergroup *UserGroup) String() string {
	return Stringify(usergroup)
}
