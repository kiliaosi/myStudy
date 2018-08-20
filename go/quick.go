package main

import "fmt"

func main(){
	var arr = []int{1,3,4,5,2,123,4,5,321,3,3,4,5,12,32123,89};
	var diff = quickSort(arr)
	fmt.Println(diff)
	}
func quickSort(arr []int) []int{

	if(len(arr)<=2){

		return arr
	}
	var left  []int;
	var right  []int;
	var middleIndex int= len(arr)/2;
	var middle = arr[middleIndex];
	arr = append(arr[:middleIndex],arr[middleIndex+1:]...)
	var i  int;
	for i = 0;i<len(arr);i++{
		if(arr[i]<middle){
			left = append(left,arr[i])
		}else{
			right = append(right,arr[i])
		}
	}


	//return left\
	left = append(left,middle)
	//return left
	return append(quickSort(left),quickSort(right)...)

	//return left
	}