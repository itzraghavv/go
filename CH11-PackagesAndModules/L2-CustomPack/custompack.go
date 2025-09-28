package l2custompack

func Reverse(s string) string {
	result := ""
	for _, v := range s {
		result = string(v) + result
	}
	return result
}
