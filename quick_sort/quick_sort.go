package main

import "fmt"

// QuickSort 普通快速排序算法
// 先从数列中取出一个数作为基准数。一般取第一个数
// 分区过程，将比这个数大的数全放到它的右边，小于或等于它的数全放到它的左边
// 再对左右区间重复第二步，直到各区间只有一个数

func QuickSort(arry []int, begin, end int) {
	if begin < end {
		loc, err := partition(arry, begin, end)
		if err != nil {
			fmt.Printf("partition error: %s", err.Error())
		}
		QuickSort(arry, begin, loc-1)
		QuickSort(arry, loc+1, end)
	}
}

func partition(arry []int, begin, end int) (i int, err error) {
	i = begin + 1
	j := end
	for i < j {
		if arry[i] > arry[begin] {
			arry[i], arry[j] = arry[j], arry[i]
			j--
		} else {
			i++
		}
	}
	if arry[i] >= arry[begin] {
		i--
	}
	arry[begin], arry[i] = arry[i], arry[begin]
	return
}
func main() {
	var arry = []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}
	QuickSort(arry, 0, len(arry)-1)
	fmt.Println(arry)
}
