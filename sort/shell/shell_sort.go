package shell

func Sort(arr []int) {
	n := len(arr)
	for i := n >> 1; i >= 1; i = i >> 1 {
		// fmt.Println("i:", i)
		for j := i; j < n; j++ {
			temp := arr[j]
			k := j
			for ; k >= i && arr[k-i] > temp; k = k - i {
				// fmt.Println("Comp:", arr[k-i], arr[k], k, "j:", j)
				arr[k] = arr[k-i]
			}
			arr[k] = temp
		}
	}
}
