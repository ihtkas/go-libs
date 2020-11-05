package merge

func Sort(arr []int) {
	mergeSortRec(arr, 0, len(arr)-1)
}

func mergeSortRec(arr []int, s, e int) {
	if s < e {
		m := (s + e) / 2
		mergeSortRec(arr, s, m)
		mergeSortRec(arr, m+1, e)
		merge(arr, s, e)
	}
}

func merge(arr []int, s, e int) {
	clone := make([]int, e-s+1)
	ind := 0
	i := s
	m := (s + e) / 2
	j := m + 1
	// fmt.Println(s, e, m)
	for i <= m && j <= e {
		if arr[i] < arr[j] {
			clone[ind] = arr[i]
			i++
		} else {
			clone[ind] = arr[j]
			j++
		}
		ind++
	}
	for i <= m {
		clone[ind] = arr[i]
		ind++
		i++
	}

	for j <= e {
		clone[ind] = arr[j]
		ind++
		j++
	}
	for _, e := range clone {
		arr[s] = e
		s++
	}

}
