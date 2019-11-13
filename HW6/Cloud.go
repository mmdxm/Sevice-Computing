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