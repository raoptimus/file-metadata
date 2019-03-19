package iofile

import (
	"fmt"
	"os/exec"
	"strings"
)

func MimeType(file string) (string, error) {
	out, err := exec.Command("file", "--mime-type", file).CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("identify: err: %s; detail: %s", err, string(out))
	}

	mime := strings.Split(strings.Split(string(out), ": ")[1], "/")[0]

	return mime, nil
}
