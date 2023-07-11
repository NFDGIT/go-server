该项目是一个个人学习项目

该项目是一个多模块结构项目

每个子目录对应一个模块：每个模块的配置在 相应 文件夹下的 .mod 文件中
一级目录的```.work``` 文件管理和组织各个模块

>go mod tidy 类似flutter项目中的 flutter pub get
多模块项目中只需cd 到入口module中 ，执行该命令



1. 入口模块路径 ：``` go mod tidy```

2. 运行命令: ```go run web-service-gin```
3. ```go get -u golang.org/x/tools/.. ```   used to install or update the Go tools provided by the golang.org/x/tools package.

-------

### 编程范式
1. 面向对象 使用 struct 和 method 实现 面向编程 具体如下
```go 
// 定义一个 Circle 结构体的方法
func (c Circle) Area() float64 {
    return 3.14 * c.radius * c.radius
}

// 定义另一个 Circle 结构体的方法
func (c Circle) Circumference() float64 {
    return 2 * 3.14 * c.radius
}

func main() {
    // 创建一个 Circle 结构体对象
    c := Circle{radius: 5}

    // 调用 Circle 结构体的方法
    area := c.Area()
    fmt.Println("Area:", area)

    circumference := c.Circumference()
    fmt.Println("Circumference:", circumference)
}
```
2. 面向过程
   
