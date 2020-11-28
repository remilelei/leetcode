package Q0018_中等_四数之和

import (
	"fmt"
	"testing"
)

func TestQuestion18_1(t *testing.T) {
	inputNums := []int{1, 0, -1, 0, -2, 2}
	res := fourSum(inputNums, 0)
	/*
	   [
	     [-1,  0, 0, 1],
	     [-2, -1, 1, 2],
	     [-2,  0, 0, 2]
	   ]
	*/
	fmt.Printf("res=%v", res)
}

func TestQuestion18_2(t *testing.T) {
	inputNums := []int{-1, -5, -5, -3, 2, 5, 0, 4}
	res := fourSum(inputNums, -7)
	/*
	   [[-5,-5,-1,4],[-5,-3,-1,2]]
	*/
	fmt.Printf("res=%v", res)
}
