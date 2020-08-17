package main

import "fmt"

//InsertSort 插入排序
//[]表示排好序
//第一轮： [4] 2 9 1 拿待排序的第二个数 2，插入到排好序的数列 [4]
//与排好序的数列 [4] 比较
//第一轮进行中：2 比 4 小，插入到 4 前
//第二轮： [2 4] 9 1 拿待排序的第三个数 9，插入到排好序的数列 [2 4]
//与排好序的数列 [2 4] 比较
//第二轮进行中： 9 比 4 大，不变化
//第三轮： [2 4 9] 1 拿待排序的第四个数 1，插入到排好序的数列 [2 4 9]
//与排好序的数列 [2 4 9] 比较
//第三轮进行中： 1 比 9 小，插入到 9 前
//第三轮进行中： 1 比 4 小，插入到 4 前
//第三轮进行中： 1 比 2 小，插入到 2 前
//结果： [1 2 4 9]
func InsertSort(list []int) {
	n := len(list)
	if n <= 1 {
		return
	}
	for i := 1; i < len(list); i++ {
		deal := list[i]
		j := i - 1
		if deal < list[j] {
			for ; j >= 0 && deal < list[j]; j-- {
				list[j+1] = list[j]
			}
			list[j+1] = deal
		}

	}
}

func main() {
	var list = []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}
	InsertSort(list)
	fmt.Println(list)
}
