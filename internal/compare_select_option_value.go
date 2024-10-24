package internal

func CompareSelectOptionValue(a, b string) string {
	if a == b {
		return "true"
	}
	return "false"
}
