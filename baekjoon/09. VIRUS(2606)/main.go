//https://www.acmicpc.net/problem/2606
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var n, m int32
var metrix [][]int32
var chk []bool
var cnt int32

func dfs(v int32) {
	chk[v] = true
	for i := 1; i < int(n)+1; i++ {
		if metrix[v][i] == 1 && !chk[i] {
			cnt++
			dfs(int32(i))
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	input := readLine(reader)
	nTemp, _ := strconv.ParseInt(input, 10, 64)
	input = readLine(reader)
	mTemp, _ := strconv.ParseInt(input, 10, 64)

	n = int32(nTemp)
	m = int32(mTemp)
	chk = make([]bool, n+1)
	metrix = make([][]int32, n+1)
	for i := 0; i < int(n)+1; i++ {
		metrix[i] = make([]int32, n+1)
	}

	for i := 0; i < int(m); i++ {
		input := strings.Split(readLine(reader), " ")
		a, _ := strconv.ParseInt(input[0], 10, 64)
		b, _ := strconv.ParseInt(input[1], 10, 64)
		metrix[a][b] = 1
		metrix[b][a] = 1
	}

	dfs(1)
	fmt.Fprintln(writer, cnt)
	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}
	return strings.TrimRight(string(str), "\r\n")
}
