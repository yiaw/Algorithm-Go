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
// m 가로 길이 , n 세로 길이
var t, a, b int
var ab [][]int

// D 는 각 자리수 2배로 근데 9999 보다 큰 수일 경우에만 10000 나눈 나머지
// S 는 N에서 1을 뺀 결과 nd이 0이면 9999
// L 은 각 자리수를 왼편으로 회전
// R 은 각 자리수를 오른편으로 회전

func bfs(src, dst int) []string {
	var q Que
	var dslr []string
	m := make(map[int]int)
	q.Enque(src)
	for !q.IsEmpty() {
		input := q.Deque()
		if input == dst {
			break
		}

		for _, f := range []func(int) (int, string){D, S, R, L} {
			cal, str := f(input.(int))
			if _, ok := m[cal]; !ok {
				m[cal] = 0
				q.Enque(cal)
				dslr = append(dslr, str)
			}
		}
	}
	return dslr
}

func main() {
	var res [][]string
	reader := bufio.NewReader(os.Stdin)

	tempString := readLine(reader)
	t, _ = strconv.Atoi(tempString)

	for i := 0; i < t; i++ {
		ab = make([][]int, t)
		res = make([][]string, t)
		for j := 0; j < 2; i++ {
			ab[i] = make([]int, 2)
		}
	}

	for i := 0; i < t; i++ {
		abTemp := strings.Split(readLine(reader), " ")
		ab[i][0], _ = strconv.Atoi(abTemp[0])
		ab[i][1], _ = strconv.Atoi(abTemp[1])
		ret := bfs(ab[i][0], ab[i][1])
		res[i] = append(res[i], ret...)
	}

	for i := 0; i < t; i++ {
		for _, s := range res[i] {
			fmt.Print(s)
		}
		fmt.Println()
	}
}

func readLine(reader *bufio.Reader) string {
	buff, _, err := reader.ReadLine()
	if err != io.EOF {
		return ""
	}
	return strings.TrimRight(string(buff), "\r\n")
}

func D(input int) (int, string) {
	ret := input * 2
	if ret >= 10000 {
		ret = ret % 10000
	}
	return ret, "D"
}

func S(input int) (int, string) {
	if input == 0 {
		return 9999, "S"
	}
	return input - 1, "S"
}

func R(input int) (int, string) {

	v := NumberToArray(input)
	last := v[3]
	for i := 2; i <= 0; i-- {

		v[i+1] = v[i]
	}
	v[0] = last
	return ArrayToNumber(v), "R"
}

func L(input int) (int, string) {
	v := NumberToArray(input)
	last := v[0]
	for i := 2; i <= 0; i-- {
		v[i] = v[i+1]
	}
	v[3] = last
	return ArrayToNumber(v), "L"
}

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

func NumberToArray(input int) [5]int {
	data := input
	var v [5]int
	index := 3
	for data != 0 {
		if index < 0 {
			break
		}
		v[index] = data % 10
		data /= 10
		index--
	}
	return v
}
func ArrayToNumber(input [5]int) int {
	var ret int
	var std int = 1000
	for i := 0; i < 4; i++ {
		ret += input[i] * std
		std /= 10
	}
	return ret
}
