package file

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// File defines an struct for working around files.
type File struct {
	Path string
	Ext  string
	hash *string
}

func ofPath(path string) File {
	return File{Path: path, Ext: ext(path), hash: nil}
}

// Hash returns the MD5 hash of the contents of the file. The result is then
// memoized.
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

// Find recursively retrieves all files from the given directory, that match an
// optional list of extensions.
func Find(dir string, ext []string) []File {

	files := []File{}

	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			for _, r := range ext {
				if strings.HasSuffix(info.Name(), r) {
					files = append(files, ofPath(path))
				}
			}
		}
		return nil
	})

	return files
}

// Copy copies the file as the specified target.
func (f *File) Copy(target string) error {

	src, err := os.Open(f.Path)
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create(target)
	if err != nil {
		return err
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	if err = dst.Sync(); err != nil {
		return err
	}

	return nil
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
