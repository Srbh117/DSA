package dsa

func TwoSum(nums []int, target int) []int {
	if len(nums) < 2 {
		return []int{}
	}
	seen := map[int]int{}
	for i, v := range nums {
		complement := target - v
		if j, ok := seen[complement]; ok {
			return []int{j, i}
		}
		seen[v] = i
	}
	return []int{}
}

func minPartitions(n string) int {
	ans := 0
	for _, v := range n {
		ans = max(ans, int(v-'0'))
	}
	return ans
}
