package overlapping

import "regexp"

func IsOverlap(str string) bool {
	return regexp.MustCompile("(^|[^B]|AB)AB").MatchString(str) && regexp.MustCompile("(^|[^A]|BA)BA").MatchString(str)
}
