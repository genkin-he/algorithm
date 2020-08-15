package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	arry := []int{5}
	BubbleSort(arry)
	if reflect.DeepEqual(arry, []int{5}) {
		fmt.Printf("BubbleSort Test arry succuess\n")
	} else {
		t.Logf("BubbleSort Test arry Fail:")
	}

	arry2 := []int{5, 3}
	BubbleSort(arry2)
	if reflect.DeepEqual(arry2, []int{3, 5}) {
		fmt.Printf("BubbleSort Test arry2 succuess\n")
	} else {
		t.Logf("BubbleSort Test arry2 Fail:")
	}

	arry3 := []int{5, 3, 3, 4, 5, 6, 8, 9, 54, 32, 5, 8}
	BubbleSort(arry3)
	if reflect.DeepEqual(arry3, []int{3, 3, 4, 5, 5, 5, 6, 8, 8, 9, 32, 54}) {
		fmt.Printf("BubbleSort Test arry3 succuess\n")
	} else {
		t.Logf("BubbleSort Test arry3 Fail:")
	}
}
