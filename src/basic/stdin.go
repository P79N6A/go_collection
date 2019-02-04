//Go程序设计语言
//统计键盘输入的内容出现两次或以上
package main

import(
  "fmt"
  "bufio"
  "os"
)

func main() {
  counts := make(map[string]int)
  input := bufio.NewScanner(os.Stdin) //创建扫描器类型
  for input.Scan() {                  //扫描器不断监听键盘输入，并将结尾的换行符去掉
    if input.Text() == "#" {          //当输入"#"时跳出循环
      break
    }
    counts[input.Text()]++
  }
  
  for line, n := range counts {
    if n > 1 {
      fmt.Printf("%s\t%d",line,n)
    }
  }
}
