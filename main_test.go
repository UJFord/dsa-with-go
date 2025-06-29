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

var list = SinglyLinkedList{}
var data = []string{"A", "B", "C", "D", "E"}

func TestPrepend(t *testing.T) {

	t.Run("prepend", func(t *testing.T) {
		for _, data := range data {
			list.Prepend(&Node{value: data})
		}

		got := list.Display()
		want := "E -> D -> C -> B -> A"

		if got != want {
			t.Errorf("got '%q' want '%q'", got, want)
		}
	})

	t.Run("append", func(t *testing.T) {
		for _, data := range data {
			list.Append(&Node{value: data})
		}

		got := list.Display()
		want := "E -> D -> C -> B -> A -> A -> B -> C -> D -> E"

		if got != want {
			t.Errorf("got '%q' want '%q'", got, want)
		}
	})

	t.Run("insertAt", func(t *testing.T) {
		list.InsertAt(1, &Node{value: data[0]})

		got := list.Display()
		want := "E -> A -> D -> C -> B -> A -> A -> B -> C -> D -> E"

		if got != want {
			t.Errorf("got '%q' want '%q'", got, want)
		}
	})
}

func TestRemove(t *testing.T) {

	t.Run("remove head", func(t *testing.T) {
		repeatCount := 5
		for range repeatCount {
			list.RemoveHead()
		}

		got := list.Display()
		want := "A -> A -> B -> C -> D -> E"

		if got != want {
			t.Errorf("got '%q' want '%q'", got, want)
		}
	})

	t.Run("remove tail", func(t *testing.T) {
		repeatCount := 2
		for range repeatCount {
			list.RemoveTail()
		}

		got := list.Display()
		want := "A -> A -> B -> C"

		if got != want {
			t.Errorf("got '%q' want '%q'", got, want)
		}
	})

	t.Run("removeAt", func(t *testing.T) {
		list.RemoveAt(1)

		got := list.Display()
		want := "A -> B -> C"

		if got != want {
			t.Errorf("got '%q' want '%q'", got, want)
		}
	})
}

func TestGet(t *testing.T) {
	got := list.Get(2)
	want := "C"

	if got.value != want {
		t.Errorf("got %q want %q", got.value, want)
	}
}

var queue = Queue{}

func TestEnqueue(t *testing.T) {
	for _, data := range data {
		queue.Enqueue(&Node{value: data})
	}

	got := queue.Display()
	want := "A -> B -> C -> D -> E"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestDeque(t *testing.T) {
	queue.Deque()

	got := queue.Display()
	want := "B -> C -> D -> E"

	if got != want {
		t.Errorf("got %q want '%q", got, want)
	}
}

func TestQueuePeek(t *testing.T) {

	got := queue.Peek()
	want := "B"

	fmt.Printf("Peeking Queue Result: %v\n\n", got.value)
	if got.value != want {
		t.Errorf("got %v want '%q", got, want)
	}
}

var stack = Stack{}

func TestPush(t *testing.T) {
	for _, data := range data {
		stack.Push(&Node{value: data})
	}

	got := stack.Display()
	want := "A <- B <- C <- D <- E"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestPop(t *testing.T) {

	repeatCount := 3
	for range repeatCount {
		stack.Pop()
	}

	got := stack.Display()
	want := "A <- B"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestStackPeek(t *testing.T) {
	got := stack.Peek()
	want := "B"

	fmt.Printf("Peeking Queue Result: %v\n\n", got.value)
	if got.value != want {
		t.Errorf("got %v want %q", got, want)
	}
}
