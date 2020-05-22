package model

type User struct {
	Id        int64  `json:"id,omitempty" key:"primary" column:"id"`
	FirstName string `json:"first_name" column:"FirstName"`
	LastName  string `json:"last_name" column:"LastName"`
	UserName  string `json:"user_name" column:"UserName"`
	Password  string `json:"password" column:"password"`
	//IsRoot        int64  `json:"is_root" column:"IsUserRoot"`
}

func (user *User) Table() string {
	return "Users"
}

func (user *User) String() string {
	return Stringify(user)
}
