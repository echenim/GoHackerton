package main

func merger(barr []int, carr []int, b int, c int) []int {
	darr := make([]int, b+c)
	i := 0
	j := 0
	k := 0
	// m := b - 1
	// n := c - 1
	for i < b && j < c {
		if barr[i] < carr[j] {
			darr[k] = barr[i]
			k++
			i++
		} else {
			darr[k] = carr[j]
			k++
			j++
		}
	}

	for i < b {
		darr[k] = barr[i]
		k++
		i++
	}

	for j < c {
		darr[k] = carr[j]
		k++
		j++
	}

	return darr
}
