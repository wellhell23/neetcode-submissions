func twoSum(numbers []int, target int) []int {
	res:= []int{}
	start,end := 0, len(numbers)-1
	
	for start < end {
		sum := numbers[start]+numbers[end] 
		if sum < target {
			start++
			continue
		}
		if sum > target {
			end--
			continue 
		}
		return []int{start+1, end+1}

	}
	return res
}
