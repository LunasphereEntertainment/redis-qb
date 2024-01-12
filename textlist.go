package redis_qb

import "strings"

type TextList []string

func (t TextList) String(quoted bool) string {
	sb := strings.Builder{}
	//
	//if len(t) == 1 {
	//	if quoted {
	//		sb.WriteRune('"')
	//	}
	//	sb.WriteString(t[0])
	//	if quoted {
	//		sb.WriteRune('"')
	//	}
	//} else {

	for i, str := range t {
		if i > 0 {
			sb.WriteString(" | ")
		}

		if quoted {
			sb.WriteRune('"')
		}
		sb.WriteString(str)
		if quoted {
			sb.WriteRune('"')
		}
	}
	//}

	return sb.String()
}
