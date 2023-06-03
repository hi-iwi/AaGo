package atype

type Paging struct {
	Page   uint `json:"page"`
	Offset uint `json:"offset"`
	Limit  uint `json:"limit"`
	Prev   uint `json:"prev"`
	Next   uint `json:"next"`
}
