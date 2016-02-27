package pascal

func Triangle(n int) [][]int {
	var ans [][]int
	for i := 0; i < n; i++ {
		row := make([]int, i+1, i+1)
		row[0], row[i] = 1, 1
		for j := 1; j <= i/2; j++ {
			row[j] = ans[i-1][j-1] + ans[i-1][j]
			row[i-j] = row[j]
		}
		ans = append(ans, row)
	}
	return ans
}
