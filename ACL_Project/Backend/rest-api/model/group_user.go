package model

type GroupUser struct {
    Id		  int64   `json:"user_id" column:"id"`
    GroupId		  int64   `json:"group_id" column:"gid"`
}

func (group_user *GroupUser) Table() string {
	return "UserGroup"
}

func (group_user *GroupUser) String() string {
	return Stringify(group_user)
}