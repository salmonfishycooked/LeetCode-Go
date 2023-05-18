> [LeetCode 4][https://leetcode.cn/problems/median-of-two-sorted-arrays]

# 解法：二分查找

```go
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n1, n2 := len(nums1), len(nums2)
	quantum := n1 + n2
	if quantum%2 == 1 {
		return float64(findRank(nums1, nums2, quantum/2))
	}
	return float64(findRank(nums1, nums2, quantum/2)+findRank(nums1, nums2, quantum/2-1)) / 2.0
}

// 找到两个有序数组合起来后的索引为rank的数
func findRank(nums1 []int, nums2 []int, rank int) int {
	n1, n2 := len(nums1), len(nums2)
	l1, r1 := -1, n1
	l2, r2 := -1, n2
	// 在 nums1 里面找 rank
	for l1+1 != r1 {
		mid := l1 + (r1-l1)>>1
		l2, r2 = -1, n2
		for l2+1 != r2 {
			mid2 := l2 + (r2-l2)>>1
			if nums2[mid2] < nums1[mid] {
				l2 = mid2
			} else {
				r2 = mid2
			}
		}

		// 防止卡重复数
		for r2 < n2 && nums2[r2] == nums1[mid] {
			nowRank := mid + (r2 + 1)
			if nowRank == rank {
				return nums1[mid]
			}
			r2++
		}

		// 计算nums1[mid]的rank值
		nowRank := mid + (l2 + 1)
		if nowRank < rank {
			l1 = mid
		} else if nowRank > rank {
			r1 = mid
		} else {
			return nums1[mid]
		}
	}

	// 在 nums2 里面找 rank
	return findRank(nums2, nums1, rank)
}
```

> 时间复杂度：$O(log(m + n))$
>
> 空间复杂度：$O(1)$