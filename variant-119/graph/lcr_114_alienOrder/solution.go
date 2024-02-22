/*
 * Author: robin-luo
 * Created time: 2024-02-20 20:40:28
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-20 20:54:44
 */

package solution

func alienOrder(words []string) string {
	g := map[byte][]byte{}
	for _, c := range words[0] {
		g[byte(c)] = nil
	}
next:
	for i := 1; i < len(words); i++ {
		s, t := words[i-1], words[i]
		for _, c := range t {
			g[byte(c)] = g[byte(c)]
		}
		for j := 0; j < len(s) && j < len(t); j++ {
			if s[j] != t[j] {
				g[s[j]] = append(g[s[j]], t[j])
				continue next
			}
		}
		if len(s) > len(t) {
			return ""
		}
	}

	const visiting, visited = 1, 2
	order := make([]byte, len(g))
	i := len(g) - 1
	state := map[byte]int{}
	var dfs func(u byte) bool
	dfs = func(u byte) bool {
		state[u] = visiting
		for _, v := range g[u] {
			if state[v] == 0 {
				if !dfs(v) {
					return false
				}
			} else if state[v] == visiting {
				return false
			}
		}
		order[i] = u
		i--
		state[u] = visited
		return true
	}
	for u := range g {
		if state[u] == 0 && !dfs(u) {
			return ""
		}
	}
	return string(order)
}
