//https://www.acmicpc.net/problem/2667

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

var chk [][]bool
var metric [][]int32
var n int32
var cnt int32
var result []int
var innerCount int32
var dx = []int32{0, 1, 0, -1}
var dy = []int32{1, 0, -1, 0}

func dfs(x, y int32) {
	innerCount++
	for i := 0; i < 4; i++ {
		ny := y + dy[i]
		nx := x + dx[i]
		if (0 <= nx && nx < n) && (0 <= ny && ny < n) {
			if metric[ny][nx] == 1 && !chk[ny][nx] {
				chk[ny][nx] = true
				dfs(nx, ny)
			}
		}

	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	input := strings.Split(readLine(reader), " ")

	nTemp, _ := strconv.ParseInt(input[0], 10, 64)
	n = int32(nTemp)

	metric = make([][]int32, n)
	chk = make([][]bool, n)

	for i := 0; i < int(n); i++ {
		metric[i] = make([]int32, n)
		chk[i] = make([]bool, n)
	}

	for i := 0; i < int(n); i++ {
		str := readLine(reader)
		for j := 0; j < int(n); j++ {
			mTemp, _ := strconv.ParseInt(string(str[j]), 10, 64)
			metric[i][j] = int32(mTemp)
		}
	}

	for i := 0; i < int(n); i++ {
		for j := 0; j < int(n); j++ {
			if metric[i][j] == 1 && !chk[i][j] {
				chk[i][j] = true
				cnt++
				dfs(int32(j), int32(i))
				result = append(result, int(innerCount))
				innerCount = 0
			}
		}
	}

	fmt.Fprintln(writer, cnt)
	sort.Ints(result)
	for _, v := range result {
		fmt.Fprintln(writer, v)
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
