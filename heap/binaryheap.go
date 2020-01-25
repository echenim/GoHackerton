package main

//Insert function for heap
func Insert(arr []int, size int) {

	i := size
	temp := arr[i]

	for i > 1 && temp > arr[i/2] {
		arr[i] = arr[i/2]
		i = i / 2
	}
	arr[i] = temp

}
