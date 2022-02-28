package main

/*
问题描述：
    请实现一个算法，确定一个字符串的所有字符【是否全部不同】, 这里我们要求【不允许使用额外的存储结构】。给定一个string,请返回一个bool值 true代表所有字符全部不同，false代表存在相同字符、
    确保字符串为【ASCII字符】，字符串长度小于等于【3000】
解题思路：
    重点：第一个是ASCII字符，ASCII字符一共256个，其中128个是常用字符，可以再键盘上输入。128之后是无法再键盘上找到的。
          然后是全部不同，也就是字符串中的字符没有重复的，
          再次，不允许使用额外存储结构，且长度小于等于3000.
    如果允许使用其他额外存储结构，这个题目很好做。如果不允许的话，可以使用golang内置的方式实现。
*/ 
import (
    "fmt"
    "strings"
)



func isUniqueString(s string) bool {
    if strings.Count(s,"") >3000{
        return false 
    }

    for _,v := range s{
        if v > 127 {
            return false

        }
        if strings.Count(s,string(v)) > 1 {
            return false

        }
    }
    return true

}


func isUniqueStrings(s string) bool {
    if strings.Count(s,"") > 3000{
        return false

    }

    for k,v := range s{
    if v > 127{
        return false

    }
    if strings.Index(s,string(v)) != k {
        return false

    }
    
    }
    return true

}

func main(){
    testStr1 := "abvdefghopu"
    result1 := isUniqueString(testStr1)
    result2 := isUniqueStrings(testStr1)
    fmt.Println(result1)
    fmt.Println(result2)
    testStr2 := "aabbccdd"
    result3 := isUniqueString(testStr2)
    result4 := isUniqueStrings(testStr2)
    fmt.Println(result3)
    fmt.Println(result4)

}

