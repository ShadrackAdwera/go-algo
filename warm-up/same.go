package warmup

import (
	"math"
)

/*
 function accepts two arrays and returns true if
 every value in the array has its corresponding value squared in the second array
 the frequency of values must be the same
*/
func same(num1 []int, num2 []int) bool {
	if len(num1) != len(num2) {
		return false
	}

	freqMap := make(map[int]int)

	// Count the frequency of squared values from num1 and num2 in a single map
	for _, num := range num1 {
		squared := int(math.Pow(float64(num), 2))
		freqMap[squared]++
	}
	for _, num := range num2 {
		freqMap[num]--
		if freqMap[num] < 0 {
			return false
		}
	}

	return true
}



// UnOptimized Version
// func Same(num1 []int, num2 []int) bool {
// 	if (len(num1) != len(num2)) {
// 		return false
// 	}

// 	/*
// 	1. create a map 
// 	2. put first array in a hashmap incl the frequency of each element
// 	{
// 		1:2,
// 		2:3,
// 		3:1
// 	}
// 	3. put second array in hashmap incl frequency of each element
// 	{
// 		1: 2,
// 		4: 3,
// 		9: 1
// 	}
// 	4. iterate through first array map (square this value) - compare with second array
// 	*/

// 	mapFirstArry := map[int]int{}
// 	mapSecondArry := map[int]int{}

// 	for _, v := range num1 {
// 		mapFirstArry[v]++
// 	}

// 	for _, v := range num2 {
// 		mapSecondArry[v]++
// 	}

// 	for _, v := range mapFirstArry {
// 		// check if value exists
// 		idx := int(math.Pow(float64(v),2))
// 		if(mapSecondArry[idx] == 0) {
// 			return false
// 		}
// 		// check for the frequency
// 		if(mapFirstArry[v] != mapSecondArry[idx]) {
// 			return false
// 		}
// 	}

// 	return true
// }