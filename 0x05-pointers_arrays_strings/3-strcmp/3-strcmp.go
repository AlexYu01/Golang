package main

// strcmp will compare two strings based on length and decimal value. Returns less than 0 if: s1 has smaller decimal value, 0 if s1 and s2 match, and greater than 0 if s1 has a greater decimal value.
func strCmp(s1, s2 string) int {
	var i, len1, len2 int

	len1 = len(s1)
	len2 = len(s2)
	for i = 0; i < len1 && i < len2; i++ {
		if s1[i] != s2[i] {
			return int(s1[i]) - int(s2[i])
		}
	}
	if i == len2 {
		return 0
	} else {
		return -15
	}
}
