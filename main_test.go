package main

import (
	"fmt"
	"math/rand"
	"slices"
	"testing"
	"time"
)

var sorted = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
var unsorted = []int{13, 11, 17, 3, 1, 15, 10, 19, 6, 9, 2, 8, 20, 5, 14, 18, 7, 12, 16, 4}

var target = 12

func TestLinearSearch(t *testing.T) {

	got := LinearSearch(sorted, target)
	want := true

	if !got {
		t.Errorf("got '%t' want '%t'", got, want)
	}
}

func TestBinarySearch(t *testing.T) {

	got := BinarySearch(sorted, target)
	want := true

	if got != want {
		t.Errorf("got '%t' want '%t'", got, want)
	}
}

func TestCrystalBall(t *testing.T) {

	t.Run("random", func(t *testing.T) {
		size := 100
		rand.NewSource(time.Now().UnixNano())
		true_index := rand.Intn(size)
		items := make([]bool, size)
		for i := true_index; i < len(items); i++ {
			items[i] = true
		}

		got := CrystalBall(items)
		want := true_index

		if got != want {
			t.Errorf("got '%d' want '%d'", got, want)
		}
	})

	t.Run("fixed", func(t *testing.T) {
		true_index := 99
		size := 100
		items := make([]bool, size)
		for i := true_index; i < len(items); i++ {
			items[i] = true
		}

		got := CrystalBall(items)
		want := true_index

		if got != want {
			t.Errorf("got '%d' want '%d'", got, want)
		}
	})
}

func TestBubbleSort(t *testing.T) {

	got := BubbleSort(unsorted)
	want := sorted

	if !slices.Equal(got, want) {
		t.Errorf("got %v want %v", got, want)
	}

}

// LinkedSinglyLinkedList
var (
	data = []string{"A", "B", "C", "D", "E"}
)

var list = SinglyLinkedList{}
var currentListLength int

func TestGetLength(t *testing.T) {

	t.Run("no head", func(t *testing.T) {
		list := SinglyLinkedList{}

		got := list.length
		want := currentListLength

		if got != want {
			t.Errorf("got '%d' want '%d'", got, want)
		}
	})

	t.Run("1 node", func(t *testing.T) {

		head := &Node{value: data[0]}
		tail := &Node{}
		list = SinglyLinkedList{1, head, tail}

		currentListLength += 1

		got := list.length
		want := currentListLength

		if got != want {
			t.Errorf("got '%d' want '%d'", got, want)
		}

	})
}

func TestPrepend(t *testing.T) {

	t.Run("prepend", func(t *testing.T) {
		for _, data := range data {
			list.Prepend(&Node{value: data})
		}
		currentListLength += len(data)

		got := list.length
		want := currentListLength

		if got != want {
			t.Errorf("got '%d' want '%d'", got, want)
		}
	})

	t.Run("append", func(t *testing.T) {
		for _, data := range data {
			list.Append(&Node{value: data})
		}
		currentListLength += len(data)

		got := list.length
		want := currentListLength

		if got != want {
			t.Errorf("got '%d' want '%d'", got, want)
		}
	})

	t.Run("insertAt", func(t *testing.T) {
		list.InsertAt(1, &Node{value: data[0]})
		currentListLength += 1

		got := list.length
		want := currentListLength

		if got != want {
			t.Errorf("got '%d' want '%d'", got, want)
		}
	})
}

func TestRemove(t *testing.T) {

	t.Run("remove head", func(t *testing.T) {

		repeatCount := 5
		for i := 0; i < repeatCount; i++ {
			list.RemoveHead()
		}
		currentListLength -= repeatCount

		got := list.length
		want := currentListLength

		if got != want {
			t.Errorf("got '%d' want '%d'", got, want)
		}
	})

	t.Run("remove tail", func(t *testing.T) {

		repeatCount := 2
		for i := 0; i < repeatCount; i++ {
			list.RemoveTail()
		}
		currentListLength -= repeatCount

		got := list.length
		want := currentListLength

		if got != want {
			t.Errorf("got '%d' want '%d'", got, want)
		}
	})

	t.Run("removeAt", func(t *testing.T) {
		list.RemoveAt(1)
		currentListLength -= 1

		got := list.length
		want := currentListLength

		if got != want {
			t.Errorf("got '%d' want '%d'", got, want)
		}
	})

	list.Display()
}

func TestGet(t *testing.T) {
	got := list.Get(2)
	want := "B"

	if got.value != want {
		t.Errorf("got %q want %q", got.value, want)
	}
}

var queue = SinglyLinkedList{length: 0}

func TestEnqueue(t *testing.T) {
	for _, data := range data {
		queue.Enqueue(&Node{value: data})
	}

	queue.Display()

	got := queue.length
	want := 5

	if got != want {
		t.Errorf("got '%d' want '%d'", got, want)
	}
}

func TestDeque(t *testing.T) {
	for range data {
		dequed := queue.Deque()
		fmt.Printf("dequed: %s\n", dequed.value)
		queue.Display()
	}

	got := queue.length
	want := 0
	if got != want {
		t.Errorf("got '%d' want '%d'", got, want)
	}
}

func TestPeek(t *testing.T) {

	peekVal := queue.Peek()

	if peekVal != nil {
		fmt.Println(peekVal.value)
	}
}
