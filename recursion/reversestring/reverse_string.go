package reversestring

// ReverseString ...
func ReverseString(s []byte) {
	strLen := len(s)
	if strLen == 1 || strLen == 0 {
		return
	}
	start := s[0]
	end := s[strLen-1]
	ReverseString(s[1 : strLen-1])
	s[strLen-1] = start
	s[0] = end
}
