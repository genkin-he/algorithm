package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	list := []int{5}
	BubbleSort(list)
	if reflect.DeepEqual(list, []int{5}) {
		fmt.Printf("BubbleSort Test list succuess\n")
	} else {
		t.Logf("BubbleSort Test list Fail:")
	}

	list2 := []int{5, 3}
	BubbleSort(list2)
	if reflect.DeepEqual(list2, []int{3, 5}) {
		fmt.Printf("BubbleSort Test list2 succuess\n")
	} else {
		t.Logf("BubbleSort Test list2 Fail:")
	}

	list3 := []int{5, 3, 3, 4, 5, 6, 8, 9, 54, 32, 5, 8}
	BubbleSort(list3)
	if reflect.DeepEqual(list3, []int{3, 3, 4, 5, 5, 5, 6, 8, 8, 9, 32, 54}) {
		fmt.Printf("BubbleSort Test list3 succuess\n")
	} else {
		t.Logf("BubbleSort Test list3 Fail:")
	}
}
