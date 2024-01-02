/*
__author__ = 'robin-luo'
__date__ = '2023/12/25 11:54'
*/

package solution

func simplifyPath(path string) string {
	var stack []string
	i := 0
	res := ""

	for i < len(path) {
		end := i + 1
		for end < len(path) && path[end] != '/' {
			end++
		}

		sub := path[i+1 : end]
		if len(sub) > 0 {
			if sub == ".." {
				if len(stack) != 0 {
					stack = stack[:len(stack)-1]
				}
			} else if sub != "." {
				stack = append(stack, sub)
			}
		}

		i = end
	}

	if len(stack) == 0 {
		return "/"
	}

	for _, v := range stack {
		res += "/" + v
	}

	return res
}
