package bubble

func Sort(arr []int) {
	l := len(arr)
	for i := 0; i < l-1; i++ {
		flag := false
		for j := 0; j < l-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j+1], arr[j] = arr[j], arr[j+1]
				flag = true
			}
		}
		if !flag {
			return
		}
	}
}
