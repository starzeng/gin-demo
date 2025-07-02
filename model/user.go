package model

type User struct {
	Username   string
	Password   string
	Role       string
	Permission []string
}

var Users = map[string]User{
	"admin": {"admin", "123456", "admin", []string{"read", "write", "delete"}},
	"jack":  {"jack", "123456", "user", []string{"read"}},
}
