package main 

import (
	"binarysearchrecursive/algo"
	"fmt"
)

func main() {
	array:=[]int{2, 3, 5, 7, 11, 13}
	length:=len(array)
	
	fmt.Println(algo.BinarySearchRecursive(array,7,0,length))
	
	fmt.Println(algo.BinarySearchRecursive(array,1,0,length))
}
