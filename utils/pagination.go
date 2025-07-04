package utils

type Pagination struct {
	Page     int `json:"page" binding:"omitempty,min=1"`              // 当前页码，默认1
	PageSize int `json:"page_size" binding:"omitempty,min=1,max=100"` // 每页条数，默认10，最大100
}

type PageResult struct {
	Total    int64       `json:"total"`     // 总记录数
	Page     int         `json:"page"`      // 当前页码
	PageSize int         `json:"page_size"` // 每页条数
	List     interface{} `json:"list"`      // 数据列表，通常是切片，如 []Book
}

// GetOffset 计算偏移量方法，方便传给数据库查询
func (p *Pagination) GetOffset() int {
	if p.Page < 1 {
		p.Page = 1
	}
	if p.PageSize < 1 {
		p.PageSize = 10
	}
	return (p.Page - 1) * p.PageSize
}
