package controller

import (
	"github.com/gin-gonic/gin"
	"starzeng.com/gin-demo/common"
	"starzeng.com/gin-demo/internal/book/model"
	"starzeng.com/gin-demo/internal/book/service"
	"starzeng.com/gin-demo/middleware"
	"starzeng.com/gin-demo/router"
	"starzeng.com/gin-demo/utils"
	"strconv"
)

type bookController struct {
}

func init() {
	router.Register(&bookController{})
}

func (b bookController) RouteRegister(group *gin.RouterGroup) {
	book := group.Group("/book", middleware.JWTAuth())

	book.POST("", CreateBook)
	book.POST("/list", ListBooks)
	book.GET("/:id", GetBook)
	book.PUT("/:id", UpdateBook)
	book.DELETE("/:id", DeleteBook)
}

// DeleteBook 删除书籍
// @Summary 删除书籍
// @Description 根据书籍 ID 删除对应书籍
// @Tags 书籍管理
// @Produce json
// @Param id path int true "书籍ID"
// @Security BearerAuth
// @Success 200 {object} common.Response
// @Failure 400 {object} common.Response
// @Failure 500 {object} common.Response
// @Router /api/book/{id} [delete]
func DeleteBook(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		common.Error(c, common.CodeInvalidParams, err.Error())
		return
	}

	err = service.DeleteBook(c, uint64(id))
	if err != nil {
		common.Error(c, common.CodeInternalError, err.Error())
		return
	}
	common.Success(c, nil)
}

// UpdateBook 更新书籍
// @Summary 更新书籍信息
// @Description 根据书籍 ID 更新书籍数据
// @Tags 书籍管理
// @Accept json
// @Produce json
// @Param id path int true "书籍ID"
// @Param data body model.Book true "更新的书籍信息"
// @Security BearerAuth
// @Success 200 {object} common.Response
// @Failure 400 {object} common.Response
// @Failure 500 {object} common.Response
// @Router /api/book/{id} [put]
func UpdateBook(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		common.Error(c, common.CodeInvalidParams, err.Error())
		return
	}

	var book model.Book
	err = c.ShouldBindJSON(&book)
	if err != nil {
		common.Error(c, common.CodeInvalidParams, err.Error())
		return
	}

	book.ID = uint64(id)
	err = service.UpdateBook(c, book)

	if err != nil {
		common.Error(c, common.CodeInternalError, err.Error())
		return
	}

	common.Success(c, nil)
}

// GetBook 获取书籍详情
// @Summary 获取书籍详情
// @Description 根据书籍 ID 获取详细信息
// @Tags 书籍管理
// @Produce json
// @Param id path int true "书籍ID"
// @Security BearerAuth
// @Success 200 {object} common.Response{data=model.Book}
// @Failure 400 {object} common.Response
// @Failure 500 {object} common.Response
// @Router /api/book/{id} [get]
func GetBook(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		common.Error(c, common.CodeInvalidParams, err.Error())
		return
	}
	book, err := service.GetBook(c, uint64(id))
	if err != nil {
		common.Error(c, common.CodeInternalError, err.Error())
		return
	}
	common.Success(c, book)
}

// ListBooks 获取书籍列表
// @Summary 获取书籍列表
// @Description 分页获取书籍信息列表
// @Tags 书籍管理
// @Accept json
// @Produce json
// @Param data body model.BookQuery true "查询参数（如分页）"
// @Security BearerAuth
// @Success 200 {object} common.Response{data=utils.PageResult{list=[]model.Book}}
// @Failure 400 {object} common.Response
// @Failure 500 {object} common.Response
// @Router /api/book/list [post]
func ListBooks(c *gin.Context) {
	var bookQuery model.BookQuery
	err := c.ShouldBindJSON(&bookQuery)
	if err != nil {
		common.Error(c, common.CodeInvalidParams, err.Error())
		return
	}

	bookList, total, err := service.ListBook(c, bookQuery)

	if err != nil {
		common.Error(c, common.CodeInternalError, err.Error())
		return
	}

	resp := utils.PageResult{
		Total:    total,
		Page:     bookQuery.Page,
		PageSize: bookQuery.PageSize,
		List:     bookList,
	}

	common.Success(c, resp)
}

// CreateBook 创建书籍
// @Summary 创建新书籍
// @Description 根据请求体中的书籍信息创建新书籍
// @Tags 书籍管理
// @Accept json
// @Produce json
// @Param data body model.Book true "书籍信息"
// @Security BearerAuth
// @Success 200 {object} common.Response
// @Failure 400 {object} common.Response
// @Failure 500 {object} common.Response
// @Router /api/book [post]
func CreateBook(c *gin.Context) {
	var book model.Book
	err := c.ShouldBindJSON(&book)
	if err != nil {
		common.Error(c, common.CodeInvalidParams, err.Error())
		return
	}
	err = service.CreateBook(c, book)
	if err != nil {
		common.Error(c, common.CodeInternalError, err.Error())
		return
	}
	common.Success(c, "创建成功")
}
