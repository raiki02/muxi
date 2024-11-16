package main

import "fmt"

func main() {
	var (
		l   int
		num int
		sum int
		avg int
	)
	fmt.Scanln(&l)
	arr := make([]int, 0)
	for i := 0; i < l; i++ {
		fmt.Scanln(&num)
		arr = append(arr, num)
	}

	//-+操作 尽量让数相近
	//回溯 开 n-2个子树 是外sigma，n-1个子节点内sigma
	//求平均数，向平均数靠
	for i := 0; i < l; i++ {
		sum += arr[i]
	}
	avg = sum / l
	for 

}
