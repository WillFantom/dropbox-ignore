package files

type File struct {
	RelativePath string
	Directory bool
	Ignored bool
}

func NewFile(path string, isDir bool, defaultIgnored bool) (*File) {
	var file File
	file.RelativePath = path
	file.Directory = isDir
	file.Ignored = defaultIgnored
	return &file
}