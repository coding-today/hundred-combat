package main

import (
	"fmt"
	"strings"
)

/*
问题描述
	给定两个字符串，请编写程序，确定其中一个字符串的字符重新排列后，能否变成另一个字符串。 这里规定【大小写为不同字符】，且考虑字符串重点空格。给定一个string s1和一个string s2，请返回一个bool， 代表两串是否重新排列后可相同。
保证两串的⻓度都小于等于5000。

解题思路
	首先要保证字符串⻓度小于5000。之后只需要一次循环遍历s1中的字符在s2是否都存在即可。

*/

func isRegroup(s1, s2 string) bool {
	group1 := []rune(s1)
	group2 := []rune(s2)

	if len(group1) > 5000 || len(group2) > 5000 || len(group1) != len(group2) {
		return false
	}
	//在ascii码中48-57为是个阿拉伯数字，0-31为控制字符或通信专用字符 32 是空格 65-90为26个大写字母 97-122为26个小写字母 rune其实就是int32
	fmt.Println(group1)
	for _, v := range s1 {
		fmt.Println(v)
		fmt.Println(string(v))
		if strings.Count(s1, string(v)) != strings.Count(s2, string(v)) {
			return false
		}

	}
	return true
}

func main() {
	str1 := "456abcdrft444"
	str2 := "555bacdrft555"
	ok := isRegroup(str1, str2)
	if ok {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}
