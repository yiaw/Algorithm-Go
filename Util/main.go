package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	str := readLine(reader)

	for _, v := range str {
		fmt.Fprintln(writer, string(v))
	}

	fmt.Fprintf(writer, "%d", 1)
	writer.Flush()

}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}
	return strings.TrimRight(string(str), "\r\n")
}
