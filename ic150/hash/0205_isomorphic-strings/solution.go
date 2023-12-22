/*
__author__ = 'robin-luo'
__date__ = '2023/12/21 11:05'
*/

package solution

func isIsomorphic(s, t string) bool {
	m1 := make(map[byte]byte, len(s))
	m2 := make(map[byte]byte, len(s))
	for i := 0; i < len(s); i++ {
		if m1[s[i]] == 0 && m2[t[i]] == 0 {
			m1[s[i]] = t[i]
			m2[t[i]] = s[i]
		} else if m1[s[i]] != t[i] {
			return false
		}
	}

	return true
}
