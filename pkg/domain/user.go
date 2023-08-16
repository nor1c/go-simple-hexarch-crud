package domain

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type UserRepository interface {
	GetAllUsers() (*[]User, error)
}
