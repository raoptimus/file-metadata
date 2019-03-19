package iofile

import "testing"

func TestFileInfo(t *testing.T) {
	info, err := Info("testdata/test.txt")
	if err != nil {
		t.Fatal(err)
	}

	expectedInfo := FileInfo{
		MimeType: "text",
		HashMD5:  "ee7e7e64fd93df342b00b5a516b8b900",
		Size:     112,
	}

	if expectedInfo != *info {
		t.Fatalf("expected info `%v`, actual `%v`", expectedInfo, *info)
	}
}
