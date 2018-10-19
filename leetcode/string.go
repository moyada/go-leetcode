package main

import (
	"math"
)

func main() {
	//fmt.Println(reverseString("hello"))
	//fmt.Println(reverse(901000))
	//fmt.Println(firstUniqChar("loveleetcode"))
	//fmt.Println(isPalindrome("`l;`` 1o1 ??;l`"))
	//fmt.Println(myAtoi("0-1"))
	//fmt.Println(longestCommonPrefix([]string{"sssdfg", "ssss"}))
}

/**
请编写一个函数，其功能是将输入的字符串反转过来。
 */
func reverseString(s string) string {
	size := len(s)
	bytes := make([]byte, size)
	i := 0
	for index := size - 1; index >= 0; index-- {
		bytes[i] = s[index]
		i++
	}
	return string(bytes)
}

/**
给定一个 32 位有符号整数，将整数中的数字进行反转。

示例 1:

输入: 123
输出: 321
 示例 2:

输入: -123
输出: -321
示例 3:

输入: 120
输出: 21
 */
func reverse(x int) int {
	if x == 0 {
		return 0
	}

	positive := x < 0
	if positive {
		x = -x
	}
	reverse := 0
	var tmp int
	for x != 0 {
		tmp = x % 10
		x /= 10
		if 0 != reverse {
			reverse *= 10
		}
		if tmp == 0 {
			continue
		}
		reverse += tmp
	}

	if positive {
		reverse = -reverse
	}

	if math.MinInt32 > reverse || reverse > math.MaxInt32 {
		return 0
	}

	return reverse
}

/**
给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1。
 */
func firstUniqChar(s string) int {
	if s == "" {
		return -1
	}

	mark := [26]int {}
	for _, v := range s {
		mark[v - 'a'] += 1;
	}

	index := -1
	for k, v := range s {
		if mark[v - 'a'] == 1 {
			index = k
			break
		}
	}
	return index
}

/**

给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的一个字母异位词。

示例 1:

输入: s = "anagram", t = "nagaram"
输出: true
示例 2:

输入: s = "rat", t = "car"
输出: false
 */
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	markS := [26]int {}
	markT := [26]int {}

	keys := make([]int, 0)
	var index int
	for _, v := range s {
		index = int(v) - 'a'
		if 0 == markS[index] {
			keys = append(keys, index)
		}
		markS[index] += 1
	}
	for _, v := range t {
		markT[v - 'a'] += 1
	}

	for _, v := range keys {
		if markS[v] != markT[v] {
			return false;
		}
	}

	return true
}

/**
给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。
 */
func isPalindrome(s string) bool {
	if s == "" {
		return true
	}

	length := len(s)
	var result bool
	j := length-1
	var x, y int
	for i := 0; i < length; {
		if i >= j {
			result = true
			break
		}
		if !isAlphanumeric(s[i]) {
			i++
			continue
		}
		if !isAlphanumeric(s[j]) {
			j--
			continue
		}

		x = getInt(s[i])
		y = getInt(s[j])

		if x != y {
			result = false
			break
		}
		i++
		j--
	}
	return result
}

func isAlphanumeric(t byte) bool {
	return '0' <= t && t <= '9' || 'A' <= t && t <= 'Z' || 'a' <= t && t <= 'z'
}

func getInt(t byte) int {
	if t < 'a' {
		return int(t) + 32
	} else {
		return int(t)
	}
}

/**
实现 atoi，将字符串转为整数。

在找到第一个非空字符之前，需要移除掉字符串中的空格字符。如果第一个非空字符是正号或负号，选取该符号，并将其与后面尽可能多的连续的数字组合起来，这部分字符即为整数的值。如果第一个非空字符是数字，则直接将其与之后连续的数字字符组合起来，形成整数。

字符串可以在形成整数的字符后面包括多余的字符，这些字符可以被忽略，它们对于函数没有影响。

当字符串中的第一个非空字符序列不是个有效的整数；或字符串为空；或字符串仅包含空白字符时，则不进行转换。

若函数不能执行有效的转换，返回 0。
 */
func myAtoi(str string) int {
	if str == "" {
		return 0
	}
	first := str[0]
	if first != '-' && first != '+' && (first < '0' || first > '9') && first != ' ' {
		return 0
	}

	sum := 0
	positive := 0
	for _, s := range str {
		if positive == 0 && s == ' ' {
			continue
		}
		if s == '-' || s == '+' {
			if positive > 0 {
				break
			}
			if s == '-' {
				positive = 1
			} else {
				positive = 2
			}
			continue
		}

		if '0' <= s && s <= '9' {
			if positive == 0 {
				positive = 2
			}
			sum *= 10
			sum += int(s) - '0'
		} else {
			break
		}
	}


	if sum < math.MinInt32 {
		return math.MaxInt32
	}

	if 1 == positive {
		sum = -sum
	}

	if sum < math.MinInt32 {
		return math.MinInt32
	}

	if sum > math.MaxInt32 {
		return math.MaxInt32
	}

	return sum
}

/**
14. 编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。
 */
func longestCommonPrefix(strs []string) string {
	common := "";
	if nil == strs {
		return common
	}
	length := len(strs)
	if length == 0 {
		return common
	}

	bytes := make([][]byte, length)
	minLen := math.MaxInt8
	for i := 0; i < length; i++ {
		bytes[i] = []byte(strs[i])
		tmpLen := len(bytes[i])
		if minLen > tmpLen {
			minLen = tmpLen
		}
	}

	if 0 == minLen {
		return common
	}

	for i := 0; i < minLen; i++ {
		b :=  bytes[0][i]
		for j := 1; j < length; j++ {
			if b != bytes[j][i] {
				return common
			}
		}
		common = common + string(b)
	}
	return common
}