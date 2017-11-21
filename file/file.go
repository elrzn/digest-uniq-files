package file

import (
	"strings"
)

type File struct {
	Dir       string
	Extension string
	Name      string
	hash      *string
}

func New(fullname, dir string) File {
	name, ext := nameExt(fullname)
	return File{Name: name, Extension: ext, hash: nil, Dir: dir}
}

func Hash(f *File) string {
	if f.hash == nil {
		hash := "TODO"
		f.hash = &hash
	}
	return *f.hash
}

func nameExt(fullname string) (string, string) {

	const dot = "."

	if fullname == "" {
		return "", ""
	}

	xs := strings.Split(fullname, dot)
	size := len(xs)

	if size == 1 {
		return fullname, ""
	}

	return strings.Join(xs[0:size-1], dot), xs[size-1]
}
