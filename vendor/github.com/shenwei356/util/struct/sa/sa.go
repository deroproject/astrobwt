package sa

import "sort"

// SuffixArray returns the suffix array of s
func SuffixArray(s []byte) []int {
	n := len(s)
	suffixMap := make(map[string]int, n)
	for i := 0; i < n; i++ {
		suffixMap[string(s[i:])] = i
	}
	suffixes := make([]string, n)
	i := 0
	for suffix := range suffixMap {
		suffixes[i] = suffix
		i++
	}
	indice := make([]int, n)
	i = 0
	sort.Strings(suffixes)
	for _, suffix := range suffixes {
		indice[i] = suffixMap[suffix]
		i++
	}
	return indice
}
