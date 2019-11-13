# 开发 web 服务程序

### 1.概述

​      开发简单 web 服务程序 cloudgo，了解 web 服务器工作原理。

##### 任务目标：

	1.熟悉 go 服务器工作原理
	2.基于现有 web 库，编写一个简单 web 应用类似 cloudgo。
	3.使用 curl 工具访问 web 程序
	4.对 web 执行压力测试
### 2、任务要求

    1.编程 web 服务程序 类似 cloudgo 应用。
        ~要求有详细的注释
        ~是否使用框架、选哪个框架自己决定 请在 README.md 说明你决策的依据
    2.使用 curl 测试，将测试结果写入 README.md
    3.使用 ab 测试，将测试结果写入 README.md。并解释重要参数。
### 3、完成情况

​	我选取的是beego框架， beego是一个快速开发Go应用的http框架，beego可以用来快速开发API、Web、后端服务等各种应用， 目前，今日头条、百度云盘、腾讯等公司都使用了beego，而且beego框架相对灵活，支持工程化的环境搭建，也支持单个小型服务器的搭建。

​	首先我在工作区中配置了beego的环境，然后顺着老师给的博客和网上的教程，构建了一个如下的简单web程序。

```go
package main
import (
    "github.com/astaxie/beego"//用beego框架
)
type MainController struct {
    beego.Controller          //定义Controll为maincontroller，maincontroller可以调用controller的所有函数
}
func (this *MainController) Get() {
	this.Ctx.WriteString("This is a webpage run by beego")  //利用controller，往网页里写入内容
}

func main() {
    beego.Router("/", &MainController{})//注册路由器，告诉beego用户有网络请求时，beego调用对应的beego
    beego.Run()                      //最终run这个微型服务器

} //默认端口880

//访问  http://localhost:8080/
```
首先运行Cloud.go文件

![1](https://github.com/mmdxm/Sevice-Computing/blob/master/HW6/Image/1.png)接下来在浏览器访问http://localhost:8080/，可以看到

![2](https://github.com/mmdxm/Sevice-Computing/blob/master/HW6/Image/2.png)

下面是curl测试

![3](https://github.com/mmdxm/Sevice-Computing/blob/master/HW6/Image/3.png)

下面是ab测试

![4](https://github.com/mmdxm/Sevice-Computing/blob/master/HW6/Image/4.png)

![5](https://github.com/mmdxm/Sevice-Computing/blob/master/HW6/Image/5.png)



![6](https://github.com/mmdxm/Sevice-Computing/blob/master/HW6/Image/6.png)
