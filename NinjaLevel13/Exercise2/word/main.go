package word

import "strings"

// UseCount returns a map
func UseCount(s string) map[string]int {
	// no need to write an example for this one
	// writing a test for this one is a bonus challenge; harder
	xs := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m
}

// Count calculate the number of characters in a given string
func Count(s string) int {
	// write the code for this func
	xs := strings.Fields(s)
	return len(xs)
}
