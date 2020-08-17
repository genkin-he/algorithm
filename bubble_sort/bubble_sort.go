package main

import "fmt"

// BubbleSort 冒泡排序
//现在有一堆乱序的数，比如：5 9 1 6 8 14 6 49 25 4 6 3。
//
//第一轮迭代：从第一个数开始，依次比较相邻的两个数，如果前面一个数比后面一个数大，那么交换位置，直到处理到最后一个数，最后的这个数是最大的。
//
//第二轮迭代：因为最后一个数已经是最大了，现在重复第一轮迭代的操作，但是只处理到倒数第二个数。
//
//第三轮迭代：因为最后一个数已经是最大了，最后第二个数是次大的，现在重复第一轮迭代的操作，但是只处理到倒数第三个数。
//
//第N轮迭代：....
//
//经过交换，最后的结果为：1 3 4 5 6 6 6 8 9 14 25 49，我们可以看到已经排好序了。
//
//因为小的元素会慢慢地浮到顶端，很像碳酸饮料的汽泡，会冒上去，所以这就是冒泡排序取名的来源。

func BubbleSort(list []int) {
	lena := len(list)
	for i := 0; i <= lena-1; i++ {
		for j := 0; j < lena-i-1; j++ {
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
			}
		}
	}
}

func main() {
	var list = []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}
	BubbleSort(list)
	fmt.Println(list)
}
