package main

//https://www.acmicpc.net/problem/1987
import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// 세로 R
// 가로 c  // 1<= R,C <= 20
var r, c int
var ans int

// 입력 값
var metrix [21][21]string

// 이동 방향은 오른쪽, 아래쪽, 왼쪽 위쪽
var my = [4]int{0, -1, 0, 1}
var mx = [4]int{1, 0, -1, 0}

// 경로 위치를 표시할 인덱스 구조체
type position struct {
	x int
	y int
}

// 알파벳 방문 확인
//var chk = make(map[string]struct{})

var chk [100]bool

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func dfs(y, x, cnt int) {

	// 말이 시작된 알파벳을 방문했다고 표시한다.
	ans = max(ans, cnt)
	for i := 0; i < 4; i++ {
		dy := y + my[i]
		dx := x + mx[i]

		// 말의 이동경로가 metrix 범위를 벗어나서는 안된다.
		if (0 <= dx && dx < c) && (0 <= dy && dy < r) {
			// 방문한적 없는 위치일 경우
			as := []rune(metrix[dy][dx])
			if !chk[as[0]] {

				chk[as[0]] = true
				dfs(dy, dx, cnt+1)
				// 다른 깊이 우선 탐색을 위해 방문 경로에서 다시 뺀다.
				chk[as[0]] = false

			}
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	// 먼저 공백으로 구분된 두개의 RC를 받아야함
	tempInput := strings.Split(readLine(reader), " ")

	rTemp, _ := strconv.ParseInt(tempInput[0], 10, 64)
	cTemp, _ := strconv.ParseInt(tempInput[1], 10, 64)

	r = int(rTemp)
	c = int(cTemp)

	// 문자열로 이루어진 metrix를 input 을 받아야함

	for i := 0; i < r; i++ {
		tempInput := readLine(reader)
		alpha := []rune(tempInput)
		for j := 0; j < c; j++ {
			metrix[i][j] = string(alpha[j])
		}
	}

	// metrix 입력값 확인
	//fmt.Println(metrix)

	// 시작점의 위치도 포함시킨다.
	as := []rune(metrix[0][0])
	chk[as[0]] = true
	// 말의 시작 위치는 좌측  상단
	dfs(0, 0, 1)

	fmt.Println(ans)
}

func readLine(reader *bufio.Reader) string {
	buff, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}
	return strings.TrimRight(string(buff), "\r\n")
}
