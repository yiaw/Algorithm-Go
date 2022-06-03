// https://www.acmicpc.net/problem/10971
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

// 실패 문제
var metrix [][]int32
var chk []int32
var n, min, each int32

func findMetrix(m []int32, v int32) int32 {
	var min int32
	var idx int32
	// 한번 방문한 노드는 선태갛면 안됌
	min = 100000
	for i := 1; i < len(m); i++ {
		if m[i] == 0 || i == int(v) {
			continue
		}

		if m[i] < min {
			min = m[i]
			idx = int32(i)
		}

	}
	return idx
}
func bfs(v int32) int32 {
	var q Que
	var di, dq, min int32
	q.Enque(v)
	chk[v] = 1
	for !q.IsEmpty() {
		dq, _ = q.Deque()
		min = 1000001
		// 방문하지 않았던 NODE를 찾기
		// 길이 있을 경우 찾기
		for i := 1; i < int(n)+1; i++ {
			if metrix[dq][i] > 0 && chk[i] == 0 {
				if metrix[dq][i] < min {
					min = metrix[dq][i]
					di = int32(i)
				}
			}
		}

		if min != 1000001 {
			q.Enque(di)
			chk[di] = chk[dq] + metrix[dq][di]
		}

	}

	if metrix[dq][v] > 0 {
		// 되돌아 오는 길이 존재 한다면
		chk[v] = chk[dq] + metrix[dq][v] - 1

		return chk[v]
	}

	return -1
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	input := readLine(reader)
	nTemp, _ := strconv.ParseInt(input, 10, 64)
	n = int32(nTemp)

	metrix = make([][]int32, n+1)
	chk = make([]int32, n+1)

	var load int32
	for i := 1; i < int(n)+1; i++ {
		metrix[i] = make([]int32, n+1)

		input := strings.Split(readLine(reader), " ")
		for j := 0; j < int(n); j++ {
			mTemp, _ := strconv.ParseInt(input[j], 10, 64)
			metrix[i][j+1] = int32(mTemp)
		}
	}

	for j := 1; j < int(n)+1; j++ {
		reset()
		for i := 1; i < int(n)+1; i++ {
			if metrix[j][i] > 0 && chk[j] == 0 {
				//chk[j] = metrix[j][i]
				each = bfs(int32(j))
				//fmt.Println("each", each)
			}
		}

		if each == -1 {
			continue
		}

		if load == 0 {
			load = each
		}

		if each < load {
			load = each
		}
	}

	fmt.Fprintln(writer, load)
	writer.Flush()
}

func reset() {
	for i := 0; i < int(n)+1; i++ {
		chk[i] = 0
	}

}
func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}
	return strings.TrimRight(string(str), "\r\n")
}

type Que []int32

func (q *Que) IsEmpty() bool {
	return len(*q) == 0
}

func (q *Que) Enque(data int32) {
	*q = append(*q, data)
}

func (q *Que) Deque() (int32, error) {
	if q.IsEmpty() {
		return 0, errors.New("empty")
	}

	data := (*q)[0]
	*q = (*q)[1:]
	return data, nil
}
