package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	dpath := "./dir"
	filename := "something3.txt"

	_ = os.MkdirAll(dpath, 0777)

	fpath := filepath.Join(dpath, filename)

	file, _ := os.Create(fpath)

	file.Close()

	files, _ := ioutil.ReadDir("./dir")

	for _, f := range files {
		fmt.Println(f.Name(), f.ModTime())
	}

	// filpath walkFunc
	scan := func(path string, i os.FileInfo, _ error) error {
		fmt.Println(i.IsDir(), path)
		return nil
	}

	_ = filepath.Walk("./dir", scan)
}
