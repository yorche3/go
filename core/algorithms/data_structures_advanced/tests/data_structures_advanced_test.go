package tests

import (
	"reflect"
	"testing"

	. "example.com/data_structures_advanced/src"
)

const (
	parentValue = 10
	leftValue   = 20
	rightValue  = 30

	rootValue    = 50
	lowerValue   = 30
	higherValue  = 70
	deepLowLeft  = 20
	deepLowRight = 40
	absentValue  = 99

	labelA      = 10
	labelB      = 30
	labelC      = 50
	labelD      = 70
	missingLow  = 98
	missingHigh = 99
)

func assertNeighbors(t *testing.T, g *Graph, label int, expected []int, subject string) {
	t.Helper()
	got := g.Neighbors(label)
	if got == nil || !reflect.DeepEqual(got, expected) {
		t.Errorf("%s: neighbours of %d should be %v (non-failure), got %#v", subject, label, expected, got)
	}
}

func assertNeighborsFail(t *testing.T, g *Graph, label int, subject string) {
	t.Helper()
	if got := g.Neighbors(label); got != nil {
		t.Errorf("%s: neighbours of %d should be the failure indicator (nil), got %#v", subject, label, got)
	}
}

func assertCounts(t *testing.T, g *Graph, vertices, edges int, subject string) {
	t.Helper()
	if got := g.VertexCount(); got != vertices {
		t.Errorf("%s: vertex count should be %d, got %d", subject, vertices, got)
	}
	if got := g.EdgeCount(); got != edges {
		t.Errorf("%s: edge count should be %d, got %d", subject, edges, got)
	}
}

func assertBool(t *testing.T, got, expected bool, subject string) {
	t.Helper()
	if got != expected {
		t.Errorf("%s should be %v, got %v", subject, expected, got)
	}
}

func newGraphWith(t *testing.T, labels ...int) *Graph {
	t.Helper()
	g := NewGraph()
	if g == nil {
		t.Fatalf("NewGraph should return an initialised graph")
	}
	for _, l := range labels {
		assertBool(t, g.AddVertex(l), true, "AddVertex")
	}
	return g
}

func newTree(t *testing.T) *BinaryTree {
	t.Helper()
	tree := NewBinaryTree()
	if tree == nil {
		t.Fatalf("NewBinaryTree should return an initialised tree")
	}
	return tree
}

