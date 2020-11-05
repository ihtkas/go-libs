package rabinkarp

// Search assumes that the string has only lower case letters.
func Search(str string, key string) []int {
	if len(str) < len(key) {
		return nil
	}
	hash, lastMult := findHash(str, 0, len(key))
	kLen := len(key)
	keyHash, _ := findHash(str, 0, kLen)
	var indexes []int
	if hash == keyHash && str[:kLen] == key {
		indexes = []int{0}
	}
	strLen := len(str)
	for i := 1; i <= strLen-kLen; i++ {
		hash = rehash(str, hash, lastMult, i+kLen-1, kLen)
		if hash == keyHash && str[i:i+kLen] == key {
			indexes = append(indexes, i)
		}
	}
	return indexes
}

var p = 31

func findHash(str string, i int, len int) (int, int) {
	val := 0
	j := 1
	limit := i + len
	var prev int
	for ; i < limit; i++ {
		val += (int(str[i]) - 'a') * j
		prev = j
		j *= p
	}
	return val, prev

}

func rehash(str string, oldVal int, lastMult int, i int, len int) int {
	newVal := (oldVal - (int(str[i-len]) - 'a')) / p
	newVal += (int(str[i]) - 'a') * lastMult
	return newVal
}

// func main() {
// 	fmt.Println(Search("abcdabcabca", "abc"))
// }
