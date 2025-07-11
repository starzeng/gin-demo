package repository

import (
	"errors"
	"gorm.io/gorm"
	"starzeng.com/gin-demo/internal/book/model"
	"starzeng.com/gin-demo/pkg/db"
	"starzeng.com/gin-demo/pkg/logger"
)

func Create(book model.Book) error {
	return db.DB.Create(&book).Error
}

func List(bookQuery model.BookQuery) ([]model.Book, int64, error) {
	var bookList []model.Book
	var total int64

	offset := bookQuery.GetOffset()
	limit := bookQuery.PageSize
	author := bookQuery.Author
	title := bookQuery.Title

	query := db.DB.Model(&model.Book{})
	if title != "" {
		query.Where("title like ?", title+"%")
	}
	if author != "" {
		query.Where("author like ?", author+"%")
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.Offset(offset).Limit(limit).Find(&bookList).Error

	return bookList, total, nil
}

func GetById(id uint64) (*model.Book, error) {
	logger.Info("开始 get Gook repository")
	var book *model.Book
	err := db.DB.First(&book, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return book, err
}

func Update(book *model.Book) error {
	return db.DB.Save(book).Error
}

func Delete(id uint64) error {
	return db.DB.Delete(&model.Book{}, id).Error
}
