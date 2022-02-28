package main
/*

问题描述
   使用两个goroutine交替打印序列，一个goroutine打印数数字，另一个goroutine打印字母,最终效果如下：
   `12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ2728打印字母,最终效果如下：
 思路：
    使用channel来控制打印进度，使用两个channel,分别控制数据和字母的打印序列，数字打印完成后通过channel来通知字母打印，字母打印完成后通知数字打印，然后周而复始的工作。
   ```
*/

import ( 
    "fmt"
    "sync"
)

func main(){

    letter,number := make(chan bool), make(chan bool)
    wait := sync.WaitGroup{}
    go func () {
        i := 1
        for {
            select{
            case <- number:
                fmt.Print(i)
                i++
                fmt.Print(i)
                i++
                letter <- true

            }
        }
    }()
    wait.Add(1)
    go func(wait *sync.WaitGroup){
        i := 'A'
        for {
            select{
            case <- letter:
                if i>= 'Z'{
                    wait.Done()
                    return

                }
                fmt.Print(string(i))
                i++
                fmt.Print(string(i))
                i++
                number <- true

            }
        }
    }(&wait)

    number <-true
    wait.Wait()

}



