package pageQuery

import (
	"fmt"
	"math"
)

type Page struct {
	PageSize  int
	CurPage   int
	TotalPage int
	TotalItem int
}

func NewPage(curPage, pageSize, totalItem int) *Page {
	return &Page{
		CurPage:   curPage,
		PageSize:  pageSize,
		TotalItem: totalItem,
		TotalPage: int(math.Ceil(float64(totalItem) / float64(pageSize))),
	}
}

// IsHasPrev 判断是否有上一页
func (p *Page) IsHasPrev() bool {
	return p.CurPage > 1
}

// IsHasNext 判断是否有下一页
func (p *Page) IsHasNext() bool {
	return p.CurPage < p.TotalPage
}

// GetPrevPageNo 获取上一页
func (p *Page) GetPrevPageNo() int {
	if p.IsHasPrev() {
		return p.CurPage - 1
	}
	fmt.Println(p.CurPage)
	return 1
}

// GetNextPageNo 获取下一页
func (p *Page) GetNextPageNo() int {
	if p.IsHasNext() {
		return p.CurPage + 1
	}
	fmt.Println(p.CurPage)
	return p.CurPage
}
