package main

import (
	"errors"
	"fmt"
)

type Stack []interface{}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(data interface{}) {
	(*s) = append(*s, data)
}

func (s *Stack) Pop() (interface{}, error) {
	if (*s).IsEmpty() {
		return nil, errors.New("empty")
	}

	top := len(*s) - 1
	data := (*s)[top]

	*s = (*s)[:top]

	return data, nil
}

// O(N) Que
type Que []interface{}

func (q *Que) IsEmpty() bool {
	return len(*q) == 0
}

func (q *Que) Enque(data interface{}) {
	(*q) = append(*q, data)
}

func (q *Que) Deque() (interface{}, error) {
	if (*q).IsEmpty() {
		return nil, errors.New("empty")
	}

	data := (*q)[0]
	*q = (*q)[1:]

	return data, nil
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
func (l *Que) Enque(data interface{}) {
	n := &Node{
		data: data,
	}

	if l.head == nil {
		l.head = n
		l.tail = n
	} else {
		n.prev = l.tail
		l.tail.next = n
		l.tail = n
	}
}

// head 에서 가져오기
func (l *Que) Deque() interface{} {
	if l.head == nil {
		return nil
	}

	del := l.head
	l.head = l.head.next
	return del.data
}

func (l *Que) IsEmpty() bool {
	return l.head == nil
}

func (l *Que) Print() {
	if l.head == nil {
		return
	}

	n := l.head
	for n != nil {
		fmt.Println(n.data)
		n = n.next
	}
}

func main() {
	l := &List{}
	l.Push(1)
	l.Push(2)
	l.Print()

	fmt.Println(l.Pop())
	fmt.Println(l.Pop())
	fmt.Println(l.IsEmpty())

}
