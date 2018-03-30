package slices

// StringsContain returns true if the specified slice of strings
// contain `str`.
func StringsContain(slice []string, str string) bool {
	for _, l := range slice {
		if l == str {
			return true
		}
	}
	return false
}
