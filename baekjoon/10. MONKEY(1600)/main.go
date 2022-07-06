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

var k, h, w int32
var metrix [201][201]int32
var cnt [201][201]int32
var chk [201][201][31]bool
var mx = []int32{1, 0, -1, 0}
var my = []int32{0, -1, 0, 1}
var hx = []int32{2, 2, 1, 1, -2, -2, -1, -1}
var hy = []int32{1, -1, 2, -2, 1, -1, 2, -2}

type Location struct {
	X int32
	Y int32
}

func bfs(y, x int32) {
	var q Que
	var iy, ix int32

	chk[y][x] = true
	cnt[y][x] = 1
	// Que Append
	q.Enque(Location{X: x, Y: y})

	for !q.IsEmpty() {
		l, _ := q.Deque()
		fmt.Printf("pod location %v\n", l)
		if k > 0 {
			for i := 0; i < 8; i++ {
				iy = l.(Location).Y + hy[i]
				ix = l.(Location).X + hx[i]
				if (0 <= ix && ix < w) && (0 <= iy && iy < h) {
					if metrix[iy][ix] == 0 && !chk[iy][ix] {
						cnt[iy][ix] = cnt[l.(Location).Y][l.(Location).X] + 1
						chk[iy][ix] = true
						q.Enque(Location{X: ix, Y: iy})

						if ix == w && iy == h {
							break
						}
					}
				}
			}
		} else {
			for i := 0; i < 4; i++ {
				iy = l.(Location).Y + my[i]
				ix = l.(Location).X + mx[i]
				if (0 <= ix && ix < w) && (0 <= iy && iy < h) {
					if metrix[iy][ix] == 0 && !chk[iy][ix] {
						cnt[iy][ix] = cnt[l.(Location).Y][l.(Location).X] + 1
						chk[iy][ix] = true
						q.Enque(Location{X: ix, Y: iy})

						if ix == w && iy == h {
							break
						}
					}
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

	if cnt[h][w] == 0 {
		fmt.Fprintf(writer, "-1")
	} else {
		fmt.Fprintf(writer, "%d", cnt[h][w])
	}

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
