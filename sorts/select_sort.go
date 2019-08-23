package main

//Sort function for array
func selectSort(arr []int) []int {

	var i int
	var j int
	size := len(arr)
	//fmt.Println(arr)
	for i = 0; i < size-1; i++ {
		compareVal := i
		for j = i + 1; j < size; j++ {
			if arr[j] < arr[compareVal] {
				compareVal = j
			}
		}
		tmp := arr[i]
		arr[i] = arr[compareVal]
		arr[compareVal] = tmp
		//fmt.Println(arr)
	}
	//fmt.Printf("\nfinal sorted result \n %v  \n", arr)
	return arr
}
