package main

import "fmt"

func main(){
	fmt.Println(fibonachi(3))
	fmt.Println(fibonachi(5))
}

func fibonachi(x int) int {
	if x == 0 || x == 1 {
		return 1
	}
	return fibonachi(x-2)+fibonachi(x-1)
}

