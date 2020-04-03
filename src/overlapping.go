package overlapping

import "regexp"

func IsOverlap(str string) bool {
	return regexp.MustCompile("(^|[^B]|AB)AB").MatchString(str) && regexp.MustCompile("(^|[^A]|BA)BA").MatchString(str)
}

func IsOverlapQuicker(str string) bool {
	ab := "AB"
	ba := "BA"
	abExist := false
	baExist := false

	for i := 0; i < len(str); i++ {
		if i == len(str)-1 || (baExist && abExist) {
			break
		}

		if !abExist && str[i] == ab[0] && str[i+1] == ab[1] && (i-1 < 0 || str[i-1] != ba[0]) {
			abExist = true
			i++
			continue
		}

		if !baExist && str[i] == ba[0] && str[i+1] == ba[1] && (i-1 < 0 || str[i-1] != ab[0]) {
			baExist = true
			i++
		}
	}

	return abExist && baExist
}
