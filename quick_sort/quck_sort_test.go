package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	arry := []int{5}
	QuickSort(arry, 0, len(arry)-1)
	if reflect.DeepEqual(arry, []int{5}) {
		fmt.Printf("QuickSort Test arry succuess\n")
	} else {
		t.Logf("QuickSort Test arry Fail:")
	}

	arry2 := []int{5, 3}
	QuickSort(arry2, 0, len(arry2)-1)
	if reflect.DeepEqual(arry2, []int{3, 5}) {
		fmt.Printf("QuickSort Test arry2 succuess\n")
	} else {
		t.Logf("QuickSort Test arry2 Fail:")
	}

	arry3 := []int{5, 3, 3, 4, 5, 6, 8, 9, 54, 32, 5, 8}
	QuickSort(arry3, 0, len(arry3)-1)
	if reflect.DeepEqual(arry3, []int{3, 3, 4, 5, 5, 5, 6, 8, 8, 9, 32, 54}) {
		fmt.Printf("QuickSort Test arry3 succuess\n")
	} else {
		t.Logf("QuickSort Test arry3 Fail:")
	}
}
