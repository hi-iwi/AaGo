package atype

type Paging struct {
	Page         uint `json:"page"`
	PageEnd      uint `json:"page_end"`
	PerPageLimit uint `json:"per_page_limit"`
	Offset       uint `json:"offset"`
	Limit        uint `json:"limit"`
	Prev         uint `json:"prev"`
	Next         uint `json:"next"`
}

// @param perPageLimit 每页行数
// @param page 起始页数
// @param pageEnd 截止页数
// @param firstPageEnd  如果page是第一页时，pageEnd为此
func NewPaging(perPageLimit, page, pageEnd, firstPageEnd uint) Paging {
	var offset uint
	var prev uint
	if page <= 1 {
		page = 1
	}
	if pageEnd < page {
		pageEnd = page
	}
	if pageEnd == 1 && firstPageEnd > pageEnd {
		pageEnd = firstPageEnd
	}
	next := pageEnd + 1
	limit := perPageLimit * (next - page)
	offset = (page - 1) * perPageLimit
	prev = page - 1

	return Paging{
		Page:         page,
		PageEnd:      pageEnd,
		PerPageLimit: perPageLimit,
		Offset:       offset,
		Limit:        limit,
		Prev:         prev,
		Next:         next,
	}
}

func NewPage(page, pageEnd uint) Paging {
	return NewPaging(10, page, pageEnd, pageEnd)
}
