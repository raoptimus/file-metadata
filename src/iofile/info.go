package iofile

import "os"

func Info(name string) (*FileInfo, error) {
	mimeType, err := MimeType(name)
	if err != nil {
		return nil, err
	}

	hash, err := Hash(name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(name)
	if err != nil {
		return nil, err
	}

	return &FileInfo{
		MimeType: mimeType,
		HashMD5:  hash,
		Size:     fi.Size(),
	}, nil
}
