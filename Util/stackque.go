package main

import "errors"

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
