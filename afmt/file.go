package afmt

func DirPath(dir string) string {
	l := len(dir)
	if dir[0] != '/' {
		dir = "/" + dir
	}
	if dir[l - 1] == '/' {
		dir = dir[0: l - 2]
	}
	return dir
}

func FilePath(filepath string) string {
	if filepath[0] == '/' {
		filepath = filepath[1:]
	}
	return filepath
}

