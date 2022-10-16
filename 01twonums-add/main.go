package main

import "fmt"
//force method


func addTwoNumber(nums []int, tag int) (int, int ) {
	for index1, num1 := range nums {
		if index1 == len(nums)-1 {
			break
		}
		for index2 := index1+1; index2 < len(nums); index2++ {
			num2 := nums[index2]
			if tag == num1+num2{
				return index1, index2
			}
		}
	}
	return 0, 0
}
func main() {
	nums := []int{2,2,3,6,8}
	tag := 14
	index1,index2 := addTwoNumber(nums, tag)
	fmt.Printf("[%d %d]\n", index1,index2)
}
