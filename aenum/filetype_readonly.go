package aenum
import "strings"

const (
    UnknownType FileType = 0
    Jpeg        FileType = 1
    Png         FileType = 2
    Gif         FileType = 3
    Webp        FileType = 4
    Heic        FileType = 5
    Ico         FileType = 6
    Svg         FileType = 7
    Mp3         FileType = 1000
    Audio3gpp   FileType = 1001
    Audio3gpp2  FileType = 1002
    Aiff        FileType = 1003
    AudioWebm   FileType = 1004
    AudioWav    FileType = 1005
    Avi         FileType = 2000
    Mov         FileType = 2001
    Mpeg        FileType = 2002
    Mp4         FileType = 2003
    Video3gp    FileType = 2004
    Video3gp2   FileType = 2005
    Webm        FileType = 2006
    Wav         FileType = 2007
    Pdf         FileType = 3001
    Txt         FileType = 3002
    Md          FileType = 3003
    Doc         FileType = 3004
    Docx        FileType = 3005
    Xls         FileType = 3006
    Xlsx        FileType = 3007
    Ppt         FileType = 3008
    Pptx        FileType = 3009
    Zip         FileType = 7000
    Rar         FileType = 7001
    Bzip        FileType = 7002
    Bzip2       FileType = 7003
    Gzip        FileType = 7004
    Json        FileType = 10000
)
var ImageTypes = map[FileType][]string{
    Gif         : {".gif", "image/gif"},
    Heic        : {".heic", "image/heic", ".heif", ".avci", "image/heif"},
    Ico         : {".ico", "image/vnd.microsoft.icon", "image/x-icon"},
    Jpeg        : {".jpg", "image/jpeg", ".jpeg"},
    Png         : {".png", "image/png"},
    Svg         : {".svg", "image/svg+xml"},
    Webp        : {".webp", "image/webp"},
}
var AudioTypes = map[FileType][]string{
    Aiff        : {".aiff", "audio/aiff", ".aif", ".aifc", "audio/x-aiff"},
    Audio3gpp   : {".3gp", "audio/3gpp"},
    Audio3gpp2  : {".3g2", "audio/3gpp2"},
    AudioWav    : {".webm", "audio/webm"},
    AudioWebm   : {".wav", "audio/wav"},
    Mp3         : {".mp3", "audio/mpeg", "audio/mp3"},
}
var VideoTypes = map[FileType][]string{
    Avi         : {".avi", "video/x-msvideo"},
    Mov         : {".mov", "video/quicktime"},
    Mp4         : {".mp4", "video/mp4"},
    Mpeg        : {".mpeg", "video/mpeg"},
    Video3gp    : {".3gp", "video/3gpp"},
    Video3gp2   : {".3g2", "video/3gpp2"},
    Wav         : {".wav", "video/x-wav"},
    Webm        : {".webm", "video/webm"},
}
var DocumentTypes = map[FileType][]string{
    Doc         : {".doc", "application/msword"},
    Docx        : {".docx", "application/vnd.openxmlformats-officedocument.wordprocessingml.document"},
    Md          : {".md", "text/markdown"},
    Pdf         : {".pdf", "application/pdf"},
    Ppt         : {".ppt", "application/vnd.ms-powerpoint"},
    Pptx        : {".pptx", "application/vnd.openxmlformats-officedocument.presentationml.presentation"},
    Txt         : {".txt", "text/plain"},
    Xls         : {".xls", "application/vnd.ms-excel"},
    Xlsx        : {".xlsx", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"},
}
var CompressedTypes = map[FileType][]string{
    Bzip        : {".bz", "application/x-bzip"},
    Bzip2       : {".bz2", "application/x-bzip2"},
    Gzip        : {".gz", "application/gzip", "application/x-gzip"},
    Rar         : {".rar", "application/vnd.rar", "application/x-rar-compressed"},
    Zip         : {".zip", "application/zip", "application/x-zip-compressed", "multipart/x-zip"},
}
var DataTypes = map[FileType][]string{
    Json        : {".json", "application/json"},
}
func NewAudioType(mime string) (FileType, bool) {return ParseFileType(mime, AudioTypes)}
func NewVideoType(mime string) (FileType, bool) {return ParseFileType(mime, VideoTypes)}
func NewDocumentType(mime string) (FileType, bool) {return ParseFileType(mime, DocumentTypes)}
func NewCompressedType(mime string) (FileType, bool) {return ParseFileType(mime, CompressedTypes)}
func NewDataType(mime string) (FileType, bool) {return ParseFileType(mime, DataTypes)}
func NewImageType(mime string) (FileType, bool) {return ParseFileType(mime, ImageTypes)}
func (t FileType) ContentType() string {
    if d, ok := VideoTypes[t]; ok {return d[1]}
    if d, ok := DocumentTypes[t]; ok {return d[1]}
    if d, ok := CompressedTypes[t]; ok {return d[1]}
    if d, ok := DataTypes[t]; ok {return d[1]}
    if d, ok := ImageTypes[t]; ok {return d[1]}
    if d, ok := AudioTypes[t]; ok {return d[1]}
    return ""
}
func (t FileType) Ext() string {
    if d, ok := AudioTypes[t]; ok {return d[0]}
    if d, ok := VideoTypes[t]; ok {return d[0]}
    if d, ok := DocumentTypes[t]; ok {return d[0]}
    if d, ok := CompressedTypes[t]; ok {return d[0]}
    if d, ok := DataTypes[t]; ok {return d[0]}
    if d, ok := ImageTypes[t]; ok {return d[0]}
    return ""
}
func (t FileType) Name() string {return strings.TrimPrefix(t.Ext(), ".")}
