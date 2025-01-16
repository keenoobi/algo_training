package main

import (
	"fmt"
)

// Определяем структуру для узла дерева
type TreeNode struct {
	Value       int
	Left, Right *TreeNode
}

// Функция для создания нового узла
func NewTreeNode(value int) *TreeNode {
	return &TreeNode{Value: value, Left: nil, Right: nil}
}

// Функция для добавления узла в бинарное дерево
func (t *TreeNode) Insert(value int) *TreeNode {
	if t == nil {
		return NewTreeNode(value)
	}

	if value < t.Value {
		t.Left = t.Left.Insert(value)
	} else if value > t.Value {
		t.Right = t.Right.Insert(value)
	}

	return t
}

// Функция для обхода дерева в порядке "в глубину" (pre-order)
func (t *TreeNode) PreOrderTraversal() {
	if t == nil {
		return
	}

	fmt.Println(t.Value)
	t.Left.PreOrderTraversal()
	t.Right.PreOrderTraversal()
}

// Функция для поиска элемента в дереве
func (t *TreeNode) Search(value int) *TreeNode {
	if t == nil || t.Value == value {
		return t
	}

	if value < t.Value {
		return t.Left.Search(value)
	}

	return t.Right.Search(value)
}

// Вспомогательная функция для поиска минимального элемента в дереве
func (t *TreeNode) FindMin() *TreeNode {
	current := t
	for current.Left != nil {
		current = current.Left
	}
	return current
}

// Функция для удаления узла из дерева
func (t *TreeNode) Delete(value int) *TreeNode {
	if t == nil {
		return t
	}

	if value < t.Value {
		t.Left = t.Left.Delete(value)
	} else if value > t.Value {
		t.Right = t.Right.Delete(value)
	} else {
		// Узел найден
		if t.Left == nil {
			return t.Right
		} else if t.Right == nil {
			return t.Left
		}

		// Узел имеет двух потомков
		minRight := t.Right.FindMin()
		t.Value = minRight.Value
		t.Right = t.Right.Delete(minRight.Value)
	}

	return t
}

func main() {
	// Создаем корневой узел
	root := NewTreeNode(10)

	// Добавляем узлы в дерево
	root.Insert(5)
	root.Insert(15)
	root.Insert(3)
	root.Insert(7)
	root.Insert(12)
	root.Insert(17)

	// Обходим дерево в порядке "в глубину" (pre-order)
	fmt.Println("Pre-order traversal:")
	root.PreOrderTraversal()

	// Поиск элемента
	fmt.Println("\nSearching for 7:")
	if node := root.Search(7); node != nil {
		fmt.Println("Found:", node.Value)
	} else {
		fmt.Println("Not found")
	}

	// Удаление элемента
	fmt.Println("\nDeleting 15:")
	root = root.Delete(15)

	// Обходим дерево после удаления
	fmt.Println("Pre-order traversal after deletion:")
	root.PreOrderTraversal()
}
