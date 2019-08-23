package main

func partition(arr []int, from int, to int) int {
	pivot := arr[from]
	i := from
	j := to
	for i < j {
		for arr[i] <= pivot {
			i++
		}
		for arr[j] > pivot {
			j--
		}
		if i < j {
			swap(&arr[i], &arr[j])
		}
	}
	swap(&arr[from], &arr[j])

	return j
}

func swap(a *int, b *int) {
	tmp := *a
	*a = *b
	*b = tmp

}

func quickSort(arr []int, l int, h int) {

	if l < h {
		j := partition(arr, l, h)
		quickSort(arr, l, j)
		quickSort(arr, j+1, h)
	}
}
