/*
__author__ = 'robin-luo'
__date__ = '2023/12/06 15:53'
*/

package solution

func candy(ratings []int) int {
	n := len(ratings)
	candies := make([]int, n)
	sum := 0
	for i := range candies {
		candies[i] = 1
	}

	for i := 0; i < n-1; i++ {
		if ratings[i] < ratings[i+1] {
			candies[i+1] = candies[i] + 1
		}
	}

	for i := n - 1; i > 0; i-- {
		if ratings[i-1] > ratings[i] {
			candies[i-1] = max(candies[i-1], candies[i]+1)
		}
	}

	for _, candy := range candies {
		sum += candy
	}

	return sum
}
