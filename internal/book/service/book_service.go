package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"starzeng.com/gin-demo/internal/book/model"
	"starzeng.com/gin-demo/internal/book/repository"
	"starzeng.com/gin-demo/pkg/logger"
	"starzeng.com/gin-demo/pkg/redis"
	"strconv"
	"time"
)

func CreateBook(c *gin.Context, book model.Book) error {
	return repository.Create(book)
}

func ListBook(c *gin.Context, bookQuery model.BookQuery) ([]model.Book, int64, error) {
	return repository.List(bookQuery)
}

func GetBook(c *gin.Context, id uint64) (*model.Book, error) {
	logger.Info("开始 get Gook service")

	key := "id:" + strconv.FormatUint(id, 10)
	_ = redis.Set(key, id, 0)

	return repository.GetById(id)
}

func UpdateBook(c *gin.Context, book model.Book) error {
	// 查询数据
	exist, err := repository.GetById(book.ID)
	if err != nil {
		return err
	}

	if exist == nil {
		return errors.New("书本不存在")
	}

	if exist.Version == book.Version {
		return errors.New("乐观锁冲突")
	}

	// 属性赋值（判断指针是否为 nil）
	if book.Title != nil {
		exist.Title = book.Title
	}
	if book.Author != nil {
		exist.Author = book.Author
	}
	if book.Price != nil {
		exist.Price = book.Price
	}
	if book.Description != nil {
		exist.Description = book.Description
	}

	exist.Version += 1
	exist.UpdatedAt = time.Now()

	return repository.Update(exist)
}

func DeleteBook(c *gin.Context, id uint64) error {
	return repository.Delete(id)
}
