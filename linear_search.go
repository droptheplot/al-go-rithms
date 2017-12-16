package main

func linearSearch(a []int, n int) int {
	if len(a) == 1 {
		return a[0]
	}

	for i := range a {
		if n == a[i] {
			return i
		}
	}

	return -1
}
