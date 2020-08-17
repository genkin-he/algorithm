package main

import "fmt"

// QuickSort 普通快速排序算法
// 先从数列中取出一个数作为基准数。一般取第一个数
// 分区过程，将比这个数大的数全放到它的右边，小于或等于它的数全放到它的左边
// 再对左右区间重复第二步，直到各区间只有一个数

func QuickSort(list []int, begin, end int) {
	if begin < end {
		loc, err := partition(list, begin, end)
		if err != nil {
			fmt.Printf("partition error: %s", err.Error())
		}
		QuickSort(list, begin, loc-1)
		QuickSort(list, loc+1, end)
	}
}

func partition(list []int, begin, end int) (i int, err error) {
	i = begin + 1
	j := end
	for i < j {
		if list[i] > list[begin] {
			list[i], list[j] = list[j], list[i]
			j--
		} else {
			i++
		}
	}
	if list[i] >= list[begin] {
		i--
	}
	list[begin], list[i] = list[i], list[begin]
	return
}
func main() {
	var list = []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}
	QuickSort(list, 0, len(list)-1)
	fmt.Println(list)
}
