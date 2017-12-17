package main

func quickselect(a []int) int {
	if len(a) <= 1 {
		return a[0]
	}

	var (
		left []int
		pi   = len(a) / 2
	)

	for i := range a {
		if a[i] < a[pi] {
			left = append(left, a[i])
		}
	}

	if len(left) == 0 {
		return a[pi]
	}

	return quickselect(left)
}
