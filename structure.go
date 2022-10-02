package main

type LinkList struct {
	root   *LinkListNode
	length uint
}
type LinkListNode struct {
	Value interface{}
	Next  *LinkListNode
}

func (ll *LinkList) AddTail(value interface{}) LinkList {
	newNode := LinkListNode{
		Value: value,
	}
	if ll.root == nil {
		ll.root = &newNode
		return *ll
	}

	current := ll.root
	for current.Next != nil {
		current = current.Next
	}
	current.Next = &newNode
	ll.length += 1

	return *ll
}

type Queue struct {
	arr []any
}

func (q *Queue) Push(data any) Queue {
	q.arr = append(q.arr, data)
	return *q
}

func (q *Queue) Pop() any {
	data := q.arr[0]
	q.arr = q.arr[1:]
	return data
}

func (q *Queue) IsEmpty() bool {
	return len(q.arr) == 0
}

// func test123() {
// 	var q Queue
// 	q.Push(123)
// 	q.Push(321)
// 	fmt.Println(q)
// 	q.Pop()
// 	fmt.Println(q)
// }
