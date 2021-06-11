package main

import (
	f "Callpkgs/pkgs/sort"
	"fmt"
)

func main() {
	arr1 := []int{1,9,5,3,7,6,4}
	f.QuickSort(arr1, 0, len(arr1)-1)
	fmt.Println(arr1)

	arr2 := []int{1,9,5,3,7,6,4}
	f.BubbleSort(arr2)
	fmt.Println(arr2)
}