package main

import (
	"bufio"
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
	text := "It was a boat. A small boat."

	var builder strings.Builder
	var writer = bufio.NewWriterSize(NewCustomWriter(&builder), 20)
	for i := 0; true; {
		end := i + 5
		if end >= len(text) {
			writer.Write([]byte(text[i:]))
			writer.Flush()
			break
		}
		writer.Write([]byte(text[i:end]))
		i = end
	}

	Printfln("Written data: %v", builder.String())
}
