/*
__author__ = 'robin-luo'
__date__ = '2024/01/30 16:31'
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
