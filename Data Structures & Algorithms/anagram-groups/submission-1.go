func groupAnagrams(strs []string) [][]string {
	countsMap := make(map[[26]int][]string)

	for _, s := range strs {
		counts := [26]int{}
		for i :=range s {
			counts[s[i]-'a']++
		}
		countsMap[counts] = append(countsMap[counts], s)
	}
	res := make([][]string, 0, len(countsMap))
	for  _,value := range countsMap {
		res = append(res, value)
	}
	return res
}
