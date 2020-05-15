package algorithm_lab

import (
	"algorithm-lab/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateBST(t *testing.T) {
	bt := CreateBST([]int{1, 2, 3, 4, 7, 11, 6, 9})
	assert.Equal(t, "1,2,3,4,7,6,11,9", bt.BFS())
	assert.Equal(t, "1,$,2,$,3,$,4,$,7,6,$,$,11,9,$,$", bt.Serialize())
}

func TestFindKthNode(t *testing.T) {
	bt := CreateBST([]int{1, 7, 8, 4, 11, 6, 9})
	k := 3
	assert.Equal(t, 6, FindKthNode(bt.Root, &k).Value)

	k = 7
	assert.Equal(t, 11, FindKthNode(bt.Root, &k).Value)

	k = 10
	assert.Equal(t, true, common.IsNil(FindKthNode(bt.Root, &k)))
}

func TestFindKthNode2(t *testing.T) {
	bt := CreateBST([]int{1, 7, 8, 4, 11, 6, 9})
	k := 3
	ret := &common.BinaryTreeNode{}
	FindKthNode2(bt.Root, &k, ret)
	assert.Equal(t, 6, ret.Value)

	k = 7
	FindKthNode2(bt.Root, &k, ret)
	assert.Equal(t, 11, ret.Value)

	k = 10
	var ret2 *common.BinaryTreeNode = nil
	FindKthNode2(bt.Root, &k, ret2)
	assert.Equal(t, true, common.IsNil(ret2))
}
