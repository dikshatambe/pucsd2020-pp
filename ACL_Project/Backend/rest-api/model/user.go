package model

type User struct {
	Id            int64  `json:"id,omitempty" key:"primary" column:"id"`
	UserName      string `json:"user_name" column:"UserName"`
	Password      string `json:"password" column:"password"`
}

func (user *User) Table() string {
	return "Users"
}

func (user *User) String() string {
	return Stringify(user)
}
