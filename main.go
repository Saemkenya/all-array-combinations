package main

import (
	"fmt"
	"sort"
)

// interface to hold the doCalc func signature
type Calculator interface {
	doCalc() []int
}

// function that utilizes the Calculator interface
func Process(calc Calculator) []int {
	return calc.doCalc()
}

func main() {
	arr := []int{45, 32, 21, 35, 67, 34, 23, 31, 78}
	comb := Combination{
		_arr:    arr,
		_target: 2,
	}
	// _arr := comb.doCalc()
	_arr := Process(comb)
	getnum := GetNumber{
		_arr:    _arr,
		_target: 60,
	}
	// _arr1 := getnum.doCalc()
	_arr1 := Process(getnum)
	fmt.Println(arr)
	fmt.Println(_arr)
	fmt.Println(_arr1[0])
}

// struct to hold the logic of producing all possible combinations
// from a given array
type Combination struct {
	_arr    []int
	_target int
}

// Compute all possible combinations from array
func (cn Combination) doCalc() []int {
	length := uint(len(cn._arr))
	var result []int

	// Go through all possible combinations of objects
	// from 1 (only first object in subset) to 2^length (all objects in subset)
	for subsetBits := 1; subsetBits < (1 << length); subsetBits++ {
		var subset []int

		// go through all elements of the array in every iteration
		for object := uint(0); object < length; object++ {
			// checks if object is contained in subset
			// by checking if bit 'object' is set in subsetBits
			if (subsetBits>>object)&1 == 1 {
				// add object to subset
				subset = append(subset, cn._arr[object])
			}
		}

		// cn._target specifies the length of array to be used in
		// preceding calculations discarding the rest of the possibilities
		// append the sum of each two combinations to result(if cn._target == 2)
		if len(subset) == cn._target {
			result = append(result, subset[0]+subset[1])
		}
	}
	return result
}

// struct to hold the logic of getting the exact sum that is
// the closest to the target but still less than the given target
// from the sum of all the possibilities
type GetNumber struct {
	_arr    []int
	_target int
}

// Get the number combinations next to but less than the target int
func (gn GetNumber) doCalc() []int {
	result := []int{0}
	sort.Ints(gn._arr)
	for _, val := range gn._arr {
		if val <= gn._target {
			result[0] = val
		}
	}
	return result
}
