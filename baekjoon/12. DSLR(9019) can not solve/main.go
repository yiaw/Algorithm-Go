package main

//https://www.acmicpc.net/problem/9019
import (
	"bufio"
	"bytes"
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

// D 는 각 자리수 2배로 근데 9999 보다 큰 수일 경우에만 10000 나눈 나머지
// S 는 N에서 1을 뺀 결과 nd이 0이면 9999
// L 은 각 자리수를 왼편으로 회전
// R 은 각 자리수를 오른편으로 회전
type Value struct {
	cur   int          // 현재 int 값
	order bytes.Buffer // 여기까지 온 DSLR 계산 방식
}

func bfs(src, dst int) string {
	var q Que
	var chk [10000]bool
	q.Enque(Value{cur: src})
	chk[src] = true

	for !q.IsEmpty() {
		o := q.Deque()
		v := o.(Value)
		if v.cur == dst {
			return v.order.String()
		}

		for _, f := range []func(int) (int, string){D, S, L, R} {
			vv, str := f(v.cur)
			if vv == dst {
				v.order.WriteString(str)
				return v.order.String()
			}
			if !chk[vv] {
				chk[vv] = true
				ev := Value{
					cur: vv,
				}
				ev.order.WriteString(v.order.String())
				ev.order.WriteString(str)
				q.Enque(ev)
			}
		}

	}
	return ""
}

func main() {
	//var ret []string
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024*24)
	writer := bufio.NewWriterSize(os.Stdout, 1024*1024*24)

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
		fmt.Fprintf(writer, "%s\n", bfs(ab[i][0], ab[i][1]))
	}
	writer.Flush()
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
	return (input%10)*1000 + (input / 10), "R"
}

func L(input int) (int, string) {
	return (input%1000)*10 + (input / 1000), "L"
}

// O(1) Que
type Que struct {
	head *Node
	tail *Node
}

type Node struct {
	next *Node
	prev *Node
	data interface{}
}

// 꼬리에 붙이기
func (q *Que) Enque(data interface{}) {
	n := &Node{
		data: data,
	}

	if q.head == nil {
		q.head = n
		q.tail = n
	} else {
		n.prev = q.tail
		q.tail.next = n
		q.tail = n
	}
}

// head 에서 가져오기
func (q *Que) Deque() interface{} {
	if q.head == nil {
		return nil
	}

	del := q.head
	q.head = q.head.next
	return del.data
}

func (q *Que) IsEmpty() bool {
	return q.head == nil
}
