func isPalindrome(s string) bool {
	character := []rune(s)
	end:= len(s)-1;
	for i:=0; i < end; i,end = i+1,end-1 {
		for i < end && !unicode.IsLetter(character[i]) && !unicode.IsDigit(character[i]){
			i++
		}
		for i < end && !unicode.IsLetter(character[end]) && !unicode.IsDigit(character[end]){
			end--
		}
		if unicode.ToLower(character[i]) != unicode.ToLower(character[end]){
			return false
		}
	}
	return true
}
