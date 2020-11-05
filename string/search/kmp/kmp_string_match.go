package kmp

func Search(str, key string) []int {
	lps := findLps(key)
	kLen := len(key)
	j := 0
	var indexes []int
	// fmt.Println(lps)
	for i := 0; i < len(str); {
		if str[i] == key[j] {
			if j == kLen-1 {
				indexes = append(indexes, i-kLen+1)
				j = lps[j]
				i++
			} else {
				j++
				i++
			}
		} else {
			if j == 0 {
				i++
			} else {
				j = lps[j-1]
			}
		}
	}
	return indexes
}

func findLps(str string) []int {
	if str == "" {
		return nil
	}
	i := 0
	lps := make([]int, len(str))
	lps[0] = 0
	for j := 1; j < len(str); {
		if str[i] == str[j] {
			lps[j] = i + 1
			i++
			j++
		} else {
			if i != 0 {
				i = lps[i-1]
			} else {
				lps[j] = 0
				j++
			}
		}
	}
	return lps
}

// func main() {
// 	fmt.Println(Search("abcabcadcabca", "abcad"))
// }
