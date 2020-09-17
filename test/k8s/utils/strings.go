package utils

// StrDefault helper function that returns value if not empty
// otherwise returns dflt.
func StrDefault(value, dflt string) string {
	if len(value) == 0 {
		return dflt
	}
	return value
}

// StrEmpty returns true if the given string is empty
func StrEmpty(value string) bool {
	if value == "" {
		return true
	}
	return false
}
