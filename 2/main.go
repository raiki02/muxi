package main

import "fmt"

func main() {
	var (
		l       int
		num     int
		sum     int
		avg     int
		right   int
		left    int
		new_sum int
	)
	fmt.Scanln(&l)
	arr := make([]int, 0)
	for i := 0; i < l; i++ {
		fmt.Scanln(&num)
		arr = append(arr, num)
	}

	//-+操作 尽量让数相近
	//回溯 开 n-2个子树 是外sigma，n-1个子节点内sigma x
	//O(n-1logn-2?)
	//求平均数，向平均数靠
	for i := 0; i < l; i++ {
		sum += arr[i]
	}
	avg = sum / l
	//降序排列//保证相减为正数不考虑负号转正
	for i := 0; i < l; i++ {
		for j := i; j < l-i-1; j++ {
			if arr[j] < arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	//向中间靠
	right = l - 1
	left = 0
	for right > left {
		for arr[right] < avg && arr[left] > avg {
			arr[right]--
			arr[left]++
		}
		left++
		right--
	}

	//嵌套循环求最美
	for i := 0; i < l; i++ {
		for j := i; j < l-1-i; j++ {
			new_sum += arr[j] - arr[j+1]
		}
	}

	//打印
	fmt.Println(new_sum)
}
