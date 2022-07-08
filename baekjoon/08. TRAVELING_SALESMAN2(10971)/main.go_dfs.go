// https://www.acmicpc.net/problem/10971
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

var metrix [][]int32
var chk []bool
var n int32
var result []int

func bt(v, cnt, cost int32) {

	if cnt == n {
		// 모든 노드를 방문
		// 현재 노드에서 초기노드로 이동 가능한지 확인
		if metrix[v][1] > 0 {
			// cost 값을 저장
			result = append(result, int(cost)+int(metrix[v][1]))
			return
		}
		// 마지막 노드에서 이동할 값이 없음 ..
		return
	}
	for i := 1; i < int(n)+1; i++ {
		if metrix[v][i] > 0 && !chk[i] {
			chk[i] = true
			bt(int32(i), cnt+1, cost+metrix[v][i])
			chk[i] = false
		}
	}
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	input := readLine(reader)
	nTemp, _ := strconv.ParseInt(input, 10, 64)
	n = int32(nTemp)

	metrix = make([][]int32, n+1)
	chk = make([]bool, n+1)

	for i := 1; i < int(n)+1; i++ {
		metrix[i] = make([]int32, n+1)
		input := strings.Split(readLine(reader), " ")
		for j := 0; j < int(n); j++ {
			mTemp, _ := strconv.ParseInt(input[j], 10, 64)
			metrix[i][j+1] = int32(mTemp)
		}
	}

	cnt := int32(1)
	cost := int32(0)
	chk[1] = true
	bt(1, cnt, cost)

	sort.Ints([]int(result))

	fmt.Fprintln(writer, result[0])
	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}
	return strings.TrimRight(string(str), "\r\n")
}
