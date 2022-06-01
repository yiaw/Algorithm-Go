package main

//https://www.acmicpc.net/problem/1260
import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var dfs_array []int32
var bfs_array []int32
var n, m, s int32
var metric [][]int32
var chk []bool

func dfs(v int32) {
	chk[v] = true
	dfs_array = append(dfs_array, v)

	for i := 1; i <= int(n); i++ {
		if metric[v][i] == 1 && !chk[i] {
			dfs(int32(i))
		}

	}
}

func bfs(v int32) {
	var q Que
	q.Enque(v)
	bfs_array = append(bfs_array, v)
	chk[v] = false
	for !q.IsEmpty() {
		v, _ := q.Deque()
		for i := 1; i <= int(n); i++ {
			if metric[v][i] == 1 && chk[i] {
				q.Enque(int32(i))
				bfs_array = append(bfs_array, int32(i))
				chk[i] = false
			}
		}

	}
}
func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	n = int32(nTemp)

	mTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	m = int32(mTemp)

	sTemp, err := strconv.ParseInt(firstMultipleInput[2], 10, 64)
	checkError(err)
	s = int32(sTemp)
	var edges [][]int32
	for i := 0; i < int(m); i++ {
		edgesRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")
		var edgesRow []int32
		for _, edgesRowItem := range edgesRowTemp {
			edgesItemTemp, err := strconv.ParseInt(edgesRowItem, 10, 64)
			checkError(err)
			edgesItem := int32(edgesItemTemp)
			edgesRow = append(edgesRow, edgesItem)
		}

		if len(edgesRow) != 2 {
			panic("Bad input")
		}

		edges = append(edges, edgesRow)

	}

	metric = make([][]int32, n+1)
	for i := 0; i < int(n)+1; i++ {
		metric[i] = make([]int32, n+1)
	}

	chk = make([]bool, n+1)

	for i := 0; i < int(m); i++ {
		x := edges[i][0]
		y := edges[i][1]
		metric[x][y] = 1
		metric[y][x] = 1
	}

	dfs(s)
	for _, a := range dfs_array {
		fmt.Print(a, " ")

	}
	fmt.Println()

	bfs(s)
	for _, a := range bfs_array {
		fmt.Print(a, " ")
	}
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

type Que []int32

func (q *Que) IsEmpty() bool {
	return len(*q) == 0
}

func (q *Que) Enque(data int32) {
	(*q) = append(*q, data)
}

func (q *Que) Deque() (int32, error) {
	if (*q).IsEmpty() {
		return 0, errors.New("empty")
	}

	data := (*q)[0]
	*q = (*q)[1:]

	return data, nil
}
