package normalization

import "strings"

func TrimSpace(strs ...*string) {
	for _, str := range strs {
		if str == nil {
			continue
		}
		*str = strings.TrimSpace(*str)
	}
}
