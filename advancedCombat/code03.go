package main

/*
问题描述
    请实现一个算法，在不使用【额外数据结构和存储空间】的情况下，翻转一个给定的字符串（可以使用单个过程变量）。
    给定一个string,请返回一个string,为翻转后的字符串。保证字符串的长度小于等于5000。
思路
    翻转字符串其实是将一个字符串以中间字符为粥，前后翻转，即将str[len]赋值给ste[0],将ste[0]赋值给str[len]

*/
import "fmt"


func reverString(s string)(string, bool){
    str := []rune(s)
    length := len(s)
    if length > 5000 {
        return s,false 
    }
    fmt.Println(str)
    for i:=0;i<length/2;i++ {
        str[i],str[length-i-1] = str[length-i-1],str[i]
    }
    fmt.Println(str)
    return string(str),true 
}

func main(){
    testStr := "abcddfrty"
    result,ok := reverString(testStr)
    if ok {
        fmt.Printf("rever before:%s, rever after:%s\n",testStr,result)
    }
}

