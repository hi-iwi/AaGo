package atype

type Paging struct {
	Page    uint `json:"page"`
	PageEnd uint `json:"page_end"`
	Offset  uint `json:"offset"`
	Limit   uint `json:"limit"`
	Prev    uint `json:"prev"`
	Next    uint `json:"next"`
}

// @param perPageLimit 每页行数
// @param page 起始页数
// @param pageEnd 截止页数
func NewPaging(perPageLimit, page, pageEnd uint) Paging {

	var offset uint
	var prev uint
	if page <= 1 {
		page = 1
	}
	if pageEnd < page {
		pageEnd = page
	}
	next := pageEnd + 1
	limit := perPageLimit * (next - page)
	offset = (page - 1) * perPageLimit
	prev = page - 1

	return Paging{
		Page:    page,
		PageEnd: pageEnd,
		Offset:  offset,
		Limit:   limit,
		Prev:    prev,
		Next:    next,
	}

}
