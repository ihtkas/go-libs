package insertion

func Sort(arr []int) {
	for i := 1; i < len(arr); i++ {
		var temp = arr[i]
		j := i
		for ; j > 0 && arr[j-1] > temp; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = temp
	}
}
