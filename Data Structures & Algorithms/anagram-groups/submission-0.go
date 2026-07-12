func groupAnagrams(strs []string) [][]string {
	countsMap := make(map[[26]int][]string)

	for _, s := range strs {
		char := []rune(s)
		counts := [26]int{}
		for _ , c :=range char {
			counts[c-'a']++
		}
		countsMap[counts] = append(countsMap[counts], s)
	}
	res:= [][]string{}
	for  _,value := range countsMap {
		res = append(res, value)
	}
	return res
}
