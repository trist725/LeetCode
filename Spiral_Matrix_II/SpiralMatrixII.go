package Spiral_Matrix_II

// https://leetcode.cn/problems/spiral-matrix-ii/
// 59 螺旋矩阵

func GenerateMatrix(n int) [][]int {
	var (
		result = make([][]int, n)
		x, y   int
		round  int = 1 // 第几圈
		count  int = 1
	)

	for k, _ := range result {
		result[k] = make([]int, n)
	}

	// 总共n/2圈
	for round <= n/2 {
		// 每条边左闭右开四次循环赋值
		for ; y < n-round; y++ {
			result[x][y] = count
			count++
		}
		for ; x < n-round; x++ {
			result[x][y] = count
			count++
		}
		for ; y >= round; y-- {
			result[x][y] = count
			count++
		}
		for ; x >= round; x-- {
			result[x][y] = count
			count++
		}

		x++
		y++
		round++
	}

	// 特殊处理中心点
	if n%2 == 1 {
		result[x][y] = count
	}

	return result
}
