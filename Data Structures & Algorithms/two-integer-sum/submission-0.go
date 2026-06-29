func twoSum(nums []int, target int) []int {
       diffMap := make(map[int]int ,len(nums))
    for i, v:= range nums {
        dif := target - v
        if j,found:= diffMap[dif]; found {
            return []int{j, i}
        }
        diffMap[v] = i
    }
    return []int{} 
}
