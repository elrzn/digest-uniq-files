package file

import (
	"os"
	"path/filepath"
	"strings"
)

type File struct {
	Path string
	hash *string
}

func New(path string) File {
	return File{Path: path, hash: nil}
}

func Hash(f *File) string {
	if f.hash == nil {
		hash := "TODO"
		f.hash = &hash
	}
	return *f.hash
}

func Find(dir string, ext []string) []File {

	files := []File{}

	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			for _, r := range ext {
				if strings.HasSuffix(info.Name(), r) {
					files = append(files, New(path))
				}
			}
		}
		return nil
	})

	return files
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
