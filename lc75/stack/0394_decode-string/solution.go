/*
__author__ = 'robin-luo'
__date__ = '2023/10/31 13:33'
*/

package solution

func decodeString(s string) string {
	cntStk := make([]int, 0)
	strStk := make([]string, 0)
	currNum := 0
	currStr := ""

	for _, ch := range s {
		if ch >= '0' && ch <= '9' {
			currNum = currNum*10 + int(ch-'0')
		} else if ch == '[' {
			cntStk = append(cntStk, currNum)
			strStk = append(strStk, currStr)
			currNum = 0
			currStr = ""
		} else if ch == ']' {
			prevNum := cntStk[len(cntStk)-1]
			cntStk = cntStk[:len(cntStk)-1]
			prevStr := strStk[len(strStk)-1]
			strStk = strStk[:len(strStk)-1]

			for i := 0; i < prevNum; i++ {
				prevStr += currStr
			}

			currStr = prevStr
		} else {
			currStr += string(ch)
		}
	}

	return currStr
}
