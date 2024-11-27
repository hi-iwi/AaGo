package com

import (
	"github.com/hi-iwi/AaGo/ae"
	"github.com/hi-iwi/AaGo/atype"
	"strconv"
	"strings"
)

// {id:uint64}  or {sid:string}
func (r *Req) QueryId(p string, params ...any) (sid string, id uint64, e *ae.Error) {
	var x *ReqProp
	if x, e = r.Query(p, params...); e != nil {
		return
	}
	sid = x.String()
	if sid == "" || sid == "0" {
		return
	}
	for _, s := range sid {
		if s < '0' || (s > '9' && s < 'A') || (s > 'Z' && s < '_') || (s > '_' && s < 'a') || s > 'z' {
			e = ae.BadParam(p)
			return
		}
		if s > '9' {
			return
		}
	}
	id, _ = strconv.ParseUint(sid, 10, 64)
	return
}

// 不可再指定offset/limit了，单一原则，通过page分页
// @param firstPageLimit 首页行数

func (r *Req) QueryPaging(perPageLimit, firstPageEnd uint) atype.Paging {
	page, _ := r.QueryUint(ParamPage, false)
	pageEnd, _ := r.QueryUint(ParamPageEnd, false)
	return atype.NewPaging(perPageLimit, page, pageEnd, firstPageEnd)
}

func (r *Req) QueryPage() atype.Paging {
	page, _ := r.QueryUint(ParamPage, false)
	pageEnd, _ := r.QueryUint(ParamPageEnd, false)
	return atype.NewPage(page, pageEnd)
}

func (r *Req) BodyImage(p string, required ...bool) (atype.Image, *ae.Error) {
	x, e := r.BodyString(p, `^([\w-\/\.]+)$`, len(required) == 0 || required[0])
	if e != nil {
		return "", e
	}
	return atype.Image(x), nil
}
func (r *Req) BodyAudio(p string, required ...bool) (atype.Audio, *ae.Error) {
	x, e := r.BodyString(p, `^([\w-\/\.]+)$`, len(required) == 0 || required[0])
	if e != nil {
		return "", e
	}

	return atype.Audio(x), e
}
func (r *Req) BodyVideo(p string, required ...bool) (atype.Video, *ae.Error) {
	x, e := r.BodyString(p, `^([\w-\/\.]+)$`, len(required) == 0 || required[0])
	if e != nil {
		return "", e
	}

	return atype.Video(x), e
}
func (r *Req) BodyImages(p string, required ...bool) ([]atype.Image, *ae.Error) {
	xx, e := r.BodyStrings(p, len(required) == 0 || required[0], false)
	if e != nil || len(xx) == 0 {
		return nil, e
	}
	imgs := make([]atype.Image, len(xx))
	for i, x := range xx {
		if x == "" || (strings.LastIndexByte(x, '.') < 0 || strings.IndexByte(x, ' ') > -1 || strings.IndexByte(x, '?') > -1 || strings.IndexByte(x, '=') > -1) {
			return nil, ae.BadParam(p)
		}
		imgs[i] = atype.Image(x)
	}
	return imgs, e
}
func (r *Req) BodyAudios(p string, required ...bool) ([]atype.Audio, *ae.Error) {
	xx, e := r.BodyStrings(p, len(required) == 0 || required[0], false)
	if e != nil || len(xx) == 0 {
		return nil, e
	}
	audios := make([]atype.Audio, len(xx))
	for i, x := range xx {
		if x == "" || (strings.LastIndexByte(x, '.') < 0 || strings.IndexByte(x, ' ') > -1 || strings.IndexByte(x, '?') > -1 || strings.IndexByte(x, '=') > -1) {
			return nil, ae.BadParam(p)
		}
		audios[i] = atype.Audio(x)
	}
	return audios, e
}
func (r *Req) BodyVideos(p string, required ...bool) ([]atype.Video, *ae.Error) {
	xx, e := r.BodyStrings(p, len(required) == 0 || required[0], false)
	if e != nil || len(xx) == 0 {
		return nil, e
	}
	videos := make([]atype.Video, len(xx))
	for i, x := range xx {
		if x == "" || (strings.LastIndexByte(x, '.') < 0 || strings.IndexByte(x, ' ') > -1 || strings.IndexByte(x, '?') > -1 || strings.IndexByte(x, '=') > -1) {
			return nil, ae.BadParam(p)
		}
		videos[i] = atype.Video(x)
	}
	return videos, e
}

func (r *Req) BodyCoordinate(p string, required ...bool) (*atype.Coordinate, *ae.Error) {
	x, e := r.BodyFloat64Map(p, required...)
	if e != nil || x == nil {
		return nil, e
	}
	lat, ok := x["lat"]
	if !ok {
		return nil, ae.BadParam(p)
	}
	lng, ok := x["lng"]
	if !ok {
		return nil, ae.BadParam(p)
	}
	height, _ := x["height"]
	coord := atype.Coordinate{
		Latitude:  lat,
		Longitude: lng,
		Height:    height,
	}
	return &coord, nil
}

func (r *Req) BodyLocation(p string, required ...bool) (atype.Location, *ae.Error) {
	x, e := r.BodyInterfaceMap(p, required...)
	if e != nil || x == nil {
		return atype.Location{}, e
	}
	var loc atype.Location
	lat, ok := x["lat"]
	if !ok {
		e = ae.BadParam(p)
		return atype.Location{}, e
	}

	if loc.Latitude, ok = lat.(float64); !ok {
		e = ae.BadParam(p)
		return atype.Location{}, e
	}
	lng, ok := x["lng"]
	if !ok {
		e = ae.BadParam(p)
		return atype.Location{}, e
	}

	if loc.Longitude, ok = lng.(float64); !ok {
		e = ae.BadParam(p)
		return atype.Location{}, e
	}
	if ht, ok := x["height"]; ok {
		loc.Height, _ = ht.(float64)
	}
	loc.Valid = true
	loc.Name = atype.String(x["name"])
	loc.Address = atype.String(x["address"])
	return loc, nil
}
