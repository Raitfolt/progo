package main

import (
	"os"
	"path/filepath"
)

func main() {
	path, err := os.Getwd()
	if err == nil {
		matches, err := filepath.Glob(filepath.Join(path, "*.json"))
		if err == nil {
			for _, m := range matches {
				Printfln("Match: %v", m)
			}
		}
	}
	if err != nil {
		Printfln("Error %v", err.Error())
	}
}

//Raitfolt
//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
