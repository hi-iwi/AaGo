package adto

type Bimg struct {
	//URL    string `json:"url"`
	Path   string `json:"path"`
	Width  uint16 `json:"width"`
	Height uint16 `json:"height"`
	Size   uint32 `json:"size"`
}
type Bvideo struct {
	//URL  string `json:"url"`
	Path   string `json:"path"`
	Width  uint16 `json:"width"`
	Height uint16 `json:"height"`
	Size   uint32 `json:"size"`
}
