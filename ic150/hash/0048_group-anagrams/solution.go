/*
__author__ = 'robin-luo'
__date__ = '2023/12/21 14:03'
*/

package solution

func groupAnagrams(strs []string) [][]string {
	m := map[[26]int][]string{}
	for _, str := range strs {
		cnt := [26]int{}
		for _, ch := range str {
			cnt[ch-'a']++
		}

		m[cnt] = append(m[cnt], str)
	}

	res := make([][]string, 0, len(m))
	for _, v := range m {
		res = append(res, v)
	}

	return res
}
