package algorithm_lab

import (
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
}
