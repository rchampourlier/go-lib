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

// StringsMatch returns true if the two slices contain the same
// strings, even if they are ordered differently.
func StringsMatch(s1 []string, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}
	for _, s1Item := range s1 {
		if !StringsContain(s2, s1Item) {
			return false
		}
	}
	return true
}
