package model

type Group struct {
	Id            int64  `json:"id,omitempty" key:"primary" autoincr:"1" column:"GroupId"`
	GroupName     string `json:"GroupName" column:"GroupName"`
}

func (group *Group) Table() string {
	return "group_detail"
}

func (group *Group) String() string {
	return Stringify(group)
}
