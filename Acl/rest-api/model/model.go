package model

// prototype of db model
/*type IModel interface {
	Table() string
	String() string
}*/

type UserIModel interface {
	UserTable() string
	String() string
}
