package utils

import "bytes"

// GenFixedLengthChineseChars 指定长度随机中文字符(包含复杂字符)
func GenFixedLengthChineseChars(length int) string {

	var buf bytes.Buffer

	for i := 0; i < length; i++ {
		buf.WriteRune(rune(RandIntRange(19968, 40869)))
	}
	return buf.String()
}

// GenRandomLengthChineseChars 指定范围随机中文字符
func GenRandomLengthChineseChars(start, end int) string {
	length := RandIntRange(start, end)
	return GenFixedLengthChineseChars(length)
}
