package iofile

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
)

func Hash(name string) (string, error) {
	f, err := os.Open(name)
	if err != nil {
		return "", err
	}
	defer f.Close()

	return HashFile(f)
}

func HashFile(file *os.File) (string, error) {
	defer file.Seek(0, io.SeekStart)
	md5Hash := md5.New()
	if _, err := io.Copy(md5Hash, file); err != nil {
		return "", err
	}

	return hex.EncodeToString(md5Hash.Sum(nil)), nil
}
