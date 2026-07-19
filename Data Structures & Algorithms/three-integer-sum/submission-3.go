func threeSum(nums []int) [][]int {
	res:= [][]int {}
	i, end:= 0,len(nums)-1
	sort.Ints(nums)
	for i < end-1 {
		if nums[i] > 0 {
			return res
		}
		if i>0 && nums[i] == nums[i-1] {
			i++
			continue
		}
		j, k := i+1, end
		for j < k {
			sum:= nums[i]+nums[j] +nums[k]
			if sum == 0 {
				res = append (res, []int{nums[i],nums[j],nums[k]})
				for j < k && nums[j] == nums[j+1] {
					j++
				}
				for j < k && nums[k] == nums[k-1]{
					k--
				}
				j++
				k--
			} else if sum > 0 {
				k--
			} else {
				j++
			}
		}
		i++
	}
	return res
}
