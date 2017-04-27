package algo

import (

)

func BinarySearchRecursive(array []int,key int,startIndex , lastIndex int) (result bool){
	
	if(startIndex>lastIndex){
		return false
	}
	
	mid:=(startIndex+lastIndex)/2
	
	if(array[mid]==key){
		return true
	}
	
	if(key>array[mid]){
		return BinarySearchRecursive(array, key, mid+1, lastIndex)
	}else{
		return BinarySearchRecursive(array, key, startIndex, mid-1)
	}
	return
} 