package main

func bubbleSort(arr []int) []int {
	var i int
	var j int
	size := len(arr)
	//	fmt.Println(arr)

	for i = 0; i < size-1; i++ {

		for j = 0; j+1 < size-i; j++ {
			if arr[j] > arr[j+1] {
				tmp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = tmp
			}
		}

		//fmt.Println(arr)
	}

	//fmt.Printf("\nfinal sorted result \n %v \n",)
	return arr
}
