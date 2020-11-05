go基本程序结构
```go
package main//包 表明代码所在的模块

import "fmt"//引入代码依赖

//功能实现
func main()  {
	fmt.Println("hello world")
}

```
运行命令
```shell script
go run hello_world.go
//输出 hello world

go build hello_world.go
//生成hello_world二进制文件

./hello_world
//输出 hello world
```

go应用程序入口：需要注意以下几点

1. 必须是main包: `package main`
2. 必须是main方法：  `func main()`
3. 文件名不一定是: `main.go`

退出返回值
与其他语言的主要差异
- Go中的main函数不支持任何返回值
- Go通过使用 `os.Exit` 来返回状态

示例1:
```go
package main

import (
	"fmt"
	"os"
)

func main()  {
	fmt.Println("hello world")
	os.Exit(0);
}
```
输出：
```
hello world
```

示例2：
```go
package main

import (
	"fmt"
	"os"
)

func main()  {
	fmt.Println("hello world")
	os.Exit(-1);
}
```

输出：
```
hello world
exit status 255
```

获取命令行参数：
与其他编程语言的主要差异
- main函数不支持传入参数  func main(~~arg []string~~)
- 在程序中直接通过`os.Args` 获取命令行参数

示例：
```go
package main

import (
	"fmt"
	"os"
)

func main()  {
	fmt.Println(os.Args)
}

```
输出os.Args 会得到一个数组 数组的第一个值为 我们的二进制命令 第二个值为我们输入的值

```
[/var/folders/pr/99jdzjf12b9_xlqlpmt9s_lw0000gn/T/go-build905857895/b001/exe/hello_world yang]

```
小实战：
```go
package main

import (
	"fmt"
	"os"
)

func main()  {
	if len(os.Args) > 1{
        fmt.Println("hello world",os.Args[1])
    }
}

```
运行：
```
go run hello_world.go yang
//输出 hello world yang
```

