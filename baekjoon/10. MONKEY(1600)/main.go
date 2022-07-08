package main

//https://www.acmicpc.net/problem/1600
import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// 결과
var ret int32 = -1

// 입력값
var k, h, w int32

// 행렬 그래프
var metrix [201][201]int32

// 방문 노드 확인 + 말의 이동 사용 여부
var chk [201][201][31]bool

// 원숭이 이동
var mx = []int32{1, 0, -1, 0}
var my = []int32{0, -1, 0, 1}

// 말 이동
var hx = []int32{2, 2, 1, 1, -2, -2, -1, -1}
var hy = []int32{1, -1, 2, -2, 1, -1, 2, -2}

// 좌표
type Location struct {
	X int32 // x 좌표
	Y int32 // y 좌표
	A int32 // x,y 좌표까지 이동 시 말의 능력을 몇번 사용했는지 확인
	M int32 // 현재까지 지나온 거리
}

func bfs(y, x int32) {
	var q Que
	var iy, ix int32

	chk[y][x][0] = true

	// Que Append
	q.Enque(Location{X: x, Y: y, A: 0, M: 0})

	for !q.IsEmpty() {
		l, _ := q.Deque()
		if l.(Location).X == w-1 && l.(Location).Y == h-1 {
			ret = l.(Location).M
			break
		}

		if l.(Location).A < k {
			for i := 0; i < 8; i++ {
				iy = l.(Location).Y + hy[i]
				ix = l.(Location).X + hx[i]
				if (0 <= ix && ix < w) && (0 <= iy && iy < h) {
					if metrix[iy][ix] == 0 && !chk[iy][ix][l.(Location).A+1] {
						chk[iy][ix][l.(Location).A+1] = true
						q.Enque(Location{X: ix, Y: iy, A: l.(Location).A + 1, M: l.(Location).M + 1})

					}
				}
			}
		}

		for i := 0; i < 4; i++ {
			iy = l.(Location).Y + my[i]
			ix = l.(Location).X + mx[i]
			if (0 <= ix && ix < w) && (0 <= iy && iy < h) {
				if metrix[iy][ix] == 0 && !chk[iy][ix][l.(Location).A] {
					chk[iy][ix][l.(Location).A] = true
					q.Enque(Location{X: ix, Y: iy, A: l.(Location).A, M: l.(Location).M + 1})

				}
			}
		}
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	writer := bufio.NewWriterSize(os.Stdout, 16*1024*1024)

	kString := strings.Split(strings.TrimSpace(readLine(reader)), " ")
	kTemp, err := strconv.ParseInt(kString[0], 10, 64)
	checkError(err)
	k = int32(kTemp)

	whString := strings.Split(strings.TrimSpace(readLine(reader)), " ")
	wTemp, err := strconv.ParseInt(whString[0], 10, 64)
	checkError(err)
	w = int32(wTemp)
	hTemp, err := strconv.ParseInt(whString[1], 10, 64)
	checkError(err)
	h = int32(hTemp)

	for i := 0; i < int(h); i++ {
		str := strings.Split(strings.TrimSpace(readLine(reader)), " ")
		for j := 0; j < int(w); j++ {
			mTemp, _ := strconv.ParseInt(string(str[j]), 10, 64)
			metrix[i][j] = int32(mTemp)
		}

	}

	bfs(0, 0)

	fmt.Fprintf(writer, "%d\n", ret)

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
