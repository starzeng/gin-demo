package model

import (
	"gorm.io/gorm"
	"starzeng.com/gin-demo/utils"
	"time"
)

type Book struct {
	ID          uint64         `gorm:"primaryKey;autoIncrement" json:"id" comment:"主键ID，自增"`
	Title       *string        `gorm:"type:varchar(200);not null" json:"title" binding:"required,max=200" comment:"书名"`
	Author      *string        `gorm:"type:varchar(100);not null" json:"author" binding:"required,max=100" comment:"作者"`
	Price       *float64       `gorm:"type:decimal(10,2);not null" json:"price" binding:"required,gt=0" comment:"价格，保留两位小数，必须大于0"`
	Description *string        `gorm:"type:text" json:"description,omitempty" binding:"omitempty,max=1000" comment:"书籍简介，最多1000字符"`
	Version     int            `gorm:"default:1" json:"version" comment:"乐观锁版本号，默认为1"`
	CreatedAt   time.Time      `gorm:"autoCreateTime" json:"created_at" comment:"创建时间"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime" json:"updated_at" comment:"更新时间"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-" comment:"软删除字段，逻辑删除"`
}

type BookQuery struct {
	Title  string `json:"title" binding:"omitempty,max=200"`  // 书名模糊查询
	Author string `json:"author" binding:"omitempty,max=100"` // 作者模糊查询
	utils.Pagination
}
