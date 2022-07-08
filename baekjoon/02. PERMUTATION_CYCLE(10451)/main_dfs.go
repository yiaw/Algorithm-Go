//https://www.acmicpc.net/problem/10451

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var t int32
var n int32
var edges [][]int32
var metric [][]int32
var chk []bool
var cnt int32

func dfs(v int32) {
	chk[v] = true

	for i := 1; i <= int(n); i++ {
		if metric[v][i] == 1 && !chk[i] {
			dfs(int32(i))
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	input := readLine(reader)
	tTemp, _ := strconv.ParseInt(input, 10, 64)
	t = int32(tTemp)

	for i := 0; i < int(t); i++ {
		input = readLine(reader)
		nTemp, _ := strconv.ParseInt(input, 10, 64)
		n = int32(nTemp)
		chk = make([]bool, n+1)
		metric = make([][]int32, n+1)
		for j := 0; j < int(n)+1; j++ {
			metric[j] = make([]int32, n+1)
		}

		egs := strings.Split(readLine(reader), " ")

		for j := 0; j < len(egs); j++ {
			vTemp, _ := strconv.ParseInt(egs[j], 10, 64)
			v := int(vTemp)
			metric[v][j+1] = 1
		}

		cnt = 0
		for v := 1; v <= int(n); v++ {
			if !chk[v] {
				dfs(int32(v))
				cnt++
			}
		}
		fmt.Fprintln(writer, cnt)
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
