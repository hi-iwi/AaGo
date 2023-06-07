package atype

type Paging struct {
	Page   uint `json:"page"`
	Offset uint `json:"offset"`
	Limit  uint `json:"limit"`
	Prev   uint `json:"prev"`
	Next   uint `json:"next"`
}

// @param args[0] firstPageLimit 首页行数
// @param args[1] limitMax 其他页行数
func NewPaging(page uint, args ...uint) Paging {
	limit := uint(10)
	firstPageLimit := limit
	if len(args) > 0 {
		firstPageLimit = args[0]
		if len(args) > 1 {
			limit = args[1]
		}
	}
	var offset uint
	var prev uint
	if page <= 1 {
		page = 1
		limit = firstPageLimit
	} else {
		offset = firstPageLimit + (page-2)*limit
		prev = page - 1
	}
	return Paging{
		Page:   page,
		Offset: offset,
		Limit:  limit,
		Prev:   prev,
		Next:   page + 1,
	}

}
