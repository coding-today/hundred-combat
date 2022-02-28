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