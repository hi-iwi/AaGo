package adto

// 存储在数据库里面，图片列表，为了节省空间，用数组来
// 数据库存储方式为 dtype.NullImgSrc，即  [path,size,width,height]
type ImgSrc struct {
	Processor int    `name:"-" json:"processor"` // 图片处理ID，如阿里云图片处理、网易云图片处理等
	Fill      string `name:"-" json:"fill"`      // e.g.  https://xxx/img.jpg?width=${WIDTH}&height=${HEIGHT}
	Fit       string `name:"-" json:"fit"`       // e.g. https://xxx/img.jpg?maxwidth=${MAXWIDTH}
	Path      string `name:"path" json:"path"`   // path 可能是 filename，也可能是 带文件夹的文件名
	Size      uint32 `name:"size" json:"size"`
	Width     uint16 `name:"width" json:"width"`
	Height    uint16 `name:"height" json:"height"`
}

type VideoSrc struct {
	Processor int    `name:"-" json:"processor"`
	Fit       string `name:"-" json:"fit"` // e.g.  https://xxx/video.avi?quality=${QUALITY}
	Path      string `name:"path" json:"path"`
	Size      uint32 `name:"size" json:"size"`
	Width     uint16 `name:"width" json:"width"`
	Height    uint16 `name:"height" json:"height"`
	Duration  uint32 `name:"duration" json:"duration"` // 时长，秒
}
