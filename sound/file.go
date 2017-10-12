package sound

type File struct {
	FileName string
	Content  [][]byte
}

func NewFile(filename string, content [][]byte) *File {
	return &File{
		filename,
		content,
	}
}
