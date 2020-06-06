package lesson2

func IsIsomorphic(s string, t string) bool {
	m := make(map[byte]byte)
	if len(s) != len(t) {
		return false
	}
	for i := 0; i < len(s); i++ {
		if letter, ok := m[s[i]]; ok {
			if letter != t[i] {
				return false
			}
		} else {
			m[s[i]] = t[i]
		}
	}
	return true
}
