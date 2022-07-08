package main

//https://www.acmicpc.net/problem/7576
import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// 입력값
var h, w int32

// 행렬 그래프
var metrix [201][201]int32

// 좌표
type Location struct {
	X int32 // x 좌표
	Y int32 // y 좌표
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	writer := bufio.NewWriterSize(os.Stdout, 16*1024*1024)

	whString := strings.Split(strings.TrimSpace(readLine(reader)), " ")
	wTemp, err := strconv.ParseInt(whString[0], 10, 64)
	checkError(err)
	w = int32(wTemp)
	hTemp, err := strconv.ParseInt(whString[1], 10, 64)
	checkError(err)
	h = int32(hTemp)

	for i := 0; i < int(h); i++ {
		str := strings.Split(strings.TrimSpace(readLine(reader)), " ")
		for j := 0; j < int(w); j++ {
			mTemp, _ := strconv.ParseInt(string(str[j]), 10, 64)
			metrix[i][j] = int32(mTemp)
		}

	}

	fmt.Fprintf(writer, "%d\n")

	writer.Flush()

}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

type Que []interface{}

func (q *Que) IsEmpty() bool {
	return len(*q) == 0
}

func (q *Que) Enque(data interface{}) {
	(*q) = append(*q, data)
}

func (q *Que) Deque() (interface{}, error) {
	if (*q).IsEmpty() {
		return 0, errors.New("empty")
	}

	data := (*q)[0]
	*q = (*q)[1:]

	return data, nil
}
