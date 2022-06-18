package template

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val int
	Left, Right *TreeNode
}
type ListNode struct {
	Val int
	Next *ListNode
}

func listBuild(nums []int) *ListNode{
	dummyHead := &ListNode{}
	cur := dummyHead
	for _, num := range nums {
		tmp := &ListNode{num, nil}
		cur.Next = tmp
		cur = tmp
	}
	head := dummyHead.Next
	return head
}

//层序遍历
func treeBuild(data string) *TreeNode {
	if data == "null" {
		return nil
	}
	list := strings.Split(data, ",")
	Val, _ := strconv.Atoi(list[0])
	root := &TreeNode{Val: Val}
	q := []*TreeNode{root}
	cursor := 1
	for cursor < len(list) {
		node := q[0]
		q = q[1:]
		leftVal := list[cursor]
		rightVal := list[cursor+1]
		if leftVal != "null" {
			v, _ := strconv.Atoi(leftVal)
			leftNode := &TreeNode{Val: v}
			node.Left = leftNode
			q = append(q, leftNode)
		}
		if rightVal != "null" {
			v, _ := strconv.Atoi(rightVal)
			rightNode := &TreeNode{Val: v}
			node.Right = rightNode
			q = append(q, rightNode)
		}
		cursor += 2
	}
	return root
}

func serialize(root *TreeNode) string {
	q := []*TreeNode{root}
	res := []string{}
	for len(q) != 0 {
		node := q[0]
		q = q[1:]
		fmt.Println(node)
		if node != nil {
			res = append(res, strconv.Itoa(node.Val))
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		} else {
			fmt.Println(node)
			res = append(res, "null")
		}
	}
	return strings.Join(res, ",")
}
