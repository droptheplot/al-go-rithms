package main

func fibonacci(n uint64) uint64 {
	if n == 1 || n == 2 {
		return 1
	}

	var j uint64 = 1
	var k uint64 = 1

	for i := uint64(2); i < n; i++ {
		j, k = j+k, j
	}

	return j
}
