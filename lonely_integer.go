package main

func lonelyInteger(a []int) int {
	var k int

	for _, e := range a {
		k ^= e
	}

	return k
}
