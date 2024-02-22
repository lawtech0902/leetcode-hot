/*
 * Author: robin-luo
 * Created time: 2024-02-20 15:02:16
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-20 15:50:15
 */

package solution

import "slices"

func openLock(deadends []string, target string) int {
	visited := make(map[string]bool)
	q := []string{"0000"}
	var level int
	for len(q) > 0 {
		qLen := len(q)
		for i := 0; i < qLen; i++ {
			lock := q[0]
			q = q[1:]
			if slices.Contains(deadends, lock) {
				continue
			}

			if lock == target {
				return level
			}

			for j := 0; j < 4; j++ {
				next := getNextOrPrevLock(lock[:j], lock[j+1:], lock[j], false)
				if !visited[next] {
					q = append(q, next)
					visited[next] = true
				}

				prev := getNextOrPrevLock(lock[:j], lock[j+1:], lock[j], true)
				if !visited[prev] {
					q = append(q, prev)
					visited[prev] = true
				}
			}
		}

		level++
	}

	return -1
}

func getNextOrPrevLock(prefix, postfix string, ch byte, prev bool) string {
	if prev {
		if ch == '0' {
			ch = '9'
		} else {
			ch--
		}
	} else {
		if ch == '9' {
			ch = '0'
		} else {
			ch++
		}
	}

	return prefix + string(ch) + postfix
}
