package code_interviews

import (
	"algorithm-lab/common"
	"testing"
)

func TestSerializeBT(t *testing.T) {
	tree := common.NewBinaryTree()
	root := common.NewBinaryTreeNode(1)
	tree.Root = root

	child1 := common.NewBinaryTreeNode(2)
	child2 := common.NewBinaryTreeNode(2)
	child3 := common.NewBinaryTreeNode(4)
	child4 := common.NewBinaryTreeNode(3)
	child5 := common.NewBinaryTreeNode(3)
	child6 := common.NewBinaryTreeNode(4)

	root.SetLeft(child1)
	root.SetRight(child2)

	child1.SetLeft(child3)
	child1.SetRight(child4)

	child2.SetLeft(child5)
	child2.SetRight(child6)

	SerializeBT(tree.Root)
	t.Log(s)
}

//func TestDeserialize(t *testing.T) {
//	str := []string{"1", "2", "4", "$", "$", "3", "$", "$", "2", "3", "$", "$", "4", "$", "$"}
//	root := Deserialize(str)
//	SerializeBT(root)
//	assert.Equal(t, str, s)
//}
