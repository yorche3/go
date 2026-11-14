// data_structures_advanced.go — Árbol binario de búsqueda y grafo dirigido
//
// Especificación: 07_Data_Structures_Advanced. Vecinos ausentes: nil; vértice
// aislado: slice vacío.
package src

type TreeNode struct {
	value int
	left  *TreeNode
	right *TreeNode
}

func NewTreeNode(value int) *TreeNode {
	return &TreeNode{value: value}
}

func (n *TreeNode) Value() int {
	return n.value
}

func (n *TreeNode) Left() *TreeNode {
	return n.left
}

func (n *TreeNode) Right() *TreeNode {
	return n.right
}

func (n *TreeNode) SetLeft(node *TreeNode) {
	n.left = node
}

func (n *TreeNode) SetRight(node *TreeNode) {
	n.right = node
}

type BinaryTree struct {
	root  *TreeNode
	count int
}

func NewBinaryTree() *BinaryTree {
	return &BinaryTree{}
}

func (t *BinaryTree) Insert(value int) bool {
	if t.root == nil {
		t.root = NewTreeNode(value)
		t.count++
		return true
	}
	current := t.root
	for {
		switch {
		case value == current.value:
			return true
		case value < current.value:
			if current.left == nil {
				current.left = NewTreeNode(value)
				t.count++
				return true
			}
			current = current.left
		default:
			if current.right == nil {
				current.right = NewTreeNode(value)
				t.count++
				return true
			}
			current = current.right
		}
	}
}

func (t *BinaryTree) Contains(value int) bool {
	current := t.root
	for current != nil {
		switch {
		case value == current.value:
			return true
		case value < current.value:
			current = current.left
		default:
			current = current.right
		}
	}
	return false
}

func (t *BinaryTree) IsEmpty() bool {
	return t.count == 0
}

func (t *BinaryTree) Len() int {
	return t.count
}

type GraphNode struct {
	value      int
	neighbors  *neighborNode
	nextVertex *GraphNode
}

type neighborNode struct {
	label int
	next  *neighborNode
}

func newGraphNode(label int) *GraphNode {
	return &GraphNode{value: label}
}

type Graph struct {
	vertices    *GraphNode
	vertexCount int
	edgeCount   int
}

func NewGraph() *Graph {
	return &Graph{}
}

func (g *Graph) FindVertex(label int) *GraphNode {
	for n := g.vertices; n != nil; n = n.nextVertex {
		if n.value == label {
			return n
		}
	}
	return nil
}

func (g *Graph) AddVertex(label int) bool {
	if g.FindVertex(label) != nil {
		return true
	}
	newVertex := newGraphNode(label)
	newVertex.nextVertex = g.vertices
	g.vertices = newVertex
	g.vertexCount++
	return true
}

func insertNeighbor(head *neighborNode, label int) *neighborNode {
	if head == nil || label < head.label {
		return &neighborNode{label: label, next: head}
	}
	head.next = insertNeighbor(head.next, label)
	return head
}

func (g *Graph) AddEdge(u, v int) bool {
	uVertex := g.FindVertex(u)
	if uVertex == nil || g.FindVertex(v) == nil {
		return false
	}
	for n := uVertex.neighbors; n != nil; n = n.next {
		if n.label == v {
			return true
		}
	}
	uVertex.neighbors = insertNeighbor(uVertex.neighbors, v)
	g.edgeCount++
	return true
}

func (g *Graph) HasEdge(u, v int) bool {
	uVertex := g.FindVertex(u)
	if uVertex == nil {
		return false
	}
	for n := uVertex.neighbors; n != nil; n = n.next {
		if n.label == v {
			return true
		}
	}
	return false
}

func (g *Graph) Neighbors(u int) []int {
	uVertex := g.FindVertex(u)
	if uVertex == nil {
		return nil
	}
	var neighbors []int
	for n := uVertex.neighbors; n != nil; n = n.next {
		neighbors = append(neighbors, n.label)
	}
	if neighbors == nil {
		neighbors = []int{}
	}
	return neighbors
}

func (g *Graph) IsEmpty() bool {
	return g.vertexCount == 0
}

func (g *Graph) VertexCount() int {
	return g.vertexCount
}

func (g *Graph) EdgeCount() int {
	return g.edgeCount
}
