package main

import "fmt"

//variadic functions
//khoi tao 1 variadic func
func addItem(item int, list ...int) {
	//100,200,300,400 -> int[]{100,200,300,400}
	//[]int{1, 2, 3, 4} -> {int[] {1,2,3,4}}
	list = append(list, item, 11, 22, 33, 44, 55, 66)
	fmt.Println(list)
}

//change slice
func change(list ...int) {
	list[0] = 999
}
func main() {
	addItem(1, 100, 200, 300, 400)

	//pass 1 slice vao 1 variadic func
	var list = []int{1, 2, 3, 4}
	addItem(100, list...)

	change(list...)
	fmt.Println(list)
}
