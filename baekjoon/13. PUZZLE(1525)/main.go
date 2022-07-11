package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var metrix [3][3]int
var mx = []int{1, 0, -1, 0}
var my = []int{0, -1, 0, 1}
var dst = "123456780"
var chk map[string]int

// key 는 변경된 string을 확인한 경우, 해당 방문 경로까지 개수

type Value struct {
	cur string
	cnt int
}

func bfs(y, x int) {
	var q Que
	q.Enque(Value{cur: MetrixToString(metrix), cnt: 0})
	chk[MetrixToString(metrix)] = 0

	for !q.IsEmpty() {
		o := q.Deque()
		v := o.(Value)

		if v.cur == dst {
			fmt.Println(v.cnt)
			break
		}

		m := StringToArray(v.cur)
		y, x := findZero(m)
		for i := 0; i < 4; i++ {
			dy := y + my[i]
			dx := x + mx[i]
			if (0 <= dx && dx < 3) && (0 <= dy && dy < 3) {
				tm := swap(m, x, y, dx, dy)
				if _, ok := chk[MetrixToString(tm)]; !ok {
					chk[MetrixToString(tm)] = 0
					q.Enque(Value{cur: MetrixToString(tm), cnt: v.cnt + 1})
				}
			}
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	chk = make(map[string]int)
	for i := 0; i < 3; i++ {
		strTemp := strings.Split(readLine(reader), " ")
		metrix[i][0], _ = strconv.Atoi(strTemp[0])
		metrix[i][1], _ = strconv.Atoi(strTemp[1])
		metrix[i][2], _ = strconv.Atoi(strTemp[2])
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if metrix[i][j] == 0 {
				bfs(i, j)
			}
		}
	}

	fmt.Println(chk[dst])
}

func readLine(reader *bufio.Reader) string {
	buff, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}
	return strings.TrimRight(string(buff), "\r\n")
}

type Que []interface{}

func (q *Que) IsEmpty() bool {
	return len(*q) == 0
}

func (q *Que) Enque(arg interface{}) {
	*q = append(*q, arg)
}

func (q *Que) Deque() interface{} {
	if q.IsEmpty() {
		return nil
	}

	data := (*q)[0]
	(*q) = (*q)[1:]
	return data
}

func StringToArray(str string) [3][3]int {
	for i := 0; i < 9; i++ {

	}
}

func findZero(metrix [3][3]int) (int, int) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if metrix[i][j] == 0 {
				// return y,x
				return i, j
			}
		}
	}
}

func swap(m [3][3]int, x1, y1, x2, y2 int) [3][3]int {
	v := m[x2][y2]
	m[y2][x2] = m[y1][x1]
	m[y1][x1] = v
	return m
}

func MetrixToString(m [3][3]int) string {
	return fmt.Sprintf("%d%d%d%d%d%d%d%d%d",
		m[0][0], m[0][1], m[0][2],
		m[1][0], m[1][1], m[1][2],
		m[2][0], m[2][1], m[2][2])
}
