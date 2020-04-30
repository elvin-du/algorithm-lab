package algorithm_lab

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReconstructBinaryTree(t *testing.T) {
	preorder := []int{1, 2, 4, 7, 3, 5, 6, 8}
	inorder := []int{4, 7, 2, 1, 5, 3, 8, 6}
	assert.Equal(t, "12473568", ReconstructBinaryTree(preorder, inorder).PreOrder())
	assert.Equal(t, "47215386", ReconstructBinaryTree(preorder, inorder).InOrder())
}
