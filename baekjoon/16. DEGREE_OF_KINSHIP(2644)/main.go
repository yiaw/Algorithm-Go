package main

//https://www.acmicpc.net/problem/2644
import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// 전체 사람 수 1 <= n  <= 100
var n int

// P1, P2의 번호
var p1, p2 int

// 부모 자실들의 간의 관계 수
var m int

// 부모 자식들의 관계를 인접 행렬로 표현한다.
// 최대 인원수 값이 100명 임으로 index를 101까지 할당
var metrix [101][101]int

// 방문여부 확인
var chk [101]bool
var ans int = -1

// 한명이 주어지고 해당 촌수가 구해지는 것임으로 dfs 깊이 우선으로 탐색
func dfs(p int, cnt int) {
	chk[p] = true
	if p == p2 {
		ans = cnt
		return
	}

	for i := 1; i <= n; i++ {
		if metrix[p][i] == 1 && !chk[i] {
			dfs(i, cnt+1)
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	// 최대 인원 수
	nTemp := readLine(reader)
	n, _ = strconv.Atoi(nTemp)

	// 2명 사람의 번호
	pTemp := strings.Split(readLine(reader), " ")
	p1, _ = strconv.Atoi(pTemp[0])
	p2, _ = strconv.Atoi(pTemp[1])

	// 관계 개수
	mTemp := readLine(reader)
	m, _ = strconv.Atoi(mTemp)

	for i := 1; i <= m; i++ {
		xyTemp := strings.Split(readLine(reader), " ")
		x, _ := strconv.Atoi(xyTemp[0])
		y, _ := strconv.Atoi(xyTemp[1])

		// 위 or 아래로 움직여도 촌수 계산이 됨으로 무방향 그래프로 생각
		metrix[x][y] = 1
		metrix[y][x] = 1
	}

	//DFS 깊이 우선 탐색으로 탐색
	dfs(p1, 0)

	fmt.Printf("%d\n", ans)
}

func readLine(reader *bufio.Reader) string {
	buff, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}
	return strings.TrimRight(string(buff), "\r\n")
}
