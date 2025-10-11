package leetcode

/*
1010. 总持续时间可被 60 整除的歌曲

在歌曲列表中，第 i 首歌曲的持续时间为 time[i] 秒。
返回其总持续时间（以秒为单位）可被 60 整除的歌曲对的数量。形式上，我们希望下标数字 i 和 j 满足  i < j 且有 (time[i] + time[j]) % 60 == 0
*/
func numPairsDivisibleBy60(time []int) int {
	result := 0

	l := make([]int, 60)
	for _, n := range time {
		left := n % 60
		result += l[(60-left)%60]
		l[left]++
	}

	return result
}
