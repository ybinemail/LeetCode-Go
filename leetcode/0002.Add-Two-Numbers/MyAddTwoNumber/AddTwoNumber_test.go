package MyAddTwoNumber

import (
	"testing"
)

/*给出两个非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照逆序的方式存储的，并且它们的每个节点只能存储一位数字。

如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

您可以假设除了数字 0 之外，这两个数都不会以 0开头。

示例：

输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807
*/
func TestAddTwoNumber(t *testing.T) {

	//numOne,numTwo := "5","5"
	//numOne, numTwo := "1", "99"
	//numOne, numTwo := "199", "1"
	numOne, numTwo := "634", "382"

	linkOne := getLinkNode(numOne)
	linkTwo := getLinkNode(numTwo)

	for p := linkOne; ; p = p.Next {
		t.Log(p.Val)
		if p.Next == nil {
			break
		}
	}

	for p := linkTwo; ; p = p.Next {
		t.Log(p.Val)
		if p.Next == nil {
			break
		}
	}

	linksum := addTwoNumbers(linkOne, linkTwo)
	t.Log("===============sum :")
	for p := linksum; ; p = p.Next {
		t.Log(p.Val)
		if p.Next == nil {
			break
		}
	}
}
