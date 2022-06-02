//https://www.acmicpc.net/problem/1926

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

var cnt int32
var maxPainting int32
var n, m int32
var metric [][]int32
var chk [][]bool

var dx = []int32{1, 0, -1, 0}
var dy = []int32{0, 1, 0, -1}

type Pos struct {
	X int32
	Y int32
}

func bfs(x, y int32) int32 {
	var q Que
	var innerMax int32

	q.Enque(Pos{X: x, Y: y})
	innerMax++
	for !q.IsEmpty() {
		pos, _ := q.Deque()
		for i := 0; i < 4; i++ {
			nx := pos.(Pos).X + dx[i]
			ny := pos.(Pos).Y + dy[i]
			if (0 <= nx && nx < m) && (0 <= ny && ny < n) {
				if metric[ny][nx] == 1 && !chk[ny][nx] {
					chk[ny][nx] = true
					innerMax++
					q.Enque(Pos{X: nx, Y: ny})
				}
			}
		}
	}

	return innerMax

}
func main() {

	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	inputLine := strings.Split(strings.TrimSpace(readLine(reader)), " ")
	nTemp, _ := strconv.ParseInt(inputLine[0], 10, 64)
	mTemp, _ := strconv.ParseInt(inputLine[1], 10, 64)

	n = int32(nTemp)
	m = int32(mTemp)

	metric = make([][]int32, n)
	chk = make([][]bool, n)

	for i := 0; i < int(n); i++ {
		metric[i] = make([]int32, m)
		chk[i] = make([]bool, m)
	}

	for i := 0; i < int(n); i++ {
		str := strings.Split(readLine(reader), " ")
		for j := 0; j < int(m); j++ {
			mTemp, _ := strconv.ParseInt(str[j], 10, 64)
			metric[i][j] = int32(mTemp)
		}
	}

	for i := 0; i < int(n); i++ {
		for j := 0; j < int(m); j++ {
			if metric[i][j] == 1 && !chk[i][j] {

				cnt++
				chk[i][j] = true
				temp := bfs(int32(j), int32(i))
				if temp > maxPainting {
					maxPainting = temp
				}
			}
		}
	}

	fmt.Fprintln(writer, cnt)
	fmt.Fprintln(writer, maxPainting)
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
		return 0, errors.New("empty")
	}

	data := (*q)[0]
	*q = (*q)[1:]

	return data, nil
}
