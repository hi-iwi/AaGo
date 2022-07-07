package atype

type Image string
type Video string
type Images struct{ NullStrings }
type Videos struct{ NullStrings }

func (im Image) Src(filler func(path string) ImgSrc) *ImgSrc {
	return ToImgSrcPtr(string(im), filler)
}

func (im Images) Srcs(filler func(path string) ImgSrc) []ImgSrc {
	return ToImgSrcs(im.Strings(), filler)
}
