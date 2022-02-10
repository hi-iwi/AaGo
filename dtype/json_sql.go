package dtype

import (
	"database/sql"
	"encoding/json"
	"github.com/hi-iwi/AaGo/adto"
)

type NullJson sql.NullString
type NullUint8s sql.NullString        // uint8 json array
type NullUint16s sql.NullString       // uint16 json array
type NullUint32s sql.NullString       // uint32 json array
type NullInts sql.NullString          // int json array
type NullUints sql.NullString         // uint json array
type NullUint64s sql.NullString       // uint64 json array
type NullStrings sql.NullString       // string json array
type NullStringMap sql.NullString     // map[string]string
type Null2dStringMap sql.NullString   // map[string]map[string]string
type NullStringMaps sql.NullString    // []map[string]string
type NullStringMapsMap sql.NullString // map[string][]map[string]string
type NullImgSrc sql.NullString        // adto.ImgSrc
type NullImgSrcs sql.NullString       // []adto.ImgSrc
type NullVideoSrc sql.NullString      // adto.VideoSrc
type NullVideosSrcs sql.NullString    // []adto.VideoSrc

func (t NullUint8s) Uint8s() []uint8 {
	if !t.Valid || t.String == "" {
		return nil
	}
	var v []uint8
	json.Unmarshal([]byte(t.String), &v)
	return v
}

func (t NullUint16s) Uint16s() []uint16 {
	if !t.Valid || t.String == "" {
		return nil
	}
	var v []uint16
	json.Unmarshal([]byte(t.String), &v)
	return v
}
func (t NullUint32s) Uint32s() []uint32 {
	if !t.Valid || t.String == "" {
		return nil
	}
	var v []uint32
	json.Unmarshal([]byte(t.String), &v)
	return v
}
func (t NullInts) Ints() []int {
	if !t.Valid || t.String == "" {
		return nil
	}
	var v []int
	json.Unmarshal([]byte(t.String), &v)
	return v
}
func (t NullUints) Uints() []uint {
	if !t.Valid || t.String == "" {
		return nil
	}
	var v []uint
	json.Unmarshal([]byte(t.String), &v)
	return v
}
func (t NullUint64s) Uint64s() []uint64 {
	if !t.Valid || t.String == "" {
		return nil
	}
	var v []uint64
	json.Unmarshal([]byte(t.String), &v)
	return v
}
func (t NullStrings) Strings() []string {
	if !t.Valid || t.String == "" {
		return nil
	}
	var v []string
	json.Unmarshal([]byte(t.String), &v)
	return v
}

func (t NullStringMap) StringMap() map[string]string {
	if !t.Valid || t.String == "" {
		return nil
	}
	var v map[string]string
	json.Unmarshal([]byte(t.String), &v)
	return v
}

func (t Null2dStringMap) TStringMap() map[string]map[string]string {
	if !t.Valid || t.String == "" {
		return nil
	}
	var v map[string]map[string]string
	json.Unmarshal([]byte(t.String), &v)
	return v
}
func (t NullStringMaps) StringMaps() []map[string]string {
	if !t.Valid || t.String == "" {
		return nil
	}
	var v []map[string]string
	json.Unmarshal([]byte(t.String), &v)
	return v
}
func (t NullStringMapsMap) StringMapsMap() map[string][]map[string]string {
	if !t.Valid || t.String == "" {
		return nil
	}
	var v map[string][]map[string]string
	json.Unmarshal([]byte(t.String), &v)
	return v
}

func (t NullImgSrc) ImgSrc() *adto.ImgSrc {
	if !t.Valid || t.String == "" {
		return nil
	}

	// 为了节省存储空间，这里使用 [path,size,width,height] 数组方式存储
	//  If you sent the JSON value through browser then any number you sent that will be the type float64
	var m [4]interface{}
	err := json.Unmarshal([]byte(t.String), &m)
	if err != nil {
		return nil
	}
	return &adto.ImgSrc{
		Path:   New(m[0]).String(),
		Size:   New(m[1]).DefaultUint32(0),
		Width:  New(m[2]).DefaultUint16(0),
		Height: New(m[3]).DefaultUint16(0),
	}
}

func (t NullImgSrcs) ImgSrcs() []adto.ImgSrc {
	if !t.Valid || t.String == "" {
		return nil
	}

	// 为了节省存储空间，这里使用 [[path,size,width,height],[path,size,width,height]...] 数组方式存储
	//  If you sent the JSON value through browser then any number you sent that will be the type float64
	var ms [][4]interface{}
	err := json.Unmarshal([]byte(t.String), &ms)
	if err != nil {
		return nil
	}
	v := make([]adto.ImgSrc, len(ms))
	for i, m := range ms {
		v[i] = adto.ImgSrc{
			Path:   New(m[0]).String(),
			Size:   New(m[1]).DefaultUint32(0),
			Width:  New(m[2]).DefaultUint16(0),
			Height: New(m[3]).DefaultUint16(0),
		}
	}
	return v
}
