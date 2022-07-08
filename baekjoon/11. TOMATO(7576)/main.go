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
}

var my = []int32{-1, 0, 1, 0}
var mx = []int32{0, 1, 0, -1}
var q Que

func bfs() {
	for !q.IsEmpty() {
		l, _ := q.Deque()
		pos := l.(Location)
		for i := 0; i < 4; i++ {
			ix := pos.X + mx[i]
			iy := pos.Y + my[i]
			if (0 <= ix && ix < m) && (0 <= iy && iy < n) {
				if metrix[iy][ix] == 0 {
					metrix[iy][ix] = metrix[pos.Y][pos.X] + 1
					q.Enque(Location{X: ix, Y: iy})
				}
			}
		}
	}
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
			if metrix[i][j] == 1 {
				l := Location{X: int32(j), Y: int32(i)}
				q.Enque(l)
			}
		}
	}

	bfs()
	var ret int32
	var check bool
	for i := 0; i < int(n); i++ {
		for j := 0; j < int(m); j++ {
			if metrix[i][j] == 0 {
				ret = -1
				check = true
				break
			}

			if ret < metrix[i][j] {
				ret = metrix[i][j]
			}
		}
		if check {
			break
		}
	}
	if ret == -1 {
		ret = 0
	}
	fmt.Fprintf(writer, "%d\n", ret-1)
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
