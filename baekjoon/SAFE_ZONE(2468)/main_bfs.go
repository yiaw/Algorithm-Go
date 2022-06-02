//https://www.acmicpc.net/problem/2468
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

type Pos struct {
	X int32
	Y int32
}

var n, max int32
var metrix [][]int32
var chk [][]bool
var dy = []int32{-1, 0, 1, 0}
var dx = []int32{0, 1, 0, -1}
var cnt int32

func bfs(y, x, h int32) {
	var q Que

	q.Enque(Pos{X: x, Y: y})

	for !q.IsEmpty() {
		p, _ := q.Deque()
		for i := 0; i < 4; i++ {
			ny := p.(Pos).Y + int32(dy[i])
			nx := p.(Pos).X + int32(dx[i])

			if (0 <= ny && ny < n) && (0 <= nx && nx < n) {
				if metrix[ny][nx] > h && !chk[ny][nx] {
					chk[ny][nx] = true
					q.Enque(Pos{X: nx, Y: ny})
				}
			}
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	input := readLine(reader)

	nTemp, _ := strconv.ParseInt(input, 10, 64)

	n = int32(nTemp)

	metrix = make([][]int32, n)
	chk = make([][]bool, n)
	for i := 0; i < int(n); i++ {
		metrix[i] = make([]int32, n)
		chk[i] = make([]bool, n)
	}

	// input value
	for i := 0; i < int(n); i++ {
		input := strings.Split(readLine(reader), " ")
		for j := 0; j < int(n); j++ {
			mTemp, _ := strconv.ParseInt(input[j], 10, 64)
			metrix[i][j] = int32(mTemp)
		}
	}

	for h := 0; h <= 100; h++ {
		cnt = 0
		reset()
		for i := 0; i < int(n); i++ {
			for j := 0; j < int(n); j++ {
				if metrix[j][i] > int32(h) && !chk[j][i] {
					cnt++
					chk[j][i] = true
					bfs(int32(j), int32(i), int32(h))
				}

			}
		}

		if max < cnt {
			max = cnt
		}

	}

	fmt.Fprintln(writer, max)
	writer.Flush()
}

func reset() {
	for i := 0; i < int(n); i++ {
		for j := 0; j < int(n); j++ {
			chk[i][j] = false
		}
	}
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
