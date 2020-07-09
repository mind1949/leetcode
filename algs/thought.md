## 思路
要求不用乘、除、取余这三种运算来计算商，想到除法其实就是减法，用减法来计算商。
根据题目的说法除数不为0，所以就不用考虑除数为零的情况。这样就有三个注意点了：
* 判断商的正负性；
* 计算被除数可以减去多少个除数；
* 若商大于`2^31 -1`，则返回`2^31-1`；

### 代码
```golang
func divide(dividend int, divisor int) int {
	// 判断被除数与除数的正负性
	positive := true
	if dividend < 0 {
		positive = !positive
		dividend = -dividend
	}
	if divisor < 0 {
		positive = !positive
		divisor = -divisor
	}

	// 计算商
	quotient := 0
	for dividend > divisor {
		dividend -= divisor
		quotient++
		if quotient > math.MaxInt32 {
			quotient = math.MaxInt32
		}
	}

	if !positive {
		return -quotient
	}
	return quotient
}
```

**提交结果**

![图片.png](https://pic.leetcode-cn.com/76e90cfcd0f4dd84605ea19112e4c4d23228d01d5676591b7e1b61fdc38566a6-%E5%9B%BE%E7%89%87.png)

这就尴尬了。。超时了。先看下复杂度，后面再优化下

**复杂度**
时间复杂度：**O(m/n)**；空间复杂度**O(1)**
