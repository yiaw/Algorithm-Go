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
// m 가로 길이 , n 세로 길이
var m, n int32

// 출력
// 모두 익지 못하는 case의 경우 -1
// 이미 모두 익은   case의 경우 1
// 모두 익은 상태의 case의 경우 해당 일수

// 행렬 그래프
var metrix [1001][1001]int32

// 좌표
type Location struct {
	X int32 // x 좌표
	Y int32 // y 좌표
	M int32 // 최소 일수
}

var chk [1001][1001]bool

func bfs(y, x int32) int32 {
	var q Que
	var pos Location
	var my = []int32{-1, 0, 1, 0}
	var mx = []int32{0, 1, 0, -1}

	l := Location{X: x, Y: y, M: 0}
	chk[y][x] = true
	q.Enque(l)

	for !q.IsEmpty() {
		l, _ := q.Deque()
		pos = l.(Location)
		fmt.Printf("x:%d y:%d m:%d\n", pos.X, pos.Y, pos.M)
		for i := 0; i < 4; i++ {
			ix := pos.X + mx[i]
			iy := pos.Y + my[i]
			if (0 <= ix && ix < m) && (0 <= iy && iy < n) {
				if metrix[iy][ix] == 0 && !chk[iy][ix] {
					chk[iy][ix] = true
					q.Enque(Location{X: ix, Y: iy, M: pos.M + 1})
				}
			}
		}
	}
	return pos.M + 1
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	writer := bufio.NewWriterSize(os.Stdout, 16*1024*1024)

	tempString := strings.Split(strings.TrimSpace(readLine(reader)), " ")
	mTemp, err := strconv.ParseInt(tempString[0], 10, 64)
	checkError(err)
	m = int32(mTemp)
	nTemp, err := strconv.ParseInt(tempString[1], 10, 64)
	checkError(err)
	n = int32(nTemp)

	for i := 0; i < int(n); i++ {
		str := strings.Split(strings.TrimSpace(readLine(reader)), " ")
		for j := 0; j < int(m); j++ {
			tempValue, _ := strconv.ParseInt(string(str[j]), 10, 64)
			metrix[i][j] = int32(tempValue)
		}

	}

	for i := 0; i < int(n); i++ {
		for j := 0; j < int(m); j++ {
			if metrix[i][j] == 0 && !chk[i][j] {
				move := bfs(int32(i), int32(j))
				fmt.Printf("%d\n", move)
			}
		}
	}

	//fmt.Fprintf(writer, "%d\n", cnt)

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
