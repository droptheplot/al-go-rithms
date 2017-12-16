package main

func bubbleSort(a []int) []int {
	if len(a) <= 1 {
		return a
	}

	var isSorted = false

	for !isSorted {
		isSorted = true

		for i := range a {
			if i+1 >= len(a) {
				continue
			}

			if a[i] > a[i+1] {
				a[i], a[i+1] = a[i+1], a[i]
				isSorted = false
			}
		}
	}

	return a
}
