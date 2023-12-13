package bst

import (
	. "github.com/smartystreets/goconvey/convey"
	"go-playground/datastructure/tree"
	"reflect"
	"testing"
)

func TestBiTree(t *testing.T) {
	var root *tree.TreeNode

	Convey("insert a value into empty bst", t, func() {
		root = InsertValue(root, 5)

		So(root, ShouldNotBeNil)
		So(root.Value, ShouldEqual, 5)
	})

	Convey("insert some values into bst and use in-order traversal to output slices", t, func() {
		root = InsertValue(root, 4)
		root = InsertValue(root, 7)
		root = InsertValue(root, 12)
		root = InsertValue(root, 8)
		root = InsertValue(root, 2)
		list := InOrder(root)

		So(reflect.DeepEqual(list, []int{2, 4, 5, 7, 8, 12}), ShouldBeTrue)
	})

	Convey("find value in bst", t, func() {
		So(IsInBST(root, 7), ShouldBeTrue)
		So(IsInBST(root, 3), ShouldBeFalse)
	})

	Convey("plus one", t, func() {
		PlusOne(root)
		list := InOrder(root)

		So(reflect.DeepEqual(list, []int{3, 5, 6, 8, 9, 13}), ShouldBeTrue)
	})

	Convey("is valid bst", t, func() {
		So(IsValidBST(root), ShouldBeTrue)
	})

	Convey("is same tree", t, func() {
		var t2 *tree.TreeNode
		t2 = InsertValue(t2, 6)
		t2 = InsertValue(t2, 5)
		t2 = InsertValue(t2, 8)
		t2 = InsertValue(t2, 13)
		t2 = InsertValue(t2, 9)
		t2 = InsertValue(t2, 3)

		So(IsSameTree(root, t2), ShouldBeTrue)
	})

	Convey("delete a value", t, func() {
		DeleteValue(root, 6)
		list := InOrder(root)

		So(reflect.DeepEqual(list, []int{3, 5, 8, 9, 13}), ShouldBeTrue)
	})
}
