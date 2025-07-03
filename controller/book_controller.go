package controller

import (
	"github.com/gin-gonic/gin"
	"starzeng.com/gin-demo/common"
	"starzeng.com/gin-demo/middleware"
	"starzeng.com/gin-demo/model"
	"starzeng.com/gin-demo/router"
	"starzeng.com/gin-demo/service"
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
func DeleteBook(context *gin.Context) {
	idStr := context.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		common.Error(context, common.CodeInvalidParams, err.Error())
		return
	}

	err = service.DeleteBook(uint64(id))
	if err != nil {
		common.Error(context, common.CodeInternalError, err.Error())
		return
	}
	common.Success(context, nil)
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
func UpdateBook(context *gin.Context) {
	idStr := context.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		common.Error(context, common.CodeInvalidParams, err.Error())
		return
	}

	var book model.Book
	err = context.ShouldBindJSON(&book)
	if err != nil {
		common.Error(context, common.CodeInvalidParams, err.Error())
		return
	}

	book.ID = uint64(id)
	err = service.UpdateBook(book)

	if err != nil {
		common.Error(context, common.CodeInternalError, err.Error())
		return
	}

	common.Success(context, nil)
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
func GetBook(context *gin.Context) {
	idStr := context.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		common.Error(context, common.CodeInvalidParams, err.Error())
		return
	}
	book, err := service.GetBook(uint64(id))
	if err != nil {
		common.Error(context, common.CodeInternalError, err.Error())
		return
	}
	common.Success(context, book)
}

// ListBooks 获取书籍列表
// @Summary 获取书籍列表
// @Description 分页获取书籍信息列表
// @Tags 书籍管理
// @Accept json
// @Produce json
// @Param data body model.BookQuery true "查询参数（如分页）"
// @Security BearerAuth
// @Success 200 {object} common.Response{data=model.PageResult{list=[]model.Book}}
// @Failure 400 {object} common.Response
// @Failure 500 {object} common.Response
// @Router /api/book/list [post]
func ListBooks(context *gin.Context) {
	var bookQuery model.BookQuery
	err := context.ShouldBindJSON(&bookQuery)
	if err != nil {
		common.Error(context, common.CodeInvalidParams, err.Error())
		return
	}

	bookList, total, err := service.ListBook(bookQuery)

	if err != nil {
		common.Error(context, common.CodeInternalError, err.Error())
		return
	}

	resp := model.PageResult{
		Total:    total,
		Page:     bookQuery.Page,
		PageSize: bookQuery.PageSize,
		List:     bookList,
	}

	common.Success(context, resp)
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
func CreateBook(context *gin.Context) {
	var book model.Book
	err := context.ShouldBindJSON(&book)
	if err != nil {
		common.Error(context, common.CodeInvalidParams, err.Error())
		return
	}
	err = service.CreateBook(book)
	if err != nil {
		common.Error(context, common.CodeInternalError, err.Error())
		return
	}
	common.Success(context, "创建成功")
}
