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

type T struct {
	prev  *T
	value string
	nxt   *T
}

type LinkedList interface {
	length() int
	insert_at(T, int)
	remove(T) T
	remove_at(int) T
	append(T)
	get(int) T
}

//
// func DataSingly(linked_list T, list []string, index int) (T, int) {
//
// 	linked_list = T{}
// 	linked_list.value = list[index]
//
// 	if linked_list.nxt.value == "" {
// 		return linked_list, index
// 	}
//
// 	index++
//
// 	return DataSingly(linked_list, list, index)
// }

func (itm *T) length() int {

	if itm == nil {
		return 0
	}

	if itm.value == "" {
		return 0
	}

	return 1 + itm.nxt.length()
}
