package radix

func Sort(arr []int) {
	if len(arr) == 0 {
		return
	}
	max := arr[0]

	for _, v := range arr {
		if v > max {
			max = v
		}
	}

	div := 1
	for max/div != 0 {
		countSort(arr, div)
		div *= 10

	}
}

func countSort(arr []int, exp int) {
	count := make([]int, 10)

	for _, e := range arr {
		count[(e/exp)%10]++
	}
	for i := 1; i < 10; i++ {
		count[i] += count[i-1]
	}
	output := make([]int, len(arr))
	for j := len(arr) - 1; j >= 0; j-- {
		ind := (arr[j] / exp) % 10
		output[count[ind]-1] = arr[j]
		count[ind]--
	}
	for i, e := range output {
		arr[i] = e
	}
}
