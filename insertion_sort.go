package main

func insertionSort(a []int) []int {
	if len(a) == 1 {
		return a
	}

	for i := range a {
		for j := 0; j < i; j++ {
			if a[i] < a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}

	return a
}
