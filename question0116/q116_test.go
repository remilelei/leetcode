package question0116

import (
	"testing"
)

func makeTree(data []int, pos int) *Node {
	if pos < 0 || pos >= len(data) {
		return nil
	}

	node := &Node{Val: data[pos]}
	node.Left = makeTree(data, pos*2+1)
	node.Right = makeTree(data, pos*2+2)
	return node
}

func checkTree(node *Node, data []int, pos int, t *testing.T) {

	if node == nil {
		t.Fatal("empty tree!")
	}

	curNode := node
	newLineNode := node.Left
	dataPos := 0
	for curNode != nil {
		if dataPos >= len(data) {
			t.Fatalf("error pos[%v]", dataPos)
		}
		if curNode.Val != data[dataPos] {
			t.Fatalf("data[%v]=%v but node.Val=%v", dataPos, data[dataPos], curNode.Val)
		}
		dataPos++
		if curNode.Next != nil {
			curNode = curNode.Next
		} else {
			curNode = newLineNode
			if newLineNode != nil {
				newLineNode = newLineNode.Left
			}
		}
	}

}

func TestQuestion116_1(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7}
	root := makeTree(data, 0)
	root = connect(root)
	checkTree(root, data, 0, t)
}
