package main

import "fmt"

func main() {
	var (
		n   int
		q   int
		ai  int
		sum int //睡几个小时后的清醒值
	)
	arr := make([]int, 0) //记录每个小时的清醒值
	xj := make([]int, 0)  //记录每次询问大于的清醒值
	fmt.Scan(&n, &q, "\n")
	for i := 0; i < n; i++ {
		x := 0
		fmt.Scan(&x)
		arr = append(arr, x)
	}

	for i := 0; i < q; i++ {
		fmt.Scanln(&ai)
		xj = append(xj, ai)
	}

	//将每小时增加的清醒值降序排列
	for i := 0; i < n; i++ {
		for j := i; j < n-i-1; j++ {
			if arr[j] < arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	//对于每一次询问
	//每次从排序后的切片遍历，增加的清醒值记录在sum
	//与询问的xj【】比较
	//大于等于则存在最小值 为下标 + 1
	//若遍历完后sum< xj【】不存在 返回 -1
	for i := 0; i < q; i++ {
		for j := 0; j < n; j++ {
			sum += arr[j]
			if sum >= xj[i] {
				fmt.Println(j + 1)
				break
			} else {
				fmt.Println("-1")
				break
			}
		}
	}

}
