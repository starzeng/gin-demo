package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	zapgorm2 "moul.io/zapgorm2"
	"starzeng.com/gin-demo/config"
	logger2 "starzeng.com/gin-demo/pkg/logger"
)

var DB *gorm.DB

func InitMySQL() {
	cfg := config.AppConfig.MySQL

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)

	// 添加zap日志
	zapLogger := zapgorm2.New(logger2.Log)
	zapLogger.SetAsDefault()

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: zapLogger,
	})
	if err != nil {
		panic("MySQL连接失败: " + err.Error())
	}

	// 设置连接池参数
	sqlDB, err := DB.DB()
	if err != nil {
		panic("数据库连接失败: " + err.Error())
	}

	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)

	if cfg.LogDebug {
		DB = DB.Debug()
	}

	fmt.Println("MySQL连接成功")
}
