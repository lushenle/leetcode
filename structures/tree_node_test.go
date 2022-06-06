package structures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	// 同一个 TreeNode 的不同表达方式
	//            1
	//      	/  \
	//         2    3
	//        / \  /  \
	//       4  5  6   7
	LeetCodeOrder = []int{1, 2, 3, 4, 5, 6, 7}
	preOrder      = []int{1, 2, 4, 5, 3, 6, 7}
	inOrder       = []int{4, 2, 5, 1, 6, 3, 7}
	postOrder     = []int{4, 5, 2, 6, 7, 3, 1}
)

func Test_preIn2Tree(t *testing.T) {
	ast := assert.New(t)

	actual := Tree2Postorder(PreIn2Tree(preOrder, inOrder))
	expected := postOrder
	ast.Equal(expected, actual)

	ast.Panics(func() { PreIn2Tree([]int{1}, []int{}) })

	ast.Nil(PreIn2Tree([]int{}, []int{}))
}

func Test_indexOf(t *testing.T) {
	ast := assert.New(t)

	ast.Equal(1, indexOf(1, []int{0, 1, 2, 3}))

	ast.Panics(func() { indexOf(0, []int{1, 2, 3}) })
}
