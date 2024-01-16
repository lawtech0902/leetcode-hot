/*
__author__ = 'robin-luo'
__date__ = '2024/01/16 9:50'
*/

package solution

func groupAnagrams(strs []string) [][]string {
	hash := make(map[[26]int][]string, 0)
	for _, s := range strs {
		cnt := [26]int{}
		for _, ch := range s {
			cnt[ch-'a']++
		}

		hash[cnt] = append(hash[cnt], s)
	}

	res := make([][]string, 0, len(hash))
	for _, v := range hash {
		res = append(res, v)
	}

	return res
}
