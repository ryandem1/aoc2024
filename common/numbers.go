package common

func CountDigits(n int) int {
	count := 0
	for n != 0 {
		n /= 10
		count++
	}
	return count
}
