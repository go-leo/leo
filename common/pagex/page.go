package pagex

import (
	"errors"
)

type Page struct {
	// pageNum 页码，从1开始
	pageNum uint64
	// pageSize 页面大小
	pageSize uint64
	// offset 跳过的行数
	offset uint64
	// limit 限制行数
	limit uint64
	// total 总行数
	total uint64
	// pages 总页数
	pages uint64
	// count 包含count查询
	count bool
	// countColumn 进行count查询的列名
	countColumn string
	// orderBy 排序
	orderBy string
}

func (p *Page) init() {

}

func (p *Page) apply(opts ...Option) {
	for _, opt := range opts {
		opt(p)
	}
}

type Option func(p *Page)

func Count(count bool) Option {
	return func(p *Page) {
		p.count = count
	}
}

func CountColumn(countColumn string) Option {
	return func(p *Page) {
		p.countColumn = countColumn
	}
}

// OrderBy 设置排序字段
func OrderBy(orderBy string) Option {
	return func(p *Page) {
		// TODO 增加 SQL 注入校验
		p.orderBy = orderBy
	}
}

// UnsafeOrderBy 不安全的设置排序方法，如果从前端接收参数，请自行做好注入校验。
func UnsafeOrderBy(orderBy string) Option {
	return func(p *Page) {
		p.orderBy = orderBy
	}
}

func NewPage(pageNum uint64, pageSize uint64, opts ...Option) (*Page, error) {
	if pageNum == 0 {
		return nil, errors.New("pageNum is zero")
	}
	if pageSize == 0 {
		return nil, errors.New("pageSize is zero")
	}
	p := &Page{
		pageNum:     pageNum,
		pageSize:    pageSize,
		total:       0,
		pages:       0,
		count:       true,
		countColumn: "",
		orderBy:     "",
	}
	p.apply(opts...)
	p.init()
	// 计算出 offset 和 limit
	p.offset = (p.pageNum - 1) * p.pageSize
	p.limit = p.pageSize
	return p, nil
}

func (p *Page) PageNum() uint64 {
	return p.pageNum
}

func (p *Page) PageSize() uint64 {
	return p.pageSize
}

func (p *Page) Offset() uint64 {
	return p.offset
}

func (p *Page) Limit() uint64 {
	return p.limit
}

func (p *Page) Total() uint64 {
	return p.total
}

func (p *Page) SetTotal(total uint64) {
	p.total = total
	if total%p.pageSize == 0 {
		p.pages = total / p.pageSize
	} else {
		p.pages = total/p.pageSize + 1
	}
}

func (p *Page) Pages() uint64 {
	return p.pages
}

func (p *Page) Count() bool {
	return p.count
}

func (p *Page) CountColumn() string {
	return p.countColumn
}

func (p *Page) OrderBy() string {
	return p.orderBy
}
