package models

type Page struct {
	Books       []*Book //每页图书存放的切片
	PageNo      int     //当前页
	PageSize    int     //每页显示的条数
	TotalPageNo int     //总页数，通过计算得到
	TotalRecord int     //总记录数
	MinPrice    string  //最小价格
	MaxPrice    string  //最大价格
	IsLogin     bool    //是否登录
	UserName    string  //用户名
}

// 判断是否有上一页
func (p *Page) IsHasPres() bool {
	return p.PageNo > 1
}

// 判断是否有下一页
func (p *Page) IsHasNext() bool {
	return p.PageNo < p.TotalPageNo
}

// 获取上一页
func (p *Page) GetPrevPageNo() int {
	if p.IsHasPres() {
		return p.PageNo - 1
	} else {
		return p.PageNo
	}
}

// 获取下一页
func (p *Page) GetNextPageNo() int {
	if p.IsHasNext() {
		return p.PageNo + 1
	} else {
		return p.PageNo
	}
}
