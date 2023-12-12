package bst

import (
	"fmt"
	"go-playground/datastructure/tree"
)

// PlusOne 所有节点的值 +1
func PlusOne(root *tree.TreeNode) {
	if root == nil {
		return
	}
	root.Value++
	PlusOne(root.Left)
	PlusOne(root.Right)
}

// IsSameTree 是否是相同的树
func IsSameTree(t1 *tree.TreeNode, t2 *tree.TreeNode) bool {
	// 同时为 nil
	if t1 == nil && t2 == nil {
		return true
	}
	// 不同时为 nil
	if t1 == nil || t2 == nil {
		return false
	}
	// 都不是 nil
	if t1.Value != t2.Value {
		return false
	}

	return IsSameTree(t1.Left, t2.Left) && IsSameTree(t1.Right, t2.Right)
}

func IsValidBST(root *tree.TreeNode) bool {
	var helper func(t, min, max *tree.TreeNode) bool

	helper = func(t, min, max *tree.TreeNode) bool {
		if t == nil {
			return true
		}
		// 如果 t 小于最小的，或者大于最大的，那肯定不合法
		if min != nil && t.Value < min.Value || max != nil && t.Value > max.Value {
			return false
		}
		// 到这里，t 在最小值和最大值之间，说明 t 合法
		// 再去看 t 的左右子树是否合法
		return helper(t.Left, min, t) && helper(t.Right, t, max)
	}

	return helper(root, nil, nil)
}

// InOrder 中序遍历
func InOrder(root *tree.TreeNode) {
	var result []int
	var helper func(t *tree.TreeNode)

	helper = func(t *tree.TreeNode) {
		if t == nil {
			return
		}
		if t.Left != nil {
			helper(t.Left)
		}
		result = append(result, t.Value)
		if t.Right != nil {
			helper(t.Right)
		}
	}

	helper(root)
	fmt.Println(result)
}

// InsertValue 向树中插入值，暂不允许插入相同的值
func InsertValue(root *tree.TreeNode, val int) *tree.TreeNode {
	if root == nil {
		root = &tree.TreeNode{Value: val}
		return root
	}
	if val < root.Value {
		root.Left = InsertValue(root.Left, val)
	} else if val > root.Value {
		root.Right = InsertValue(root.Right, val)
	}
	return root
}

// DeleteValue 删除值
func DeleteValue(root *tree.TreeNode, val int) *tree.TreeNode {
	if root == nil {
		return nil
	}
	if val == root.Value {
		// 删除叶子节点（没有孩子），直接删除
		if root.Left == nil && root.Right == nil {
			return nil
		}
		// 删除只有一个孩子的节点，孩子提升
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}
		// 删除有两个孩子的节点，则让左子树的最大值，或者右子树的最小值提升
		minNode := GetMinNode(root.Right)
		// 将 minNode 的值提升
		root.Value = minNode.Value
		// 然后再删除 minNode
		root.Right = DeleteValue(root.Right, minNode.Value)
	} else if val < root.Value {
		root.Left = DeleteValue(root.Left, val)
	} else {
		root.Right = DeleteValue(root.Right, val)
	}
	return root
}

// GetMinNode 获取 BST 的最小节点
func GetMinNode(root *tree.TreeNode) (min *tree.TreeNode) {
	if root == nil {
		return nil
	}
	min = root
	for min.Left != nil {
		min = min.Left
	}
	return min
}

// IsInBST 查找值是否存在
func IsInBST(root *tree.TreeNode, val int) bool {
	if root == nil {
		return false
	}
	if root.Value == val {
		return true
	}
	if val < root.Value {
		return IsInBST(root.Left, val)
	} else {
		return IsInBST(root.Right, val)
	}
}

func Test() {
	// PlusOne
	t1 := &tree.TreeNode{
		Value: 1,
		Left:  nil,
		Right: nil,
	}
	//PlusOne(t1)
	//fmt.Println(t1)

	// InsertValue
	InsertValue(t1, 2)
	InsertValue(t1, 5)
	InsertValue(t1, 3)
	InsertValue(t1, 8)
	InsertValue(t1, 6)
	//InOrder(t1)

	// IsInBST
	//fmt.Println(IsInBST(t1, 3))
	//fmt.Println(IsInBST(t1, 10))

	// IsValidBST
	//fmt.Println(IsValidBST(t1))

	// DeleteValue
	t1 = DeleteValue(t1, 5)
	t1 = DeleteValue(t1, 1)
	InOrder(t1)
}
