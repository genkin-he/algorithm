package main

import "fmt"

//先申请一个辅助数组，长度等于两个有序数组长度的和。
//从两个有序数组的第一位开始，比较两个元素，哪个数组的元素更小，那么该元素添加进辅助数组，然后该数组的元素变更为下一位，继续重复这个操作，直至数组没有元素。
//返回辅助数组。
// 自顶向下递归归并排序
func MergeSort(list []int, begin, end int) {
	if end-begin > 1 {
		mid := begin + (end-begin+1)/2
		MergeSort(list, begin, mid)
		MergeSort(list, mid, end)
		Merge(list, begin, mid, end)
	}
}

// 自底向上的非递归实现：
func MergeSort2(list []int, begin, end int) {
	step := 1
	for end-begin > step {
		for i := begin; i < end; i += step << 1 {
			lo := i
			mid := i + step
			hi := lo + (step << 1)
			if mid > end {
				return
			}

			if end < hi {
				hi = end
			}

			Merge(list, lo, mid, hi)
		}
		step <<= 1
	}

}

func Merge(list []int, begin, mid, end int) {
	lSize := mid - begin
	rSize := end - mid
	newSize := lSize + rSize
	newList := make([]int, 0, newSize)
	r, l := 0, 0
	for r < rSize && l < lSize {
		if list[begin+l] <= list[mid+r] {
			newList = append(newList, list[begin+l])
			l++
		} else {
			newList = append(newList, list[mid+r])
			r++
		}
	}

	newList = append(newList, list[begin+l:mid]...)
	newList = append(newList, list[mid+r:end]...)
	for i := 0; i < newSize; i++ {
		list[begin+i] = newList[i]
	}
}

func main() {
	var list = []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}
	MergeSort(list, 0, len(list))
	fmt.Println(list)
	list = []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}
	MergeSort2(list, 0, len(list))
	fmt.Println(list)
}
