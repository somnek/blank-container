package main

func contains[T comparable](l []T, e T) bool {
	for _, x := range l {
		if x == e {
			return true
		}
	}
	return false
}
