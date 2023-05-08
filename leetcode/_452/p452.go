package _452

import "sort"

func findMinArrowShots(points [][]int) (ans int) {
	n := len(points)
	// 将 points 数组按每个一维切片的第二个值从小到大排序
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})
	// 遍历 points
	// 记录第一个一维切片的第二个数做为最大引爆限度
	max := points[0][1]
	for i := 0; i < n; {
		// 继续遍历数组，直到遍历到某个一维数组的第一个值大于最大引爆限度
		for i < n && points[i][0] <= max {
			i++
		}
		// 所需弓箭数量+1，更新最大引爆限度继续遍历
		ans += 1
		if i < n {
			max = points[i][1]
		}
	}
	return ans
}
