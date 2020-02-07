package user

type UserBody struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type UserDB struct {
	Name     string `db:"name" json:"name"`
	Age      int    `db:"age" json:"age"`
	Phone    string `db:"phone" json:"phone"`
	Password string `db:"password" json:"password"`
}
