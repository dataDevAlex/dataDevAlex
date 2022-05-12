/*
simple implementation of bubble-sort algorithm with randomly
generated array and printing of all the steps necessary to sort the array
*/

package main

//imports
import (
	"fmt"
	"math/rand"
	"time"
)

//global variable
var nums []int

//main function
func main() {

	nums = randomArray() //assign to nums variable the random array create by random_array function (technically a slice)
	prints()             //print results
	bubbleSort(nums)     //sorting and printing array (slice) nums using bubble_sort function

}

// ---------- functions ---------- //

//create random array function
func randomArray() []int {
	nums := make([]int, 0)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		nums = append(nums, rand.Intn(100))
	}
	return nums
}

//prints statements to make program more readable
func prints() {
	fmt.Println("Unsorted array:")
	fmt.Println(nums)
	fmt.Println()
	fmt.Println("Sorted array:")
}

//bubble sort function
func bubbleSort(nums []int) {
	k := true
	for k {
		k = false
		for i := 0; i < len(nums)-1; i++ {
			if nums[i+1] < nums[i] {
				nums[i+1], nums[i] = nums[i], nums[i+1]
				k = true
			}
		}
		if !k {
			break
		}
		fmt.Println(nums)
	}
}
