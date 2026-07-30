func isAnagram(s string, t string) bool {
	var counter int32 
	for _, r := range s {
		counter += r
	}
	for _, r := range t {
		counter -= r
	}
	return counter == 0
}
