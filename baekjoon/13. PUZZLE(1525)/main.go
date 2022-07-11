package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

var metrix [3][3]int

func main() {
	reader := bufio.NewReader(os.Stdin)

	for i := 0; i < 3; i++ {
		strTemp := strings.Split(readLine(reader), " ")
		metrix[i][0], _ = strconv.Atoi(strTemp[0])
		metrix[i][1], _ = strconv.Atoi(strTemp[1])
		metrix[i][2], _ = strconv.Atoi(strTemp[2])
	}
}

func readLine(reader *bufio.Reader) string {
	buff, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}
	return strings.TrimRight(string(buff), "\r\n")
}

type Que []interface{}

func (q Que) IsEmpty() bool {
	return len(*q) == 0
}

func (q Que) Enque(arg interface{}) {
	*q = append(*q, arg)
}

func (q Que) Deque() interface{} {
	if q.IsEmpty() {
		return nil
	}

	data := (*q)[0]
	(*q) = (*q)[1:]
	return data
}
