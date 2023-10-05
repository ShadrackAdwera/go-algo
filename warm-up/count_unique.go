package warmup

func countUniqueValues(nums []int) int {
	if(len(nums)==0) {
		return 0
	}
	i := 0
	// use two "pointers"
	// compare the element at "pointer" i if it's not equal to element at "pointer" j, move i forward, put the element at j where i is 
	// j will automatically be moed forward by the array
	for j := 0; j < len(nums); j++ {
		if(nums[i] != nums[j]) {
			i++
			nums[i] = nums[j]
		}
	}

	return i + 1
 }