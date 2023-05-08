> [LeetCode 122][https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-ii]

# 解法：贪心

在价格低谷的时候买入，在价格高峰的时候卖出

```go
func maxProfit(prices []int) (ans int) {
    n := len(prices)
    // 遍历 prices 数组，在价格低谷的时候买入，在价格高峰的时候卖出
    for i := 1; i < n; {
        for i < n && prices[i] <= prices[i - 1] {
            i++
        }
        buy := prices[i - 1] // 在 i-1 的时候是低谷，买入
        for i < n && prices[i] >= prices[i - 1] {
            i++
        }
        sell := prices[i - 1] // 在 i-1 的时候是高峰，卖出
        ans += sell - buy
    }
    return
}
```

