package iofile

import (
	"os"
	"testing"
)

func TestFileExists(t *testing.T) {
	if Exists("testdata/test.txt") != true {
		t.Fatal("file `testdata/test.txt` exists but function returns `not exists`")
	}
	if Exists("testdata/somefile.txt") != false {
		t.Fatal("file `testdata/somefile.txt` does not exists but function returns `exists`")
	}
}

func TestMkDir(t *testing.T) {
	os.Remove("testdata/test")

	err := Mkdir("testdata/test", os.ModePerm)
	if err != nil {
		t.Fatal(err)
	}

	stat, err := os.Stat("testdata/test")
	if err != nil {
		t.Fatal(err)
	}

	if !stat.IsDir() {
		t.Fatalf("testdata/test must be is directory")
	}
}
