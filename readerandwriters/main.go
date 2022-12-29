package main

import (
	"io"
	"strings"
)

func processData(reader io.Reader, writer io.Writer) {
	b := make([]byte, 2)
	for {
		count, err := reader.Read(b)
		if count > 0 {
			writer.Write(b[0:count])
			Printfln("Read %v bytes: %v", count, string(b[0:count]))
		}
		if err == io.EOF {
			break
		}
	}
}

func processData2(reader io.Reader, writer io.Writer) {
	count, err := io.Copy(writer, reader)
	if err == nil {
		Printfln("Read %v bytes", count)
	} else {
		Printfln("Error: %v", err.Error())
	}
}

func main() {
	r1 := strings.NewReader("Kayak")
	r2 := strings.NewReader("Lifejacket")
	r3 := strings.NewReader("Canoe")

	concatReader := io.MultiReader(r1, r2, r3)

	limited := io.LimitReader(concatReader, 5)
	ConsumeData(limited)
}
