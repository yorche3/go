// data_structures_basics.go — Celda enlazada compartida, lista enlazada, pila y cola
//
// Especificación: 06_Data_Structures_Basics
//
// Implementación del contrato: tipos nuevos, firmas y operaciones resueltas.
//
// Indicadores:
//   - el enlace ausente de `Node` es `nil`: solo `Node` devuelve o compara con
//     `nil` (`Next`);
//   - las operaciones que extraen un entero devuelven `int` y su fallo es `-1`;
//   - las banderas devuelven `bool` y su fallo es `false`;
//   - los contadores devuelven `int` y parten de `0`;
//   - las mutaciones (`SetNext`, `InsertHead`, `InsertTail`, `Push`, `Enqueue`)
//     no devuelven valor: mutan la instancia recibida.
//
// Go es mutable, así que las operaciones mutan la instancia en lugar de devolver
// una estructura nueva. El valor cero de `LinkedList`, `Stack` y `Queue` ya es la
// instancia vacía: es el equivalente idiomático del `init` del contrato y no
// necesita preparación previa; la única construcción explícita es `NewNode`.
package src

// ---------------------------------------------------------------------------
// Node — celda enlazada compartida por las tres estructuras
// ---------------------------------------------------------------------------

// Node is the shared linked cell used by LinkedList, Stack, and Queue.
//
// A Node is created with NewNode. Its value is fixed at creation and its link
// can be updated with SetNext. The zero value of Node is not meaningful.
type Node struct {
	value int
	next  *Node
}

// NewNode returns a node holding v with no link (init).
func NewNode(v int) *Node {
	return &Node{value: v}
}

// Value returns the node's value (get_value).
func (n *Node) Value() int {
	return n.value
}

// Next returns the linked node, or nil when the link is absent (get_next).
func (n *Node) Next() *Node {
	return n.next
}

// SetNext updates the node's link (set_next).
func (n *Node) SetNext(next *Node) {
	n.next = next
}

// ---------------------------------------------------------------------------
// LinkedList
// ---------------------------------------------------------------------------

// LinkedList is a singly linked list with head and tail pointers.
//
// The zero value is a ready-to-use empty list.
type LinkedList struct {
	head  *Node
	tail  *Node
	count int
}

// Head returns the head value, or -1 when the list is empty (get_head).
func (l *LinkedList) Head() int {
	if l.IsEmpty() {
		return -1
	}
	return l.head.Value()
}

// InsertHead adds v at the front of the list (insert_head). O(1).
func (l *LinkedList) InsertHead(v int) {
	newNode := NewNode(v)
	newNode.SetNext(l.head)
	l.head = newNode
	if l.tail == nil {
		l.tail = newNode
	}
	l.count++
}

// InsertTail adds v at the end of the list (insert_tail). O(1).
func (l *LinkedList) InsertTail(v int) {
	newNode := NewNode(v)
	if l.tail != nil {
		l.tail.SetNext(newNode)
	}
	l.tail = newNode
	if l.head == nil {
		l.head = newNode
	}
	l.count++
}

// Remove deletes the first occurrence of v (delete); it returns true when a node
// was removed and false when v is absent.
func (l *LinkedList) Remove(v int) bool {
	if l.IsEmpty() {
		return false
	}
	if l.head.Value() == v {
		l.head = l.head.Next()
		if l.head == nil {
			l.tail = nil
		}
		l.count--
		return true
	}
	prev := l.head
	curr := l.head.Next()
	for curr != nil {
		if curr.Value() == v {
			prev.SetNext(curr.Next())
			if curr == l.tail {
				l.tail = prev
			}
			l.count--
			return true
		}
		prev = curr
		curr = curr.Next()
	}
	return false
}

// IsEmpty reports whether the list contains no nodes (is_empty).
func (l *LinkedList) IsEmpty() bool {
	return l.count == 0
}

// Len returns the number of nodes in the list (size).
func (l *LinkedList) Len() int {
	return l.count
}

// ---------------------------------------------------------------------------
// Stack — LIFO sobre el mismo Node
// ---------------------------------------------------------------------------

// Stack is a LIFO structure backed by Node links.
//
// The zero value is a ready-to-use empty stack.
type Stack struct {
	top   *Node
	count int
}

// Push adds v on top of the stack (push). O(1).
func (s *Stack) Push(v int) {
	newNode := NewNode(v)
	newNode.SetNext(s.top)
	s.top = newNode
	s.count++
}

// Pop removes and returns the top value, or -1 when the stack is empty (pop).
func (s *Stack) Pop() int {
	if s.IsEmpty() {
		return -1
	}
	value := s.top.Value()
	s.top = s.top.Next()
	s.count--
	return value
}

// Peek returns the top value without removing it, or -1 when the stack is empty
// (peek).
func (s *Stack) Peek() int {
	if s.IsEmpty() {
		return -1
	}
	return s.top.Value()
}

// IsEmpty reports whether the stack contains no nodes (is_empty).
func (s *Stack) IsEmpty() bool {
	return s.count == 0
}

// Len returns the number of nodes in the stack (size).
func (s *Stack) Len() int {
	return s.count
}

// ---------------------------------------------------------------------------
// Queue — FIFO sobre el mismo Node
// ---------------------------------------------------------------------------

// Queue is a FIFO structure backed by Node links.
//
// The zero value is a ready-to-use empty queue.
type Queue struct {
	front *Node
	rear  *Node
	count int
}

// Enqueue adds v at the rear of the queue (enqueue). O(1).
func (q *Queue) Enqueue(v int) {
	newNode := NewNode(v)
	if q.rear != nil {
		q.rear.SetNext(newNode)
	}
	q.rear = newNode
	if q.front == nil {
		q.front = newNode
	}
	q.count++
}

// Dequeue removes and returns the front value, or -1 when the queue is empty
// (dequeue).
func (q *Queue) Dequeue() int {
	if q.IsEmpty() {
		return -1
	}
	value := q.front.Value()
	q.front = q.front.Next()
	if q.front == nil {
		q.rear = nil
	}
	q.count--
	return value
}

// Peek returns the front value without removing it, or -1 when the queue is
// empty (peek).
func (q *Queue) Peek() int {
	if q.IsEmpty() {
		return -1
	}
	return q.front.Value()
}

// IsEmpty reports whether the queue contains no nodes (is_empty).
func (q *Queue) IsEmpty() bool {
	return q.count == 0
}

// Len returns the number of nodes in the queue (size).
func (q *Queue) Len() int {
	return q.count
}
