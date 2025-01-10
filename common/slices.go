package common


func CopySliceWithIndexExcluded[T any](s []T, i int) []T {
	if i < 0 || i >= len(s) {
		// Return a copy of the original slice if index is out of bounds
		return append([]T(nil), s...)
	}
	// Create a new slice by copying and excluding the element
	return append(append([]T(nil), s[:i]...), s[i+1:]...)
}
