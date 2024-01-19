package models

type User struct {
	ID       uint
	Name     string
	Password string
	IsAdmin  bool
}

type Post struct {
	ID      uint
	Title   string
	Сontent string
}
