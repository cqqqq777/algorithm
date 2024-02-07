package winter20

// 一条包含字母 A-Z 的消息通过以下映射进行了 编码 ：
//
// 'A' -> "1"
// 'B' -> "2"
// ...
// 'Z' -> "26"
// 要 解码 已编码的消息，所有数字必须基于上述映射的方法，反向映射回字母（可能有多种方法）。例如，"11106" 可以映射为：
// "AAJF" ，将消息分组为 (1 1 10 6)
// "KJF" ，将消息分组为 (11 10 6)
// 注意，消息不能分组为  (1 11 06) ，因为 "06" 不能映射为 "F" ，这是由于 "6" 和 "06" 在映射中并不等价。
// 给你一个只含数字的 非空 字符串 s ，请计算并返回 解码 方法的 总数 。
// 题目数据保证答案肯定是一个 32 位 的整数。
func numDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}
	pre, cur := 1, 1
	for i := 1; i < len(s); i++ {
		tmp := cur
		// 如果当前字符是 0，那么它只能和前一个字符组合，否则就是前一个字符的解码方法数
		if s[i] == '0' {
			if s[i-1] == '1' || s[i-1] == '2' {
				cur = pre
			} else {
				return 0
			}
			// 如果当前字符是 1 或者 2 并且后一个字符是 1 到 6 之间的数字，那么它可以和前一个字符组合，否则就是前一个字符的解码方法数
		} else if s[i-1] == '1' || (s[i-1] == '2' && s[i] >= '1' && s[i] <= '6') {
			cur += pre
		}
		pre = tmp
	}
	return cur
}