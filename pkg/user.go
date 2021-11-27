package todo

type User struct {
	Id       int    `json:"-" db:"id"`
	Name     string `json:"name" binding:"required"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserList struct {
	Id     int
	UserId int
	ListId int
}
