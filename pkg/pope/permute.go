package main

// Perm calls f with each permutation of a.
func Permute[T any](a []T, f func([]T)) {
	perm(a, f, 0)
}

// Permute the values at index i to len(a)-1.
func perm[T any](a []T, f func([]T), i int) {
	if i > len(a) {
		f(a)
		return
	}
	perm(a, f, i+1)
	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		perm(a, f, i+1)
		a[i], a[j] = a[j], a[i]
	}
}
