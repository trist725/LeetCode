package Happy_Number

// https://leetcode.cn/problems/happy-number/

func isHappy(n int) bool {
	calcSum := func(n int) (sum int) {
		for n > 0 {
			sum += (n % 10) * (n % 10)
			n /= 10
		}
		return sum
	}

	exist := make(map[int]struct{})
	for n != 1 {
		if _, ok := exist[n]; ok {
			return false
		}
		exist[n] = struct{}{}
		n = calcSum(n)
	}
	return true
}
