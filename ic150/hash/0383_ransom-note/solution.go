/*
__author__ = 'robin-luo'
__date__ = '2023/12/21 10:54'
*/

package solution

func canConstruct(ransomNote string, magazine string) bool {
	hash := [26]int{}
	for _, c := range magazine {
		hash[c-'a']++
	}

	for _, c := range ransomNote {
		if hash[c-'a'] == 0 {
			return false
		}

		hash[c-'a']--
	}

	return true
}
