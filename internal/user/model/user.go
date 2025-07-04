package model

type User struct {
	Username   string   `gorm:"size:64;not null;uniqueIndex" json:"username" comment:"用户名，唯一索引"`
	Password   string   `gorm:"size:256;not null" json:"password,omitempty" comment:"密码，存储哈希值，敏感字段"`
	Role       string   `gorm:"size:32;not null" json:"role" comment:"用户角色，比如 admin、user"`
	Permission []string `gorm:"type:json" json:"permission" comment:"权限列表，JSON格式存储"`
}

var Users = map[string]User{
	"admin": {"admin", "123456", "admin", []string{"read", "write", "delete"}},
	"jack":  {"jack", "123456", "user", []string{"read"}},
}
