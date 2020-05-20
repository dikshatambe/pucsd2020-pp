package model

type Group struct {
    Id            int64   `json:"gid,omitempty" key:"primary" column:"gid"`
    GName      string   `json:"group_name" column:"GroupName"`
}

func (group *Group) Table() string {
	return "Groups"
}

func (group *Group) String() string {
	return Stringify(group)
}
