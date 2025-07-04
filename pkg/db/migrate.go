package db

import (
	"gorm.io/gorm"
	"starzeng.com/gin-demo/internal/book/model"
	model2 "starzeng.com/gin-demo/internal/user/model"
)

// AutoMigrate 自动迁移数据库模式。
// 该函数接受一个*gorm.DB实例作为参数，该实例包含了数据库连接信息。
// 函数将尝试根据定义的模型自动迁移数据库模式。
// 如果迁移过程中发生错误，该函数将返回错误。
func AutoMigrate(db *gorm.DB) error {
	// 调用db.AutoMigrate方法来执行自动迁移。
	// 这里没有提供具体的模型，因为模型应该在调用此函数之前已经注册到db实例中。
	// 如果需要为多个模型执行自动迁移，可以在实际调用中提供这些模型作为参数。
	return db.AutoMigrate(
		&model2.User{},
		&model.Book{},
	)
}
