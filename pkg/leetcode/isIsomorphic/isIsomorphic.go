package main


func isIsomorphic(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}

	hash := make(map[int8]int8)
	hash_check := make(map[int8]int8)
	for i, num := range s {
		num2 := t[i]
		if hash[int8(num)] == 0 {
			if hash_check[int8(num2)] == 0 {
				hash_check[int8(num2)]++
				hash[int8(num)] = int8(num2)
			} else {
				return false
			}
		} else if hash[int8(num)] != int8(num2) {
			return false
		}
	}
	return true
}
