/*
__author__ = 'robin-luo'
__date__ = '2025/06/13 10:51'
*/

package solution

func groupAnagrams(strs []string) [][]string {
	hash := make(map[[26]int][]string)
	for _, str := range strs {
		cnt := [26]int{}
		for _, ch := range str {
			cnt[ch-'a']++
		}

		hash[cnt] = append(hash[cnt], str)
	}

	res := make([][]string, 0)
	for _, v := range hash {
		res = append(res, v)
	}

	return res
}
