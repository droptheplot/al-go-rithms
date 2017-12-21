package main

func missingInteger(a []int) int {
	var all, given, min, max int

	for _, e := range a {
		if e < min {
			min = e
		} else if e > max {
			max = e
		}

		given ^= e
	}

	for i := min; i <= max; i++ {
		all ^= i
	}

	return all ^ given
}
