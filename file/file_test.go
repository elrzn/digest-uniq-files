package file

import "testing"

// TODO Show actual values.
func TestNameExtension(t *testing.T) {

	e := ext("")
	if e != "" {
		t.Fatal("extension should be empty")
	}

	e = ext("foo")
	if e != "" {
		t.Fatal("extension should be empty")
	}

	e = ext("foo.bar")
	if e != "bar" {
		t.Fatal("bad ext")
	}

	e = ext("foo.bar.baz")
	if e != "baz" {
		t.Fatal("bad ext")
	}
}
