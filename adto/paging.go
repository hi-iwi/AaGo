package adto

import "github.com/hi-iwi/AaGo/com"

type Paging struct {
	Page   int `alias:"page"`
	Offset int `alias:"offset"`
	Limit  int `alias:"limit"`
}

func MakePaging(r *com.Req, args ...int) Paging {
	p, _ := r.Query("page", `^\d+$`, false)
	ofs, _ := r.Query("offset", `^\d+$`, false)
	lmt, _ := r.Query("limit", `^\d+$`, false)

	page, _ := p.Int()
	offset, _ := ofs.Int()
	limit, _ := lmt.Int()

	if limit < 1 {
		if len(args) > 0 {
			limit = args[0]
		} else {
			limit = 20
		}
	} else if limit > 100 {
		limit = 100
	}

	if offset > 0 {
		page = (offset / limit) + 1
	} else {
		if page < 1 {
			page = 1
		}
	}
	// change ?limit=3&offset=10 to ?limit=0&offset=10
	offset = (page - 1) * limit

	return Paging{
		Page:   page,
		Offset: offset,
		Limit:  limit,
	}
}
