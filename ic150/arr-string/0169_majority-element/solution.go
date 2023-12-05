/*
__author__ = 'robin-luo'
__date__ = '2023/12/05 15:26'
*/

package solution

func majorityElement(nums []int) int {
	major, count := 0, 0
	for _, num := range nums {
		if count == 0 {
			major = num
		}

		if major == num {
			count++
		} else {
			count--
		}
	}

	return major
}
