package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

//https://www.acmicpc.net/problem/1389

// N : 유저의 수 2  <=  N <= 100
// M : 관계의 수
var n, m int

// 관계 그래프를 인접 행렬로 표시
// n + 1 에는 n열의 값의 합을 구함
var metrix [101][101]int

var cnt [101][101]int

// 해당 유저의 케빈 베이컨  6단계 법칙에 의해 구해진 수
// cnt[a][b] 라고 가정 하면
// a 기준으로 b를 방문 할때 몇단계인지 기입
// a / b로 기준으로하여 제일 적은 수가 기입된 a / b 값을 반환

func bfs(n int) int {
	var q Que
	var chk [101]bool
	q.Enque(n)

	for !q.IsEmpty() {
		for i := 1; i <= n; i++ {
			if !chk[i] {

			}
		}

	}
}
func main() {
	reader := bufio.NewReader(os.Stdin)

	// 사람의 수, 관계의 수 입력
	nmTemp := strings.Split(readLine(reader), " ")

	n, _ = strconv.Atoi(nmTemp[0])
	m, _ = strconv.Atoi(nmTemp[1])

	for i := 0; i < m; i++ {
		// A-B의 관계 입력
		abTemp := strings.Split(readLine(reader), " ")
		a, _ := strconv.Atoi(abTemp[0])
		b, _ := strconv.Atoi(abTemp[1])

		// 무방향 그래프
		metrix[a][b] = 1
		metrix[b][a] = 1
	}

	for i := 1; i <= n; i++ {
		// 1번 user부터 시작
		ret = append(ret, bfs(i))
	}

}

func readLine(reader *bufio.Reader) string {
	buff, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}
	return strings.TrimRight(string(buff), "\r\n")
}

type Que []interface{}

func (q *Que) IsEmpty() bool {
	return len(*q) == 0
}
func (q *Que) Enque(data interface{}) {
	*q = append(*q, data)
}
func (q *Que) Deque() interface{} {
	data := (*q)[0]
	(*q) = (*q)[1:]
	return data
}
