package main

func binarySearch(a []int, n int) int {
	if len(a) == 1 {
		return a[0]
	}

	var (
		i     int
		start = 0
		end   = len(a) - 1
	)

	for end-start >= 0 {
		i = ((end - start) / 2) + start

		if n == a[i] {
			return i
		} else if n > a[i] {
			start = i + 1
		} else {
			end = i - 1
		}
	}

	return -1
}
