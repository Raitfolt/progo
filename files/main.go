package main

import (
	"os"
)

func main() {
	path, err := os.Getwd()
	if err == nil {
		dirEntries, err := os.ReadDir(path)
		if err == nil {
			for _, dentry := range dirEntries {
				Printfln("Entry name: %v, IsDir: %v", dentry.Name(), dentry.IsDir())
			}
		}
	}
	if err != nil {
		Printfln("Error %v", err.Error())
	}
}

//Raitfolt
//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
