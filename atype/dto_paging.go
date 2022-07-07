package atype

type Paging struct {
	Page   int `json:"page"`
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}
