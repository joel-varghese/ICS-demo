package main

import "fmt"

func main() {
	var i int = 10
	var ai *int = &i
	fmt.Println("%v", ai)
	fmt.Println("%v", i)
	fmt.Println("%v", *ai)
	fmt.Println("%v", &i)
	fmt.Println("%v", &ai)

	var arr = []int{1, 2, 3, 4}
	var ari *[]int = &arr
	for i := 0; i < 4; i++ {
		fmt.Println(arr[i])
		fmt.Println((*ari)[i])
		fmt.Println(ari)
		fmt.Println(&arr)
	}
}
