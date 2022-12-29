package main

import (
	"fmt"
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

func scanFromReader(reader io.Reader, template string, vals ...interface{}) (int, error) {
	return fmt.Fscanf(reader, template, vals...)
}

func scanSingle(reader io.Reader, val interface{}) (int, error) {
	return fmt.Fscan(reader, val)
}

func main() {
	reader := strings.NewReader("Kayak Watersports $279.00")

	var name, category string
	var price float64
	scanTemplate := "%s %s $%f"

	_, err := scanFromReader(reader, scanTemplate, &name, &category, &price)
	if err != nil {
		Printfln("Error: %v", err.Error())
	} else {
		Printfln("Name: %v", name)
		Printfln("Category: %v", category)
		Printfln("Price: %.2f", price)
	}
}
