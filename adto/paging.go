package adto

import "github.com/luexu/AaGo/com"

type Paging struct {
	Page     int `alias:"page"`
	PageSize int `alias:"page_size"`
	Offset   int
	Limit    int
}

func MakePaging(r *com.Req) Paging {
	p, _ := r.Query("page", `^\d+$`, false)
	ps, _ := r.Query("page_size", `^\d+$`, false)

	page, _ := p.Int()
	size, _ := ps.Int()

	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = 20
	} else {
		if size > 100 {
			size = 100
		}
	}
	limit := size
	offset := (page - 1) * limit
	return Paging{
		Page:     page,
		PageSize: size,
		Offset:   offset,
		Limit:    limit,
	}
}
