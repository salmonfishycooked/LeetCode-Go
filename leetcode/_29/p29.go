package _29

import "math"

func divide(dividend int, divisor int) (ans int) {
	// 溢出情况（只有这一种溢出情况）
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}

	// 将被除数与除数符号化正，最终的符号用op来确定
	op := 1
	if dividend < 0 {
		op = -op
		dividend = -dividend
	}
	if divisor < 0 {
		op = -op
		divisor = -divisor
	}

	// 倍增除数，得到其倍增的最大值（保证这个值小于被除数）
	// quantum 用于保存除数的倍数
	quantum := 1
	op1, op2 := dividend, divisor
	for op2<<1 < op1 {
		op2 <<= 1
		quantum <<= 1
	}

	// 相减，不够减则缩小除数
	for op1 >= divisor {
		if op1 >= op2 {
			op1 -= op2
			ans += quantum
		}
		if op2 == divisor {
			continue
		}
		op2 >>= 1
		quantum >>= 1
	}

	if op == -1 {
		return -ans
	}
	return
}
