package main

//Insert function for heap
func Insert(arr []int, size int) {

	i := size
	temp := arr[i]

	for i > 0 && temp > arr[i/2] {
		arr[i] = arr[i/2]
		i = i / 2
	}
	arr[i] = temp

}

//Delete function for heap
//Delete function by deleting the root value and transfer it to the end of the array
func Delete(arr []int, size int) int {

	//x := arr[size]
	val := arr[0]
	arr[0] = arr[size]
	//assign the delete val to the end of the array
	arr[size] = val
	i := 0
	j := i * 2
	for j < size-1 {
		if arr[j+1] > arr[j] {
			j = j + 1
		}

		if size < 3 && j == 0 {
			j = j + 1

			if arr[i] > arr[j] {
				temp := arr[i]
				arr[i] = arr[j]
				arr[j] = temp
			}
		}

		if arr[i] < arr[j] {
			temp := arr[i]
			arr[i] = arr[j]
			arr[j] = temp
			i = j
			j = 2 * j
		}
	}

	return val
}

//Sorted function for heap
//Sorted function by moving the root value and transfer it to the end of the array
func Sorted(arr []int, size int) int {

	//x := arr[size]
	val := arr[0]
	arr[0] = arr[size]
	//assign the delete val to the end of the array
	arr[size] = val
	i := 0
	j := i * 2
	for j < size-1 {
		if arr[j+1] > arr[j] {
			j = j + 1
		}

		if size < 3 && j == 0 {
			j = j + 1

			if arr[i] > arr[j] {
				temp := arr[i]
				arr[i] = arr[j]
				arr[j] = temp
			}
		}

		if arr[i] < arr[j] {
			temp := arr[i]
			arr[i] = arr[j]
			arr[j] = temp
			i = j
			j = 2 * j
		}
	}

	return val
}
