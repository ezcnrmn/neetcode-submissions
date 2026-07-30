type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	encoded := ""
	del := "?"
	for _, s := range strs {
		encoded += strconv.Itoa(len(s)) + del + s
	}
	return encoded
}

func (s *Solution) Decode(encoded string) []string {
	var i, length int
	var strs []string
	var del byte = '?'

	for i < len(encoded) {
		if length == 0 {
			var lengthStr []byte
			for encoded[i] != del {
				lengthStr = append(lengthStr, encoded[i])
				i++
			}
			length, _ = strconv.Atoi(string(lengthStr))
			if length == 0 {
				strs = append(strs, "")
			}
			i++
		} else {
			var word []byte
			for length > 0 {
				word = append(word, encoded[i])
				i++
				length--
			}
			strs = append(strs, string(word))
		}
	}
	return strs
}
