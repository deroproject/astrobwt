package randutil

import "math/rand"

// Shuffle shuffles a slice of int
func Shuffle(s []int) {
	for i := range s {
		j := rand.Intn(i + 1)
		s[i], s[j] = s[j], s[i]
	}
}
