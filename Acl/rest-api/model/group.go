package model

type Group struct {
	Id            int64  `json:"id,omitempty" key:"primary" autoincr:"1" column:"GroupId"`
	GroupName     string `json:"GroupName" column:"GroupName"`
	//LastName      string `json:"last_name" column:"last_name"`
	Owner         string `json:"Owner" column:"Owner"`
	//password      string `json:"password" column:"password"`
}

func (group *Group) Table() string {
	return "group_detail"
}

func (group *Group) String() string {
	return Stringify(group)
}
