package file

import "testing"

// TODO Show actual values.
func TestNameExtension(t *testing.T) {

	name, ext := nameExt("")
	if name != "" {
		t.Fatal("name should be empty")
	}
	if ext != "" {
		t.Fatal("extension should be empty")
	}

	name, ext = nameExt("foo")
	if name != "foo" {
		t.Fatal("bad name")
	}
	if ext != "" {
		t.Fatal("extension should be empty")
	}

	name, ext = nameExt("foo.bar")
	if name != "foo" {
		t.Fatal("bad name")
	}
	if ext != "bar" {
		t.Fatal("bad ext")
	}

	name, ext = nameExt("foo.bar.baz")
	if name != "foo.bar" {
		t.Fatal("bad name")
	}
	if ext != "baz" {
		t.Fatal("bad ext")
	}
}
