package fenwick

import (
	"fmt"
	"strconv"
)

func buildBit(arr []int) []int {
	l := len(arr) + 1
	bit := make([]int, l)
	for i, e := range arr {
		i++
		fmt.Println("i: ", i, strconv.FormatInt(int64(i), 2))
		for i < l {
			fmt.Println(i, strconv.FormatInt(int64(i), 2))
			bit[i] += e
			i = i + (-i & i)

		}
	}
	return bit
}

func prefixSum(bit []int, i int) int {
	sum := 0
	i++
	fmt.Println("sum: ", i, strconv.FormatInt(int64(i), 2))
	for i > 0 {
		sum += bit[i]
		fmt.Println(i, strconv.FormatInt(int64(i), 2))
		i = i - (i & -i)
	}
	return sum
}

func updateBit(arr, bit []int, i, newval int) {
	diff := newval - arr[i]
	i++
	l := len(bit)
	for i < l {
		bit[i] += diff
		i = i + (-i & i)
	}

}

// func main() {
// 	arr := []int{1, 2, 3, 4, 5, 6, 7}
// 	bit := buildBit(arr)
// 	for i := 0; i < len(arr); i++ {
// 		prefixSum(bit, i)
// 	}
// 	// updateBit(arr, bit, 0, 0)
// 	// for i := 0; i < len(arr); i++ {
// 	// 	fmt.Println(prefixSum(bit, i))
// 	// }

// 	// updateBit(arr, bit, 3, 0)
// 	// for i := 0; i < len(arr); i++ {
// 	// 	fmt.Println(prefixSum(bit, i))
// 	// }

// }
