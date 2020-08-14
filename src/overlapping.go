package overlapping

func IsNonOverlap(str string) bool {
	ab := "AB"
	ba := "BA"
	containsAB := false
	containsBA := false

	for i := 0; i < len(str); i++ {
		if (containsAB && containsBA) || i == len(str)-1 {
			break
		}

		if !containsAB && str[i] == ab[0] && str[i+1] == ab[1] {
			containsAB = true
			i++
			continue
		}

		if !containsBA && str[i] == ba[0] && str[i+1] == ba[1] {
			containsBA = true
			i++
		}
	}

	return containsAB && containsBA
}
