package main

import (
	"os"
	"path/filepath"
)

func main() {
	dpath := "./dir"
	filename := "something.txt"

	_ = os.MkdirAll(dpath, 0777)

	fpath := filepath.Join(dpath, filename)

	file, _ := os.Create(fpath)

	file.Close()

	_ = os.Remove(fpath)

	_ = os.Remove(dpath)
}
