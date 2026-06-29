func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
  
    srune, trune := make(map[rune]int), make(map[rune]int)
    for i, ch := range s {
        srune[ch]++
        trune[rune(t[i])]++
    }

    for k,v := range srune {
        if trune[k] != v {
            return false
        }
    }
    return true

}
