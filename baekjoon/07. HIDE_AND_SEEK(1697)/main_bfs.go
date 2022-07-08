//https://www.acmicpc.net/problem/11724
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

const MAXCOUNT = 100000

var visted [MAXCOUNT + 1]int32

func bfs(n, k int32) int32 {
	var q Que
	q.Enque(n)
	for !q.IsEmpty() {
		cur, _ := q.Deque()
		if cur == k {
			return visted[cur]
		}

		for i := 0; i < 3; i++ {
			var ncur int32
			switch i {
			case 0:
				ncur = cur + 1
			case 1:
				ncur = cur - 1
			case 2:
				ncur = cur * 2
			}
			if (0 <= ncur && ncur <= MAXCOUNT) && visted[ncur] == 0 {
				visted[ncur] = visted[cur] + 1
				q.Enque(ncur)
			}
		}
	}
	return -1
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	input := strings.Split(readLine(reader), " ")
	nTemp, _ := strconv.ParseInt(input[0], 10, 64)
	kTemp, _ := strconv.ParseInt(input[1], 10, 64)

	n := int32(nTemp)
	k := int32(kTemp)

	ret := bfs(n, k)
	fmt.Fprintln(writer, ret)
	writer.Flush()
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
	if (*q).IsEmpty() {
		return 0, errors.New("empty")
	}
	data := (*q)[0]
	(*q) = (*q)[1:]
	return data, nil
}
