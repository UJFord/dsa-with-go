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

func TestLinkedList(t *testing.T) {
	// list := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I"}
	data := T{value: "A", nxt: &T{value: "B"}}

	t.Run("length", func(t *testing.T) {

		got := data.length()
		want := 1

		if got != want {
			t.Errorf("got '%d' want '%d'", got, want)
		}
	})

	t.Run("append", func(t *testing.T) {
		new := T{value: "A", nxt: &T{value: "B"}}
		fmt.Println(new)
	})
}
