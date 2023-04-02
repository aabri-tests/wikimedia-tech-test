package utils

func IsEmpty(s string) bool {
	for _, b := range []byte(s) {
		if b != 0 {
			return false
		}
	}
	return true
}
func InArray(target string, arr []string) bool {
	for _, s := range arr {
		if s == target {
			return true // string is in array
		}
	}
	return false // string is not in array
}
