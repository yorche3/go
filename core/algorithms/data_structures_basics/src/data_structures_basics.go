// data_structures_basics.go — Celda enlazada compartida, lista enlazada, pila y cola
//
// Especificación: 06_Data_Structures_Basics
//
// Contrato del paso 4b: declara los tipos nuevos y las firmas, y deja el cuerpo
// de cada operación en su indicador natural, sin resolver ningún caso. El
// algoritmo es del paso 5 y la suite, del 4c.
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
	return -1
}

// Next returns the linked node, or nil when the link is absent (get_next).
func (n *Node) Next() *Node {
	return nil
}

// SetNext updates the node's link (set_next).
func (n *Node) SetNext(next *Node) {
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
	return -1
}

// InsertHead adds v at the front of the list (insert_head). O(1).
func (l *LinkedList) InsertHead(v int) {
}

// InsertTail adds v at the end of the list (insert_tail). O(1).
func (l *LinkedList) InsertTail(v int) {
}

// Remove deletes the first occurrence of v (delete); it returns true when a node
// was removed and false when v is absent.
func (l *LinkedList) Remove(v int) bool {
	return false
}

// IsEmpty reports whether the list contains no nodes (is_empty).
func (l *LinkedList) IsEmpty() bool {
	return false
}

// Len returns the number of nodes in the list (size).
func (l *LinkedList) Len() int {
	return 0
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
}

// Pop removes and returns the top value, or -1 when the stack is empty (pop).
func (s *Stack) Pop() int {
	return -1
}

// Peek returns the top value without removing it, or -1 when the stack is empty
// (peek).
func (s *Stack) Peek() int {
	return -1
}

// IsEmpty reports whether the stack contains no nodes (is_empty).
func (s *Stack) IsEmpty() bool {
	return false
}

// Len returns the number of nodes in the stack (size).
func (s *Stack) Len() int {
	return 0
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
}

// Dequeue removes and returns the front value, or -1 when the queue is empty
// (dequeue).
func (q *Queue) Dequeue() int {
	return -1
}

// Peek returns the front value without removing it, or -1 when the queue is
// empty (peek).
func (q *Queue) Peek() int {
	return -1
}

// IsEmpty reports whether the queue contains no nodes (is_empty).
func (q *Queue) IsEmpty() bool {
	return false
}

// Len returns the number of nodes in the queue (size).
func (q *Queue) Len() int {
	return 0
}
