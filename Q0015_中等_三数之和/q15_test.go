package Q0015_中等_三数之和

import (
	"fmt"
	"testing"
)

func TestQuestion15_1(t *testing.T) {
	inputNums := []int{-1, 0, 1, 2, -1, -4}
	// answer := [][]int{{-1, 0, 1}, {-1, -1, 2}}
	res := threeSum(inputNums)
	fmt.Printf("res=%v", res)
}
func TestQuestion15_2(t *testing.T) {
	inputNums := []int{0, 0, 0, 0}
	// answer := [][]int{{-1, 0, 1}, {-1, -1, 2}}
	res := threeSum(inputNums)
	fmt.Printf("res=%v", res)
}
func TestQuestion15_3(t *testing.T) {
	inputNums := []int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6}
	// answer := [][]int{{-1, 0, 1}, {-1, -1, 2}}
	res := threeSum(inputNums)
	fmt.Printf("res=%v", res)
}