func TestTreeNode(t *testing.T) {
	t.Run("Create a tree node", func(t *testing.T) {
		n := NewTreeNode(parentValue)
		if n == nil {
			t.Fatalf("NewTreeNode should return a node")
		}
		if got := n.Value(); got != parentValue {
			t.Errorf("Value should be %d, got %d", parentValue, got)
		}
	})

	t.Run("Absent initial links", func(t *testing.T) {
		n := NewTreeNode(parentValue)
		if n == nil {
			t.Fatalf("NewTreeNode should return a node")
		}
		if n.Left() != nil {
			t.Errorf("Left of a new node should be absent")
		}
		if n.Right() != nil {
			t.Errorf("Right of a new node should be absent")
		}
	})

	t.Run("Link two children", func(t *testing.T) {
		parent := NewTreeNode(parentValue)
		if parent == nil {
			t.Fatalf("NewTreeNode should return a node")
		}
		parent.SetLeft(NewTreeNode(leftValue))
		parent.SetRight(NewTreeNode(rightValue))
		if parent.Left() == nil || parent.Left().Value() != leftValue {
			t.Errorf("left child should hold %d", leftValue)
		}
		if parent.Right() == nil || parent.Right().Value() != rightValue {
			t.Errorf("right child should hold %d", rightValue)
		}
	})

	t.Run("Traverse each link", func(t *testing.T) {
		parent := NewTreeNode(parentValue)
		if parent == nil {
			t.Fatalf("NewTreeNode should return a node")
		}
		parent.SetLeft(NewTreeNode(leftValue))
		parent.SetRight(NewTreeNode(rightValue))
		left, right := parent.Left(), parent.Right()
		if left == nil || right == nil {
			t.Fatalf("both children should be reachable")
		}
		if left.Value() != leftValue || right.Value() != rightValue {
			t.Errorf("children should lead to %d and %d, got %d and %d", leftValue, rightValue, left.Value(), right.Value())
		}
		if left.Left() != nil || left.Right() != nil || right.Left() != nil || right.Right() != nil {
			t.Errorf("children's links should be absent")
		}
	})

	t.Run("Unlink children", func(t *testing.T) {
		parent := NewTreeNode(parentValue)
		if parent == nil {
			t.Fatalf("NewTreeNode should return a node")
		}
		parent.SetLeft(NewTreeNode(leftValue))
		parent.SetRight(NewTreeNode(rightValue))

		parent.SetLeft(nil)
		if parent.Left() != nil {
			t.Errorf("left link should be absent after clearing it")
		}
		if parent.Value() != parentValue {
			t.Errorf("value should stay %d after clearing left, got %d", parentValue, parent.Value())
		}
		if parent.Right() == nil || parent.Right().Value() != rightValue {
			t.Errorf("right child %d should be preserved after clearing left", rightValue)
		}

		parent.SetRight(nil)
		if parent.Left() != nil || parent.Right() != nil {
			t.Errorf("both links should be absent after clearing right")
		}
		if parent.Value() != parentValue {
			t.Errorf("value should stay %d after clearing right, got %d", parentValue, parent.Value())
		}
	})
}

func TestBinaryTree(t *testing.T) {
	assertPresent := func(t *testing.T, tree *BinaryTree, size int, present []int, subject string) {
		t.Helper()
		if got := tree.Len(); got != size {
			t.Errorf("%s: size should be %d, got %d", subject, size, got)
		}
		for _, v := range present {
			if !tree.Contains(v) {
				t.Errorf("%s: tree should contain %d", subject, v)
			}
		}
	}
	insertAll := func(t *testing.T, tree *BinaryTree, values ...int) {
		t.Helper()
		for _, v := range values {
			if !tree.Insert(v) {
				t.Errorf("Insert(%d) should succeed", v)
			}
		}
	}

	t.Run("Create empty", func(t *testing.T) {
		tree := newTree(t)
		assertBool(t, tree.IsEmpty(), true, "IsEmpty")
		assertPresent(t, tree, 0, nil, "empty tree")
		assertBool(t, tree.Contains(rootValue), false, "Contains(50)")
	})

	t.Run("Insert the root", func(t *testing.T) {
		tree := newTree(t)
		insertAll(t, tree, rootValue)
		assertBool(t, tree.IsEmpty(), false, "IsEmpty")
		assertPresent(t, tree, 1, []int{rootValue}, "root only")
	})

	t.Run("Insert lower and higher", func(t *testing.T) {
		tree := newTree(t)
		insertAll(t, tree, rootValue, lowerValue, higherValue)
		assertPresent(t, tree, 3, []int{rootValue, lowerValue, higherValue}, "three values")
	})

	t.Run("Multilevel branches", func(t *testing.T) {
		tree := newTree(t)
		insertAll(t, tree, rootValue, lowerValue, higherValue, deepLowLeft, deepLowRight)
		assertPresent(t, tree, 5, []int{rootValue, lowerValue, higherValue, deepLowLeft, deepLowRight}, "five values")
	})

	t.Run("Missing value in nonempty tree", func(t *testing.T) {
		tree := newTree(t)
		insertAll(t, tree, rootValue)
		assertBool(t, tree.Contains(absentValue), false, "Contains(99)")
		assertPresent(t, tree, 1, []int{rootValue}, "after missing query")
	})

	t.Run("Idempotent duplicate", func(t *testing.T) {
		tree := newTree(t)
		insertAll(t, tree, rootValue, lowerValue, higherValue, lowerValue)
		assertPresent(t, tree, 3, []int{rootValue, lowerValue, higherValue}, "after duplicate")
	})

	t.Run("Ordered insertions", func(t *testing.T) {
		tree := newTree(t)
		insertAll(t, tree, deepLowLeft, lowerValue, deepLowRight, rootValue, higherValue)
		assertPresent(t, tree, 5, []int{deepLowLeft, lowerValue, deepLowRight, rootValue, higherValue}, "ordered values")
		assertBool(t, tree.Contains(absentValue), false, "Contains(99)")
	})
}

