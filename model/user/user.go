package user

type UserBody struct {
	Name string `json:"name" binding:"required"`
	Age int `json:"age" binding:"required"`
}

type UserDB struct {
	Name string `db:"name" json:"name"`
	Age int `db:"age" json:"age"`
}