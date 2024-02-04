/*
 * Author: robin-luo
 * Created time: 2024-02-03 17:45:27
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-03 18:04:02
 */

package solution

func candy(ratings []int) int {
	n := len(ratings)
	candies := make([]int, n)
	for i := range candies {
		candies[i] = 1
	}

	// ->
	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		}
	}

	// <-
	for i := n - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			candies[i] = max(candies[i], candies[i+1]+1)
		}
	}

	res := 0
	for _, candy := range candies {
		res += candy
	}

	return res
}
