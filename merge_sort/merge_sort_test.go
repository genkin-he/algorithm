package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	list := []int{5}
	MergeSort(list, 0, len(list))
	if reflect.DeepEqual(list, []int{5}) {
		fmt.Printf("MergeSort Test list succuess\n")
	} else {
		t.Logf("MergeSort Test list Fail:")
	}

	list2 := []int{5, 3}
	MergeSort(list2, 0, len(list2))
	if reflect.DeepEqual(list2, []int{3, 5}) {
		fmt.Printf("MergeSort Test list2 succuess\n")
	} else {
		t.Logf("MergeSort Test list2 Fail:")
	}

	list3 := []int{5, 3, 3, 4, 5, 6, 8, 9, 54, 32, 5, 8}
	MergeSort(list3, 0, len(list3))
	if reflect.DeepEqual(list3, []int{3, 3, 4, 5, 5, 5, 6, 8, 8, 9, 32, 54}) {
		fmt.Printf("MergeSort Test list3 succuess\n")
	} else {
		t.Logf("MergeSort Test list3 Fail:")
	}
}

func TestMergeSort2(t *testing.T) {
	list := []int{5}
	MergeSort2(list, 0, len(list))
	if reflect.DeepEqual(list, []int{5}) {
		fmt.Printf("MergeSort2 Test list succuess\n")
	} else {
		t.Logf("MergeSort2 Test list Fail:")
	}

	list2 := []int{5, 3}
	MergeSort2(list2, 0, len(list2))
	if reflect.DeepEqual(list2, []int{3, 5}) {
		fmt.Printf("MergeSort2 Test list2 succuess\n")
	} else {
		fmt.Printf("%v", list2)
		t.Logf("MergeSort2 Test list2 Fail:")
	}

	list3 := []int{5, 3, 3, 4, 5, 6, 8, 9, 54, 32, 5, 8}
	MergeSort2(list3, 0, len(list3))
	if reflect.DeepEqual(list3, []int{3, 3, 4, 5, 5, 5, 6, 8, 8, 9, 32, 54}) {
		fmt.Printf("MergeSort2 Test list3 succuess\n")
	} else {
		fmt.Printf("%v", list3)
		t.Logf("MergeSort2 Test list3 Fail:")
	}
}
