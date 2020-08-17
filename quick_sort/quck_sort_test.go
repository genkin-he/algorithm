package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	list := []int{5}
	QuickSort(list, 0, len(list)-1)
	if reflect.DeepEqual(list, []int{5}) {
		fmt.Printf("QuickSort Test list succuess\n")
	} else {
		t.Logf("QuickSort Test list Fail:")
	}

	list2 := []int{5, 3}
	QuickSort(list2, 0, len(list2)-1)
	if reflect.DeepEqual(list2, []int{3, 5}) {
		fmt.Printf("QuickSort Test list2 succuess\n")
	} else {
		t.Logf("QuickSort Test list2 Fail:")
	}

	list3 := []int{5, 3, 3, 4, 5, 6, 8, 9, 54, 32, 5, 8}
	QuickSort(list3, 0, len(list3)-1)
	if reflect.DeepEqual(list3, []int{3, 3, 4, 5, 5, 5, 6, 8, 8, 9, 32, 54}) {
		fmt.Printf("QuickSort Test list3 succuess\n")
	} else {
		t.Logf("QuickSort Test list3 Fail:")
	}
}
