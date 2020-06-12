package main

import (
	"io/ioutil"
	"os"
	"path"
)

func getFile(pageName string, flags int) (file []byte, err error) {
	if f, err := os.OpenFile(path.Join(*pageDir, pageName), flags, 0755); err == nil {
		return ioutil.ReadAll(f)
	}
	return file, err
}

func saveFile(pageName string, contents []byte) error {
	return ioutil.WriteFile(path.Join(*pageDir, pageName), contents, 0700)
}
