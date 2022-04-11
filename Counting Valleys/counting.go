package main

func countingValleys(steps int32, path string) int32 {
	var counts, current int32
	sealevel := true

	for i := int32(0); i < steps; i++ {
		if path[i] == 'U' {
			current++
		} else {
			current--
		}

		if sealevel && current == -1 {
			counts++
		}

		sealevel = current == 0
	}

	return counts
}
