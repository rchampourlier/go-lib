package slices

// StringsIntersect returns the intersection of the 2 specificed
// slices.
func StringsIntersect(slice1 []string, slice2 []string) []string {
	r := make([]string, 0)
	for _, str := range slice1 {
		if StringsContain(slice2, str) {
			r = append(r, str)
		}
	}
	return r
}

// StringsSubstract returns the slice made with items from slice1
// without items from slice2.
func StringsSubstract(slice1 []string, slice2 []string) []string {
	r := make([]string, 0)
	for _, str := range slice1 {
		if !StringsContain(slice2, str) {
			r = append(r, str)
		}
	}
	return r
}

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
