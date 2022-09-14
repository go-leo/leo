package pagex

import (
	"errors"
	"math"
)

type Page struct {
	/**
	 * 页码，从1开始
	 */
	pageNum uint64
	/**
	 * 页面大小
	 */
	pageSize uint64
	/**
	 * 起始行
	 */
	startRow uint64
	/**
	 * 末行
	 */
	endRow uint64
	/**
	 * 总数
	 */
	total uint64
	/**
	 * 总页数
	 */
	pages uint64
	/**
	 * 包含count查询
	 */
	count bool
	/**
	 * 进行count查询的列名
	 */
	countColumn string
	/**
	 * 排序
	 */
	orderBy string
	/**
	 * 只增加排序
	 */
	orderByOnly bool
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

func OrderByOnly(orderByOnly bool) Option {
	return func(p *Page) {
		p.orderByOnly = orderByOnly
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
		startRow:    0,
		endRow:      0,
		total:       0,
		pages:       0,
		count:       true,
		countColumn: "",
		orderBy:     "",
		orderByOnly: false,
	}
	p.apply(opts...)
	p.init()
	// 查询全部
	if p.pageNum == 1 && p.pageSize == math.MaxUint64 {
		p.startRow = 0
		p.endRow = math.MaxUint64
		return p, nil
	}
	// 计算出起始行与结束行
	p.startRow = (p.pageNum - 1) * p.pageSize
	p.endRow = p.startRow + p.pageSize
	return p, nil
}

func (p *Page) PageNum() uint64 {
	return p.pageNum
}

func (p *Page) PageSize() uint64 {
	return p.pageSize
}

func (p *Page) StartRow() uint64 {
	return p.startRow
}

func (p *Page) EndRow() uint64 {
	return p.endRow
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

func (p *Page) SetPages(pages uint64) {
	p.pages = pages
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

func (p *Page) OrderByOnly() bool {
	return p.orderByOnly
}
