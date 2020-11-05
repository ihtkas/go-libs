package heap

func Sort(arr []int) {
	l := len(arr)
	for i := (l - 2) / 2; i >= 0; i-- {
		heapify(arr, i, l)
	}
	// fmt.Println("heapify", arr)

	for i := l - 1; i > 0; i-- {
		arr[i], arr[0] = arr[0], arr[i]
		heapify(arr, 0, i)

	}
}

func heapify(arr []int, i, e int) {
	for i < e {
		max := i
		c1 := 2*i + 1
		c2 := c1 + 1
		if c1 < e && arr[i] < arr[c1] {
			max = c1
		}

		if c2 < e && arr[max] < arr[c2] {
			max = c2
		}
		// fmt.Println(max, i, c1, c2, arr)
		if max != i {
			arr[max], arr[i] = arr[i], arr[max]
			i = max
		} else {
			return
		}
	}
}
