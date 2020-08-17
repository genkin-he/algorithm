package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestInsertSort(t *testing.T) {
	list := []int{5}
	InsertSort(list)
	if reflect.DeepEqual(list, []int{5}) {
		fmt.Printf("InsertSort Test list succuess\n")
	} else {
		t.Logf("InsertSort Test list Fail:")
	}

	list2 := []int{5, 3}
	InsertSort(list2)
	if reflect.DeepEqual(list2, []int{3, 5}) {
		fmt.Printf("InsertSort Test list2 succuess\n")
	} else {
		t.Logf("InsertSort Test list2 Fail:")
	}

	list3 := []int{5, 3, 3, 4, 5, 6, 8, 9, 54, 32, 5, 8}
	InsertSort(list3)
	if reflect.DeepEqual(list3, []int{3, 3, 4, 5, 5, 5, 6, 8, 8, 9, 32, 54}) {
		fmt.Printf("InsertSort Test list3 succuess\n")
	} else {
		t.Logf("InsertSort Test list3 Fail:")
	}
}