func TestGraph(t *testing.T) {
	t.Run("Create empty", func(t *testing.T) {
		g := NewGraph()
		if g == nil {
			t.Fatalf("NewGraph should return an initialised graph")
		}
		assertBool(t, g.IsEmpty(), true, "IsEmpty")
		assertCounts(t, g, 0, 0, "empty graph")
		assertBool(t, g.HasEdge(labelA, labelB), false, "HasEdge(10,30)")
		assertNeighborsFail(t, g, labelA, "empty graph")
	})

	t.Run("Isolated vertex", func(t *testing.T) {
		g := NewGraph()
		if g == nil {
			t.Fatalf("NewGraph should return an initialised graph")
		}
		assertBool(t, g.AddVertex(labelA), true, "AddVertex(10)")
		assertBool(t, g.IsEmpty(), false, "IsEmpty")
		assertCounts(t, g, 1, 0, "isolated vertex")
		assertNeighbors(t, g, labelA, []int{}, "isolated vertex")
	})

	t.Run("Nonconsecutive labels and directed edge", func(t *testing.T) {
		g := newGraphWith(t, labelA, labelD)
		assertBool(t, g.AddEdge(labelA, labelD), true, "AddEdge(10,70)")
		assertBool(t, g.HasEdge(labelA, labelD), true, "HasEdge(10,70)")
		assertBool(t, g.HasEdge(labelD, labelA), false, "HasEdge(70,10)")
		assertCounts(t, g, 2, 1, "directed edge")
	})

	t.Run("Duplicate vertex preserves adjacency", func(t *testing.T) {
		g := newGraphWith(t, labelA, labelB)
		assertBool(t, g.AddEdge(labelA, labelB), true, "AddEdge(10,30)")
		assertBool(t, g.AddVertex(labelA), true, "AddVertex(10) again")
		assertCounts(t, g, 2, 1, "duplicate vertex")
		assertNeighbors(t, g, labelA, []int{labelB}, "duplicate vertex")
	})

	t.Run("Duplicate edge", func(t *testing.T) {
		g := newGraphWith(t, labelA, labelB)
		assertBool(t, g.AddEdge(labelA, labelB), true, "first AddEdge(10,30)")
		assertBool(t, g.AddEdge(labelA, labelB), true, "second AddEdge(10,30)")
		assertCounts(t, g, 2, 1, "duplicate edge")
		assertNeighbors(t, g, labelA, []int{labelB}, "duplicate edge")
	})

	t.Run("Missing endpoint without mutation", func(t *testing.T) {
		g := newGraphWith(t, labelA, labelB)
		assertBool(t, g.AddEdge(labelA, labelB), true, "AddEdge(10,30)")
		assertBool(t, g.AddEdge(labelA, missingHigh), false, "AddEdge(10,99)")
		assertBool(t, g.AddEdge(missingHigh, labelA), false, "AddEdge(99,10)")
		assertBool(t, g.AddEdge(missingLow, missingHigh), false, "AddEdge(98,99)")
		assertBool(t, g.AddEdge(missingHigh, missingHigh), false, "AddEdge(99,99)")
		assertCounts(t, g, 2, 1, "failed edges")
		assertNeighbors(t, g, labelA, []int{labelB}, "failed edges")
		assertNeighbors(t, g, labelB, []int{}, "failed edges")
		assertBool(t, g.HasEdge(labelA, missingHigh), false, "HasEdge(10,99)")
		assertBool(t, g.HasEdge(missingHigh, labelA), false, "HasEdge(99,10)")
		assertBool(t, g.HasEdge(missingLow, missingHigh), false, "HasEdge(98,99)")
		assertBool(t, g.HasEdge(missingHigh, missingHigh), false, "HasEdge(99,99)")
		assertNeighborsFail(t, g, missingLow, "failed edges")
		assertNeighborsFail(t, g, missingHigh, "failed edges")
	})

	t.Run("Sorted neighbours and counts", func(t *testing.T) {
		g := newGraphWith(t, labelA, labelB, labelC, labelD)
		assertBool(t, g.AddEdge(labelA, labelD), true, "AddEdge(10,70)")
		assertBool(t, g.AddEdge(labelA, labelB), true, "AddEdge(10,30)")
		assertBool(t, g.AddEdge(labelA, labelC), true, "AddEdge(10,50)")
		assertBool(t, g.AddEdge(labelB, labelC), true, "AddEdge(30,50)")
		assertNeighbors(t, g, labelA, []int{labelB, labelC, labelD}, "sorted neighbours")
		assertNeighbors(t, g, labelB, []int{labelC}, "sorted neighbours")
		assertNeighbors(t, g, labelC, []int{}, "sorted neighbours")
		assertNeighbors(t, g, labelD, []int{}, "sorted neighbours")
		assertCounts(t, g, 4, 4, "sorted neighbours")
	})

	t.Run("Missing-edge query", func(t *testing.T) {
		g := newGraphWith(t, labelA, labelB)
		assertBool(t, g.HasEdge(labelA, labelB), false, "HasEdge(10,30)")
		assertCounts(t, g, 2, 0, "missing edge")
		assertBool(t, g.IsEmpty(), false, "IsEmpty")
	})

	t.Run("Neighbours of missing vertex", func(t *testing.T) {
		g := newGraphWith(t, labelA)
		assertNeighborsFail(t, g, missingHigh, "missing vertex")
		assertNeighbors(t, g, labelA, []int{}, "existing isolated vertex")
		assertCounts(t, g, 1, 0, "missing vertex query")
	})

	t.Run("Self-loop and duplicate", func(t *testing.T) {
		g := newGraphWith(t, labelA)
		assertBool(t, g.AddEdge(labelA, labelA), true, "first AddEdge(10,10)")
		assertBool(t, g.AddEdge(labelA, labelA), true, "second AddEdge(10,10)")
		assertBool(t, g.HasEdge(labelA, labelA), true, "HasEdge(10,10)")
		assertNeighbors(t, g, labelA, []int{labelA}, "self-loop")
		assertCounts(t, g, 1, 1, "self-loop")
	})

	t.Run("Neighbour-result isolation", func(t *testing.T) {
		g := newGraphWith(t, labelA, labelB, labelC)
		assertBool(t, g.AddEdge(labelA, labelB), true, "AddEdge(10,30)")
		saved := g.Neighbors(labelA)
		if !reflect.DeepEqual(saved, []int{labelB}) {
			t.Fatalf("saved neighbours of 10 should be [30], got %#v", saved)
		}
		other := g.Neighbors(labelA)
		if len(other) > 0 {
			other[0] = absentValue
		}
		assertNeighbors(t, g, labelA, []int{labelB}, "after modifying a returned copy")
		assertBool(t, g.AddEdge(labelA, labelC), true, "AddEdge(10,50)")
		if !reflect.DeepEqual(saved, []int{labelB}) {
			t.Errorf("saved result should remain [30], got %#v", saved)
		}
		assertNeighbors(t, g, labelA, []int{labelB, labelC}, "fresh query")
		assertCounts(t, g, 3, 2, "isolation")
	})
}
