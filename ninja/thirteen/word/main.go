// Package word has some word related functions in it
package word

import "strings"

// UseCount returns the number of times a word is used in a string
// no need to write an example for this one
// writing a test for this one is a bonus challenge; harder
func UseCount(s string) map[string]int {
	xs := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m
}

// Count counts the number of words separated by 'space' in a string
func Count(s string) int {
	return len(strings.Fields(s))
}
