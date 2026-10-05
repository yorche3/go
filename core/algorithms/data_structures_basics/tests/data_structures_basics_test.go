package tests

import (
	"reflect"
	"testing"

	. "example.com/data_structures_basics/src"
)

const (
	nodeAValue = 10
	nodeBValue = 20

	listHeadValue   = 5
	listFirstValue  = 10
	listMiddleValue = 20
	listAbsentValue = 99

	stackFirstValue  = 10
	stackSecondValue = 20
	stackThirdValue  = 30
	stackReuseValue  = 40

	queueFirstValue  = 10
	queueSecondValue = 20
	queueThirdValue  = 30
	queueReuseValue  = 40

	failureValue = -1
)

type contractCase struct {
	description string
	operation   func() any
	expected    any
}

type nodeOutput struct {
	value      int
	nextValue  int
	nextAbsent bool
}

type listOutput struct {
	head    int
	empty   bool
	length  int
	removed bool
}

type stackOutput struct {
	peek   int
	popped []int
	empty  bool
	length int
}

type queueOutput struct {
	peek     int
	dequeued []int
	empty    bool
	length   int
}

func assertAllCases(t *testing.T, subject string, cases []contractCase) {
	t.Helper()

	for _, c := range cases {
		t.Run(c.description, func(t *testing.T) {
			actual := c.operation()
			if !reflect.DeepEqual(c.expected, actual) {
				t.Errorf("%s should return %#v for %s, got %#v", subject, c.expected, c.description, actual)
			}
		})
	}
}

func TestNode(t *testing.T) {
	var a *Node

	nodeCases := []contractCase{
		{
			description: "initialize and observe value and link",
			operation: func() any {
				a = NewNode(nodeAValue)
				return nodeOutput{value: a.Value(), nextAbsent: a.Next() == nil}
			},
			expected: nodeOutput{value: nodeAValue, nextAbsent: true},
		},
		{
			description: "initialize another node, link and traverse",
			operation: func() any {
				b := NewNode(nodeBValue)
				a.SetNext(b)
				return nodeOutput{
					value:      a.Next().Value(),
					nextAbsent: b.Next() == nil,
				}
			},
			expected: nodeOutput{value: nodeBValue, nextAbsent: true},
		},
	}

	assertAllCases(t, "node", nodeCases)
}

func TestLinkedList(t *testing.T) {
	var list LinkedList

	linkedListCases := []contractCase{
		{
			description: "empty state",
			operation: func() any {
				return listOutput{head: list.Head(), empty: list.IsEmpty(), length: list.Len()}
			},
			expected: listOutput{head: failureValue, empty: true, length: 0},
		},
		{
			description: "insert at both ends",
			operation: func() any {
				list.InsertTail(listFirstValue)
				list.InsertTail(listMiddleValue)
				list.InsertHead(listHeadValue)
				list.InsertTail(listFirstValue)
				return listOutput{head: list.Head(), length: list.Len()}
			},
			expected: listOutput{head: listHeadValue, length: 4},
		},
		{
			description: "delete first occurrence",
			operation: func() any {
				removed := list.Remove(listFirstValue)
				return listOutput{head: list.Head(), length: list.Len(), removed: removed}
			},
			expected: listOutput{head: listHeadValue, length: 3, removed: true},
		},
		{
			description: "absent value",
			operation: func() any {
				removed := list.Remove(listAbsentValue)
				return listOutput{head: list.Head(), length: list.Len(), removed: removed}
			},
			expected: listOutput{head: listHeadValue, length: 3, removed: false},
		},
		{
			description: "empty the list",
			operation: func() any {
				removed := []bool{
					list.Remove(listHeadValue),
					list.Remove(listMiddleValue),
					list.Remove(listFirstValue),
				}
				return struct {
					listOutput
					removed []bool
				}{
					listOutput: listOutput{head: list.Head(), empty: list.IsEmpty(), length: list.Len()},
					removed:    removed,
				}
			},
			expected: struct {
				listOutput
				removed []bool
			}{
				listOutput: listOutput{head: failureValue, empty: true, length: 0},
				removed:    []bool{true, true, true},
			},
		},
	}

	assertAllCases(t, "linked_list", linkedListCases)
}

func TestStack(t *testing.T) {
	var stack Stack

	stackCases := []contractCase{
		{
			description: "empty state and failed removal",
			operation: func() any {
				return stackOutput{
					peek:   stack.Peek(),
					popped: []int{stack.Pop()},
					empty:  stack.IsEmpty(),
					length: stack.Len(),
				}
			},
			expected: stackOutput{peek: failureValue, popped: []int{failureValue}, empty: true, length: 0},
		},
		{
			description: "LIFO and non-mutating peek",
			operation: func() any {
				stack.Push(stackFirstValue)
				stack.Push(stackSecondValue)
				stack.Push(stackThirdValue)
				return stackOutput{peek: stack.Peek(), length: stack.Len()}
			},
			expected: stackOutput{peek: stackThirdValue, length: 3},
		},
		{
			description: "removal and reuse",
			operation: func() any {
				first := stack.Pop()
				stack.Push(stackReuseValue)
				return stackOutput{
					popped: []int{first, stack.Pop(), stack.Pop(), stack.Pop()},
					empty:  stack.IsEmpty(),
					length: stack.Len(),
				}
			},
			expected: stackOutput{
				popped: []int{stackThirdValue, stackReuseValue, stackSecondValue, stackFirstValue},
				empty:  true,
				length: 0,
			},
		},
		{
			description: "empty after removal",
			operation: func() any {
				return stackOutput{popped: []int{stack.Pop()}, empty: stack.IsEmpty()}
			},
			expected: stackOutput{popped: []int{failureValue}, empty: true},
		},
	}

	assertAllCases(t, "stack", stackCases)
}

func TestQueue(t *testing.T) {
	var queue Queue

	queueCases := []contractCase{
		{
			description: "empty state and failed removal",
			operation: func() any {
				return queueOutput{
					peek:     queue.Peek(),
					dequeued: []int{queue.Dequeue()},
					empty:    queue.IsEmpty(),
					length:   queue.Len(),
				}
			},
			expected: queueOutput{peek: failureValue, dequeued: []int{failureValue}, empty: true, length: 0},
		},
		{
			description: "FIFO and non-mutating peek",
			operation: func() any {
				queue.Enqueue(queueFirstValue)
				queue.Enqueue(queueSecondValue)
				queue.Enqueue(queueThirdValue)
				return queueOutput{peek: queue.Peek(), length: queue.Len()}
			},
			expected: queueOutput{peek: queueFirstValue, length: 3},
		},
		{
			description: "removal and reuse",
			operation: func() any {
				first := queue.Dequeue()
				queue.Enqueue(queueReuseValue)
				return queueOutput{
					dequeued: []int{first, queue.Dequeue(), queue.Dequeue(), queue.Dequeue()},
					empty:    queue.IsEmpty(),
					length:   queue.Len(),
				}
			},
			expected: queueOutput{
				dequeued: []int{queueFirstValue, queueSecondValue, queueThirdValue, queueReuseValue},
				empty:    true,
				length:   0,
			},
		},
		{
			description: "empty after removal",
			operation: func() any {
				return queueOutput{dequeued: []int{queue.Dequeue()}, empty: queue.IsEmpty()}
			},
			expected: queueOutput{dequeued: []int{failureValue}, empty: true},
		},
	}

	assertAllCases(t, "queue", queueCases)
}
