package adto

type Paging struct {
	Page   int `name:"page"`
	Offset int `name:"offset"`
	Limit  int `name:"limit"`
}
