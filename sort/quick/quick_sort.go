package quick

func Sort3Way(arr []int) {
	quickSort3WayRec(arr, 0, len(arr)-1)
}

func quickSort3WayRec(arr []int, s, e int) {
	if s < e {
		i, j := partitionDutchFlag(arr, s, e)
		quickSortRec(arr, s, i)
		quickSortRec(arr, j, e)
	}
}

func partitionDutchFlag(arr []int, s, e int) (int, int) {
	p := arr[e]
	i := s
	mid := s
	j := e - 1
	for mid <= j {
		if arr[mid] < p {
			arr[i], arr[mid] = arr[mid], arr[i]
			i++
			mid++
		} else if arr[mid] == p {
			mid++
		} else {
			arr[j], arr[mid] = arr[mid], arr[j]
			j--
		}
	}

	arr[mid], arr[e] = arr[e], arr[mid]
	return i - 1, mid + 1
}

func Sort(arr []int) {
	quickSortRec(arr, 0, len(arr)-1)
}

func quickSortRec(arr []int, s, e int) {
	if s < e {
		p := partition(arr, s, e)
		quickSortRec(arr, s, p-1)
		quickSortRec(arr, p+1, e)
	}
}

func partition(arr []int, s, e int) int {
	p := arr[e]
	i := s
	j := e - 1
	for i <= j {
		if arr[j] < p {
			arr[j], arr[i] = arr[i], arr[j]
			i++
		} else {
			j--
		}
	}
	j++
	arr[j], arr[e] = arr[e], arr[j]
	return j
}
