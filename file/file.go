package file

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type File struct {
	Path string
	Ext  string
	hash *string
}

func New(path string) File {
	return File{Path: path, Ext: ext(path), hash: nil}
}

func (f *File) Hash() string {
	if f.hash == nil {
		md5, _ := f.makeMD5()
		f.hash = &md5
	}
	return *f.hash
}

func (f File) makeMD5() (string, error) {

	file, err := os.Open(f.Path)
	if err != nil {
		return "", err
	}

	defer file.Close()

	md5 := md5.New()
	if _, err := io.Copy(md5, file); err != nil {
		return "", err
	}

	return hex.EncodeToString(md5.Sum(nil)[:16]), nil
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

func ext(path string) string {

	const dot = "."

	if path == "" {
		return ""
	}

	xs := strings.Split(path, dot)
	size := len(xs)

	if size == 1 {
		return ""
	}

	return xs[size-1]
}
