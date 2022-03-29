# hundred-combat
Golang进阶实战100题
### 1.交替打印字母和数字
---
问题描述
使用两个`goroutine`交替打印序列，一个`gorountine`打印数字，另一个`goroutine`打印字母,最终效果如下：
   
`12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ2728`
   
   解决思路及[源码](https://github.com/coding-today/hundred-combat/blob/main/advancedCombat/code01.go)   
   

### 2. 判断字符串中字符是否全都不同
***
问题描述
请实现一个算法，确定一个字符串的所有字符【是否全都不同】。这里我们要求【不允许使用额外的存储结构】。 给定一个string，请返回一个bool值,true代表所有字符全都不同，false代表存在相同的字符。 保证字符串中的字符为【ASCII字符】。字符串的⻓度小于等于【3000】。      

解题思路及[源码](https://github.com/coding-today/hundred-combat/blob/main/advancedCombat/code02.go)


### 3 翻转字符串
***
问题描述
    请实现一个算法，在不使用【额外数据结构和存储空间】的情况下，翻转一个给定的字符串（可以使用单个过程变量）。
    给定一个string,请返回一个string,为翻转后的字符串。保证字符串的长度小于等于5000。
思路
    翻转字符串其实是将一个字符串以中间字符为粥，前后翻转，即将str[len]赋值给ste[0],将ste[0]赋值给str[len]

解题思路及[源码](https://github.com/coding-today/hundred-combat/blob/main/advancedCombat/code03.go)



### 4 判断两个给定的字符串排序后是否一致
***
问题描述
	给定两个字符串，请编写程序，确定其中一个字符串的字符重新排列后，能否变成另一个字符串。 这里规定【大小写为不同字符】，且考虑字符串重点空格。给定一个string s1和一个string s2，请返回一个bool， 代表两串是否重新排列后可相同。
保证两串的⻓度都小于等于5000。

思路
	首先要保证字符串⻓度小于5000。之后只需要一次循环遍历s1中的字符在s2是否都存在即可。

解题思路及[源码](https://github.com/coding-today/hundred-combat/blob/main/advancedCombat/code04.go)

### 5 字符串替换问题
***
问题描述
	请编写一个方法，将字符串中的空格全部替换为“%20”。 假定该字符串有足够的空间存放新增的字符，并且知道字符串的真实⻓度(小于等于1000)，同时保证字符串由【大 小写的英文字母组成】。
给定一个string为原始的串，返回替换后的string。

思路
	两个问题，第一个是只能是英文字母，第二个是替换空格。

解题思路及[源码](https://github.com/coding-today/hundred-combat/blob/main/advancedCombat/code05.go)

### 6 机器人坐标问题
***
问题描述
	有一个机器人，给一串指令，L左转 R右转，F前进一步，B后退一步，问最后机器人的坐标，最开始，机器人位于 0 0，方向为正Y。
可以输入重复指令n : 比如 R2(LF) 这个等于指令 RLFLF。
问最后机器人的坐标是多少?

思路
	这里的一个难点是解析重复指令。主要指令解析成功，计算坐标就简单了

解题思路及[源码](./advancedCombat/code06.go)