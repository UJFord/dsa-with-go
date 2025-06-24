package main

import (
	"math"
)

func LinearSearch(list []int, target int) bool {

	for i := 0; i < len(list); i++ {

		if list[i] == target {
			return true
		}

	}

	return false
}

func BinarySearch(list []int, target int) bool {

	hi := len(list)
	lo := 0

	for lo < hi {
		m := lo + ((hi - lo) / 2)

		if list[m] == target {
			return true
		}
		if list[m] > target {
			hi = m
		}
		if list[m] < target {
			lo = m + 1
		}
	}

	return false
}

func CrystalBall(broke []bool) int {

	n := len(broke)
	jmpAmount := int(math.Floor(math.Sqrt(float64(n))))

	var i = jmpAmount
	for ; i < n; i += jmpAmount {

		if broke[i] {

			break
		}
	}

	i -= jmpAmount

	for j := 0; j <= jmpAmount && j < n; j++ {

		if broke[i] {

			return i
		}
		i++
	}

	return -1
}

func BubbleSort(arr []int) []int {

	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {

			if arr[j] > arr[j+1] {
				n := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = n
			}
		}
	}

	return arr
}

type Node struct {
	value string
	next  *Node
}

type SinglyLinkedList struct {
	length int
	head   *Node
}

func (list SinglyLinkedList) GetLength(currentNode *Node) int {

	if list.head == nil {
		return list.length
	}

	if list.head.next != nil {
		return 1 + list.GetLength(list.head.next)
	}

	list.length++
	return list.length
}

func (list SinglyLinkedList) Prepend(item Node) {

}

// func (data SinglyLinkedList) Get(index int) string {
//
// 	if index < 0 {
// 		fmt.Println("index < 0")
// 		return ""
// 	}
//
// 	if index > 0 {
// 		fmt.Printf("index = %d; node = %s\n", index, data.currentNode.value)
// 		index--
// 		data.currentNode = *data.currentNode.next
// 		return data.Get(index)
// 	}
//
// 	if index == 0 {
// 		return data.currentNode.value
// 	}
//
// 	return ""
// }

// func (data SinglyLinkedList) Append(node Node) SinglyLinkedList {
//
// 	currentNode := data.head
//
// 	if currentNode.next != nil {
// 		currentNode.next = &node
// 	}
//
// 	l := func(l SinglyLinkedList){
// 	}
//
// 	data.length++
//
// 	return data
// }

func main() {
	// data := []string{"A", "B", "C", "D", "E"}
	// linkedList := SinglyLinkedList{}
	//
	// for _, value := range data {
	// 	node := node{value: value}
	// 	linkedList = linkedList.Prepend(node)
	// }
	//
	// got := linkedList.GetLength(linkedList.head)
	//
	// want := len(data)
	//
	// fmt.Printf("got '%d' want '%d'", got, want)
}
