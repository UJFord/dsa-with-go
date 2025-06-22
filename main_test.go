package main

import (
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

// LinkedList
var (
	data = []string{"A", "B", "C", "D", "E"}
)

func TestGetLength(t *testing.T) {
	linkedList := SinglyLinkedList{}

	got := linkedList.GetLength()
	want := 0

	if got != want {
		t.Errorf("got '%d' want '%d'", got, want)
	}
}
