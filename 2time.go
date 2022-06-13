package main

import "fmt"

func main() {

	var num int
	var arr = []int{2, 7, 8, 9, 1, 3, 5, 4, 5, 3, 4}
	for i := 0; i < len(arr); i++ {
		num = 0

		for j := 0; j < len(arr); j++ {
			if i != j {
				if arr[i] == arr[j] {
					num = num + 1
				}
			}
		}

		if num == 0 {
			fmt.Println(arr[i])
		}
	}

}
