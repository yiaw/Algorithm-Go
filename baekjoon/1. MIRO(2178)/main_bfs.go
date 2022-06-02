package main

//https://www.acmicpc.net/problem/2178
import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var n, m int32
var metric [][]int32
var cnt [][]int32
var chk [][]bool

type Location struct {
	X int32
	Y int32
}

func bfs(y, x int32) {
	var q Que
	var dx = []int32{1, 0, -1, 0}
	var dy = []int32{0, -1, 0, 1}
	chk[y][x] = true
	cnt[y][x] = 1

	// Que Append
	q.Enque(Location{X: x, Y: y})
	for !q.IsEmpty() {
		l, _ := q.Deque()
		for i := 0; i < 4; i++ {
			ny := l.(Location).Y + dy[i]
			nx := l.(Location).X + dx[i]

			if (0 <= nx && nx < m) && (0 <= ny && ny < n) {
				if metric[ny][nx] == 1 && !chk[ny][nx] {
					cnt[ny][nx] = cnt[l.(Location).Y][l.(Location).X] + 1

					chk[ny][nx] = true
					q.Enque(Location{X: nx, Y: ny})
					if nx == m-1 && ny == n-1 {
						break
					}
				}
			}
		}
	}

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	writer := bufio.NewWriterSize(os.Stdout, 16*1024*1024)
	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	n = int32(nTemp)

	mTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	m = int32(mTemp)

	metric = make([][]int32, n)
	cnt = make([][]int32, n)
	chk = make([][]bool, n)
	for i := 0; i < int(n); i++ {
		metric[i] = make([]int32, m)
		cnt[i] = make([]int32, m)
		chk[i] = make([]bool, m)
	}

	for i := 0; i < int(n); i++ {
		str := readLine(reader)
		for j := 0; j < int(m); j++ {
			mTemp, _ = strconv.ParseInt(string(str[j]), 10, 64)
			metric[i][j] = int32(mTemp)
		}

	}
	bfs(0, 0)
	fmt.Fprintf(writer, "%d", cnt[n-1][m-1])
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
