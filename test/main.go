package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'solution' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. INTEGER_ARRAY infectedHouses
 */

var count int32
var chk []bool

// 이동은 x +1, x -1 두가지 case가 존재

var cnt int32

type Que []int32

func (q *Que) IsEmpty() bool {
	return len(*q) == 0
}

func (q *Que) Enque(data int32) {
	(*q) = append(*q, data)
}

func (q *Que) Deque() (int32, error) {
	if (*q).IsEmpty() {
		return 0, errors.New("empty")
	}

	data := (*q)[0]
	*q = (*q)[1:]

	return data, nil
}

func solution(n int32, infectedHouses []int32) int32 {
	// Write your code here
	chk = make([]bool, n+1)
	var q Que
	var nv int32
	for _, i := range infectedHouses {
		chk[i] = true
		q.Enque(i)
	}

	for !q.IsEmpty() {
		v, _ := q.Deque()
		for i := 0; i < 2; i++ {
			switch i {
			case 0:
				nv = v + 1
			case 1:
				nv = v - 1
			}
		}
		if (1 <= nv && nv <= (n)+1) && !chk[nv] {
			chk[nv] = true
			q.Enque(nv)
		}

	}

	return cnt
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	writer := bufio.NewWriterSize(os.Stdout, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	infectedHousesCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var infectedHouses []int32

	for i := 0; i < int(infectedHousesCount); i++ {
		infectedHousesItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		infectedHousesItem := int32(infectedHousesItemTemp)
		infectedHouses = append(infectedHouses, infectedHousesItem)
	}

	result := solution(n, infectedHouses)

	fmt.Fprintf(writer, "%d\n", result)

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
