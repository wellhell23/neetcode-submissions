func topKFrequent(nums []int, k int) []int {
    freq:= make(map[int]int)
    for _,n:= range nums {
        freq[n]++
    }
    arr := make([][2]int, 0, len(freq))
    for f, v := range freq {
        arr = append(arr, [2]int{f,v})
    }
    sort.Slice(arr, func(i, j int) bool {
        return arr[i][1] > arr[j][1]
    }) 
    res := make([]int, k)
    for i:=0 ;i< k;i++ {
        res[i]= arr[i][0]
    }
    return res
}