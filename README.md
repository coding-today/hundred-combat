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

### 7 下面代码可以运行吗？为什么
***
```go 
type Param map[string]interface{}
type Show struct {
   Param
}
  func main1() {
     s := new(Show)
     s.Param["RMB"] = 10000
}
```

##### 解析
共发现两个问题:
+  main 函数不能加数字。
+  new 关键字无法初始化 Show 结构体中的 Param 属性，所以直接对 s.Param 操作会出错。

### 8 请说出下面代码存在什么问题。
***
```go
type student struct {
  Name string
}
func zhoujielun(v interface{}) {
  switch msg := v.(type) {
  case *student, student:
msg.Name }
}
```

##### 解析
golang中有规定， switch type 的 case T1 ，类型列表只有一个，那么 v := m.(type) 中的 v 的类型就是T1类 型。
如果是 case T1, T2 ，类型列表中有多个，那 v 的类型还是多对应接口的类型，也就是 m 的类型。
所以这里 msg 的类型还是 interface{} ，所以他没有 Name 这个字段，编译阶段就会报错。具体[解释⻅](https://go.dev/ref/spec#Type_switches)


### 9 写出打印的结果
***
```go
type People struct {
  name string `json:"name"`
}
func main() {
  js := `{
    "name":"coding-today"
  }`
  var p People
  err := json.Unmarshal([]byte(js), &p)
  if err != nil {
    fmt.Println("err: ", err)
return
}
  fmt.Println("people: ", p)
}
```

##### 解析
按照 golang 的语法，小写开头的方法、属性或 struct 是私有的，同样，在 json 解码或转码的时候也无法上线 私有属性的转换。
题目中是无法正常得到 People 的 name 值的。而且，私有属性 name 也不应该加 json 的标签。

### 10 下面的代码是有问题的，请说明原因。
***
```go
type People struct {
  Name string
}
func (p *People) String() string {
  return fmt.Sprintf("print: %v", p)
}
func main() {
  p := &People{}
  p.String()
}
```
##### 解析

在golang中String()string方法实际上是实现了String的接口的，该接口定义在fmt/print.go 中：
```go
type Stringer interface {
  String() string
}
```
在使用 fmt 包中的打印方法时，如果类型实现了这个接口，会直接调用。而题目中打印 p 的时候会直接调用 p 实现的 String() 方法，然后就产生了循环调用。

### 11 请找出下面代码的问题所在。
***
```go
func main() {
  ch := make(chan int, 1000)
  go func() {
    for i := 0; i < 10; i++ {
      ch <- i
} }()
  go func() {
    for {
      a, ok := <-ch
      if !ok {
        fmt.Println("close")
return
}
      fmt.Println("a: ", a)
    }
  }()
  close(ch)
  fmt.Println("ok")
  time.Sleep(time.Second * 100)
}
```

#####  解析
在 golang 中 goroutine 的调度时间是不确定的，在题目中，第一个写 channel 的 goroutine 可能还未调 用，或已调用但没有写完时直接 close 管道，可能导致写失败，既然出现 panic 错误。

### 12 请说明下面代码书写是否正确。
***
```go
var value int32
func SetValue(delta int32) {
  for {
    v := value
    if atomic.CompareAndSwapInt32(&value, v, (v+delta)) {
break
} }
}
```
##### atomic.CompareAndSwapInt32 函数不需要循环调用。

### 13 下面的程序运行后为什么会爆异常。
***
```go
 
type Project struct{}
func (p *Project) deferError() {
  if err := recover(); err != nil {
    fmt.Println("recover: ", err)
  }
}
func (p *Project) exec(msgchan chan interface{}) {
  for msg := range msgchan {
    m := msg.(int)
    fmt.Println("msg: ", m)
  }
}
func (p *Project) run(msgchan chan interface{}) {
  for {
    defer p.deferError()
    go p.exec(msgchan)
    time.Sleep(time.Second * 2)
} }
func (p *Project) Main() {
  a := make(chan interface{}, 100)
  go p.run(a)
  go func() {
for {
a <- "1"
      time.Sleep(time.Second)
    } 
}()
  time.Sleep(time.Second * 100000000000000)
}
func main() {
  p := new(Project)
  p.Main()
}
```

##### 有以下几个问题
+  time.Sleep 的参数数值太大，超过了 1<<63 - 1 的限制。
+  defer p.deferError() 需要在协程开始出调用，否则无法捕获 panic 。


### 14  请说出下面代码哪里写错了
***
```go  
func main() {
  abc := make(chan int, 1000)
  for i := 0; i < 10; i++ {
abc <- i }
go func() {
for a:=rangeabc {
      fmt.Println("a: ", a)
    }
  }()
  close(abc)
  fmt.Println("close")
  time.Sleep(time.Second * 100)
}
```
##### 协程可能还未启动，管道就关闭了。

### 15 请说出下面代码，执行时为什么会报错
***
```go 
type Student struct {
  name string
}
func main() {
  m := map[string]Student{"people": {"zhoujielun"}}
  m["people"].name = "wuyanzu"
}
```
##### map的value本身是不可寻址的，因为map中的值会在内存中移动，并且旧的指针地址在map改变时会变得无效。 故如果需要修改map值，可以将 map 中的非指针类型 value ，修改为指针类型，比如使用 map[string]*Student .


### 16  请说出下面的代码存在什么问题?
***
```go 
type query func(string) string
func exec(name string, vs ...query) string {
  ch := make(chan string)
  fn := func(i int) {
    ch <- vs[i](name)
  }
  for i, _ := range vs {
    go fn(i)
}
return <-ch }
func main() {
  ret := exec("111", func(n string) string {
    return n + "func1"
  }, func(n string) string {
    return n + "func2"
  }, func(n string) string {
    return n + "func3"
  }, func(n string) string {
    return n + "func4"
  })
  fmt.Println(ret)
}
```
##### 依据4个goroutine的启动后执行效率，很可能打印111func4，但其他的111func*也可能先执行，exec只会返回一 条信息。

### 17 下面这段代码为什么会卡死?
***
```go 
package main
import (
    "fmt"
"runtime"
)
func main() { var i byte
    go func() {
        for i = 0; i <= 255; i++ {
} }()
    fmt.Println("Dropping mic")
    runtime.Gosched()
    runtime.GC()
    fmt.Println("Done")
}
 
```
##### 解析
Golang 中，byte 其实被 alias 到 uint8 上了。所以上面的 for 循环会始终成立，因为 i++ 到 i=255 的时候会溢出， i <= 255 一定成立。
也即是， for 循环永远无法退出，所以上面的代码其实可以等价于这样:
```go 
go func() {
    for {}
}
```
正在被执行的 goroutine 发生以下情况时让出当前 goroutine 的执行权，并调度后面的 goroutine 执行:
+ IO 操作 
+ Channel 阻塞 
+ system call 
+ 运行较⻓时间

如果一个 goroutine 执行时间太⻓，scheduler 会在其 G 对象上打上一个标志( preempt)，当这个 goroutine 内部发生函数调用的时候，会先主动检查这个标志，如果为 true 则会让出执行权。
main 函数里启动的 goroutine 其实是一个没有 IO 阻塞、没有 Channel 阻塞、没有 system call、没有函数调用的 死循环。
也就是，它无法主动让出自己的执行权，即使已经执行很⻓时间，scheduler 已经标志了 preempt。
而 golang 的 GC 动作是需要所有正在运行 goroutine 都停止后进行的。因此，程序会卡在 runtime.GC() 等待
所有协程退出。

### 18 写出下面代码输出内容。
***
```go 
package main
import (
  "fmt"
)
func main() {
  defer_call() 
}
func defer_call() {
    defer func() { fmt.Println("打印前") }() 
    defer func() { fmt.Println("打印中") }() 
    defer func() { fmt.Println("打印后") }()
    panic("触发异常") 
}
```

##### 解析

defer 关键字的实现跟go关键字很类似，不同的是它调用的是runtime.deferproc而不是runtime.newproc。 在 defer 出现的地方，插入了指令 call runtime.deferproc ，然后在函数返回之前的地方，插入指令 call
runtime.deferreturn 。
goroutine的控制结构中，有一张表记录 defer ，调用 runtime.deferproc 时会将需要defer的表达式记录在表
中，而在调用 runtime.deferreturn 的时候，则会依次从defer表中出栈并执行。 因此，题目最后输出顺序应该是 defer 定义顺序的倒序。 panic 错误并不能终止 defer 的执行。

### 19 以下代码有什么问题，说明原因
***
```go 
type student struct {
  Name string
  Age  int
}
func pase_student() {
  m := make(map[string]*student)
  stus := []student{
    {Name: "zhou", Age: 24},
    {Name: "li", Age: 23},
    {Name: "wang", Age: 22},
  }
  for _, stu := range stus {
    m[stu.Name] = &stu
  }
}
```
##### golang 的 for ... range 语法中， stu 变量会被复用，每次循环会将集合中的值复制给这个变量，因此，会导 致最后 m 中的 map 中储存的都是 stus 最后一个 student 的值。

### 20 下面的代码会输出什么，并说明原因
***
```go
func main() { 
runtime.GOMAXPROCS(1)
  wg := sync.WaitGroup{}
  wg.Add(20)
  for i := 0; i < 10; i++ {
    go func() {
      fmt.Println("i: ", i)
      wg.Done()
}() 
}
  for i := 0; i < 10; i++ {
    go func(i int) {
      fmt.Println("i: ", i)
      wg.Done()
    }(i)
}
wg.Wait() 
}
```

##### 解析

这个输出结果决定来自于调度器优先调度哪个G。从runtime的源码可以看到，当创建一个G时，会优先放入到下一 个调度的 runnext 字段上作为下一次优先调度的G。因此，最先输出的是最后创建的G，也就是9
```go 
func newproc(siz int32, fn *funcval) {
  argp := add(unsafe.Pointer(&fn), sys.PtrSize)
  gp := getg()
  pc := getcallerpc()
  systemstack(func() {
    newg := newproc1(fn, argp, siz, gp, pc)
_p_ := getg().m.p.ptr() //新创建的G会调用这个方法来决定如何调度
    runqput(_p_, newg, true)
    if mainStarted {
      wakep()
} })
}
...
  if next {
  retryNext:
oldnext := _p_.runnext //当next是true时总会将新进来的G放入下一次调度字段中
    if !_p_.runnext.cas(oldnext, guintptr(unsafe.Pointer(gp))) {
      goto retryNext
} 
if oldnext == 0 {
    return
}
  // Kick the old runnext out to the regular run queue.
  gp = oldnext.ptr()
}
```