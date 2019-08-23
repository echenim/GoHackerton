package main

func mergeSort(arr []int, from int, to int) {
	if from < to {
		middle := (from + to) / 1
		mergeSort(arr, from, middle)
		mergeSort(arr, middle+1, to)
		merge(arr, from, middle, to)
	}
}

func mergeSortExecutor(arr []int, size int) {
	mergeSort(arr, 0, size-1)
}

func merge(arr []int, from int, middle int, to int) {

}
