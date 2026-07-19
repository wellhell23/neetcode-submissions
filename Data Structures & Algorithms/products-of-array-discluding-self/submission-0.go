func productExceptSelf(nums []int) []int {
	prod:= 1
	count:= 0
	res:= make([]int, len(nums))
	for _,v := range nums {
		if v==0 {
			count++
		} else {
			prod*=v
		}
	}
	if count >1 {
		return res
	}
	for i,v := range nums {
		if count> 0 {
			if v==0 {
				res[i] = prod
			} else {
				res[i]=0
			}
		} else {
			res[i] = prod/v
		}
	}
	return res
}
