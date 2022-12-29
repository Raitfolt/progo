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

func writeFormatted(writer io.Writer, template string, vals ...interface{}) {
	fmt.Fprintf(writer, template, vals...)
}

func writeReplaced(writer io.Writer, str string, subs ...string) {
	replacer := strings.NewReplacer(subs...)
	replacer.WriteString(writer, str)
}

func main() {
	reader := strings.NewReader("Kayak Watersports $279.00")

	for {
		var str string
		_, err := scanSingle(reader, &str)
		if err != nil {
			if err != io.EOF {
				Printfln("Error: %v", err.Error())
			}
			break
		}
		Printfln("Value: %v", str)
	}

	var writer strings.Builder
	template := "Name: %s, Category: %s, Price: $%.2f\n"
	writeFormatted(&writer, template, "Kayak", "Watersports", float64(279))

	fmt.Println(writer.String())

	text := "It was a boat. A small boat."
	subs := []string{"boat", "kayak", "small", "huge"}
	writeReplaced(&writer, text, subs...)
	fmt.Println(writer.String())
}
