//https://www.acmicpc.net/problem/2331
package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

// A = 9999
// P = 5
const (
	MAXCOUNT = 300000
)

var cnt int32
var visited [MAXCOUNT]int32

func dfs(A, P int32) {
	var next int32
	// 방문 증가
	visited[A]++
	// 전체적으로 반복되는 부분 확인하여 3일 경우 리턴 해버림
	if visited[A] > 2 {
		return
	}

	for A > 0 {
		digit := A % 10 //
		num := math.Pow(float64(digit), float64(P))
		next += int32(num)
		A = A / 10
	}

	dfs(next, P)
	// 각 자리수에 P 제곱 값
	// 해당 값에근접했던 변수 확인

}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	input := strings.Split(readLine(reader), " ")

	aTemp, _ := strconv.ParseInt(input[0], 10, 64)
	a := int32(aTemp)
	pTemp, _ := strconv.ParseInt(input[1], 10, 64)
	p := int32(pTemp)

	dfs(a, p)

	for i := 0; i < MAXCOUNT; i++ {
		if visited[i] == 1 {
			cnt++
		}
	}
	fmt.Fprintln(writer, cnt)
	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}
	return strings.TrimRight(string(str), "\r\n")
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
		return nil, errors.New("empty")
	}

	data := (*q)[0]
	*q = (*q)[1:]

	return data, nil
}
