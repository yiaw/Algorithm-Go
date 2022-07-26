//https://www.acmicpc.net/problem/15681
package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

var n int
var r int
var q int
var metrix [100001][100001]int
var chk [100001]bool
var cnt [10001]int

func dfs(root int) int {
	if cnt[root] != 0 {
		return cnt[root]
	}

	chk[root] = true

	ret := 1

	for i := 0; i < n-1; i++ {
		next := metrix[root][i]
		if !chk[next] {
			ret += dfs(next)
		}
	}
	cnt[root] = ret
	return ret
}
func main() {
	reader := bufio.NewReader(os.Stdin)

	inputTemp := strings.Split(readLine(reader), " ")

	n, _ = strconv.Atoi(inputTemp[0])
	r, _ = strconv.Atoi(inputTemp[1])
	q, _ = strconv.Atoi(inputTemp[2])

	for i := 1; i < n-1; i++ {
		vTemp := strings.Split(readLine(reader), " ")
		v1, _ := strconv.Atoi(vTemp[0])
		v2, _ := strconv.Atoi(vTemp[1])
		metrix[v1][v2] = 1
		metrix[v2][v1] = 1
	}

	cnt[r] = dfs(r)

}
func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}
	return strings.TrimRight(string(str), "\r\n")
}
