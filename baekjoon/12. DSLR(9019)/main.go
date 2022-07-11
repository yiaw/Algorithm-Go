package main

//https://www.acmicpc.net/problem/9019
import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// 입력값
// t : test case, a : src, b : dst
var t, a, b int
var ab [][2]int
var chk [10000]bool

// D 는 각 자리수 2배로 근데 9999 보다 큰 수일 경우에만 10000 나눈 나머지
// S 는 N에서 1을 뺀 결과 nd이 0이면 9999
// L 은 각 자리수를 왼편으로 회전
// R 은 각 자리수를 오른편으로 회전
type Value struct {
	cur   int    // 현재 int 값
	order string // 여기까지 온 DSLR 계산 방식
}

func bfs(src, dst int) string {
	var q Que
	q.Enque(Value{cur: src})
	chk[src] = true

	for !q.IsEmpty() {
		o := q.Deque()
		v := o.(Value)
		if v.cur == dst {
			return v.order
		}
		for _, f := range []func(int) (int, string){D, S, L, R} {
			vv, str := f(v.cur)
			fmt.Println(v.cur, vv, str)
			if vv == dst {
				return v.order + str
			}
			if !chk[vv] {
				chk[vv] = true
				q.Enque(Value{cur: vv, order: v.order + str})
			}
		}

	}
	return ""
}

func main() {
	var ret []string
	reader := bufio.NewReader(os.Stdin)

	tempString := readLine(reader)
	t, _ = strconv.Atoi(tempString)

	for i := 0; i < t; i++ {
		ab = make([][2]int, t)
	}

	for i := 0; i < t; i++ {
		abTemp := strings.Split(readLine(reader), " ")
		ab[i][0], _ = strconv.Atoi(abTemp[0])
		ab[i][1], _ = strconv.Atoi(abTemp[1])
	}

	for i := 0; i < t; i++ {
		ret = append(ret, bfs(ab[i][0], ab[i][1]))
	}

	for _, s := range ret {
		fmt.Println(s)
	}

}

func readLine(reader *bufio.Reader) string {
	buff, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}
	return strings.TrimRight(string(buff), "\r\n")
}

func D(input int) (int, string) {
	return (input * 2) % 10000, "D"
}

func S(input int) (int, string) {
	if input == 0 {
		return 9999, "S"
	}
	return input - 1, "S"
}

func R(input int) (int, string) {
	// 1234 > 4123
	// 1234 % 10 = 4 * 1000 + 1234/10 = 123 4123

	// 100 > 010 : 10

	return (input%10)*1000 + (input / 10), "R"
}

func L(input int) (int, string) {
	// 1234 > 2341
	// 1234 % 1000 = 234*10 + 1234 / 1000 = 2340 + 1

	// 100 > 001 : 1
	return (input%1000)*10 + (input / 1000), "L"
}

// Que 자료 구조
type Que []interface{}

func (q *Que) IsEmpty() bool {
	return len(*q) == 0
}

func (q *Que) Enque(data interface{}) {
	(*q) = append(*q, data)
}

func (q *Que) Deque() interface{} {
	if (*q).IsEmpty() {
		return nil
	}

	data := (*q)[0]
	*q = (*q)[1:]

	return data
}
