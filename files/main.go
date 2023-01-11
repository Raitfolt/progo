package main

import (
	"os"
	"path/filepath"
)

func callback(path string, dir os.DirEntry, dirErr error) (err error) {
	info, _ := dir.Info()
	Printfln("Path %v, Size: %v", path, info.Size())
	return
}

func main() {
	path, err := os.Getwd()
	if err == nil {
		_ = filepath.WalkDir(path, callback)
	} else {
		Printfln("Error %v", err.Error())
	}
}

//Raitfolt
//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
