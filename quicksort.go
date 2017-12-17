package main

func quicksort(a []int) []int {
	if len(a) <= 1 {
		return a
	}

	var (
		left, right []int
		pi          = len(a) / 2
	)

	for i := range a {
		if i == pi {
			continue
		} else if a[i] > a[pi] {
			right = append(right, a[i])
		} else {
			left = append(left, a[i])
		}
	}

	if len(left) > 1 {
		left = quicksort(left)
	}

	if len(right) > 1 {
		right = quicksort(right)
	}

	return append(append(left, a[pi]), right...)
}
