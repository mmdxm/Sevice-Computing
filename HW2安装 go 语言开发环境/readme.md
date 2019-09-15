# 安装 go 语言开发环境实验报告

## 一、安装VScode

我是直接在ubuntu软件里搜索VSCode的，直接安装就行了

![2019-09-15 14-56-17 的屏幕截图](https://github.com/mmdxm/Sevice-Computing/blob/master/HW2%E5%AE%89%E8%A3%85%20go%20%E8%AF%AD%E8%A8%80%E5%BC%80%E5%8F%91%E7%8E%AF%E5%A2%83/img/2019-09-15%2014-56-17%20%E7%9A%84%E5%B1%8F%E5%B9%95%E6%88%AA%E5%9B%BE.png)

安装好后如图所示

![2019-09-15 14-56-29 的屏幕截图](https://github.com/mmdxm/Sevice-Computing/blob/master/HW2%E5%AE%89%E8%A3%85%20go%20%E8%AF%AD%E8%A8%80%E5%BC%80%E5%8F%91%E7%8E%AF%E5%A2%83/img/2019-09-15%2014-56-29%20%E7%9A%84%E5%B1%8F%E5%B9%95%E6%88%AA%E5%9B%BE.png)



## 二、安装 golang

在go的官网里选择对应的版本，先下载

![2019-09-15 15-06-05 的屏幕截图](https://github.com/mmdxm/Sevice-Computing/blob/master/HW2%E5%AE%89%E8%A3%85%20go%20%E8%AF%AD%E8%A8%80%E5%BC%80%E5%8F%91%E7%8E%AF%E5%A2%83/img/2019-09-15%2015-06-05%20%E7%9A%84%E5%B1%8F%E5%B9%95%E6%88%AA%E5%9B%BE.png)

然后将下载的压缩包，并提取到 /usr/local 目录，在 /usr/local/go 中创建Go目录树。
tar -C /usr/local -xzf go$VERSION.$OS-$ARCH.tar.gz
并将此行添加到你的 /etc/profile（全系统安装）或 $HOME/.profile 文件中

![2019-09-15 15-11-19 的屏幕截图](https://github.com/mmdxm/Sevice-Computing/blob/master/HW2%E5%AE%89%E8%A3%85%20go%20%E8%AF%AD%E8%A8%80%E5%BC%80%E5%8F%91%E7%8E%AF%E5%A2%83/img/2019-09-15%2015-11-19%20%E7%9A%84%E5%B1%8F%E5%B9%95%E6%88%AA%E5%9B%BE.png)

## 三、安装和配置

首先创建工作空间
mkdir $HOME/gowork
配置环境变量，在 ~/.profile 文件中添加:

export GOPATH=$HOME/gowork

export PATH=$PATH:$GOPATH/bin

![2019-09-15 15-08-54 的屏幕截图](https://github.com/mmdxm/Sevice-Computing/blob/master/HW2%E5%AE%89%E8%A3%85%20go%20%E8%AF%AD%E8%A8%80%E5%BC%80%E5%8F%91%E7%8E%AF%E5%A2%83/img/2019-09-15%2015-08-54%20%E7%9A%84%E5%B1%8F%E5%B9%95%E6%88%AA%E5%9B%BE.png)

然后
source $HOME/.profile

检查配置，可以看到

![2019-09-15 15-14-50 的屏幕截图](https://github.com/mmdxm/Sevice-Computing/blob/master/HW2%E5%AE%89%E8%A3%85%20go%20%E8%AF%AD%E8%A8%80%E5%BC%80%E5%8F%91%E7%8E%AF%E5%A2%83/img/2019-09-15%2015-14-50%20%E7%9A%84%E5%B1%8F%E5%B9%95%E6%88%AA%E5%9B%BE.png)





然后在vocode中，编译运行，可以看到运行结果



![2019-09-15 17-49-25 的屏幕截图](https://github.com/mmdxm/Sevice-Computing/blob/master/HW2%E5%AE%89%E8%A3%85%20go%20%E8%AF%AD%E8%A8%80%E5%BC%80%E5%8F%91%E7%8E%AF%E5%A2%83/img/2019-09-15%2017-49-25%20%E7%9A%84%E5%B1%8F%E5%B9%95%E6%88%AA%E5%9B%BE.png)

## 四、安装必要的工具和插件

工具和插件方面，科学上网就可以了

![2019-09-15 17-46-17 的屏幕截图](https://github.com/mmdxm/Sevice-Computing/blob/master/HW2%E5%AE%89%E8%A3%85%20go%20%E8%AF%AD%E8%A8%80%E5%BC%80%E5%8F%91%E7%8E%AF%E5%A2%83/img/2019-09-15%2017-46-17%20%E7%9A%84%E5%B1%8F%E5%B9%95%E6%88%AA%E5%9B%BE.png)



## 五、安装与运行 go tour

科学上网的话，我go get连接的是谷歌的网页

![2019-09-15 17-51-43 的屏幕截图](https://github.com/mmdxm/Sevice-Computing/blob/master/HW2%E5%AE%89%E8%A3%85%20go%20%E8%AF%AD%E8%A8%80%E5%BC%80%E5%8F%91%E7%8E%AF%E5%A2%83/img/2019-09-15%2017-51-43%20%E7%9A%84%E5%B1%8F%E5%B9%95%E6%88%AA%E5%9B%BE.png)

不过连不上。。。

![2019-09-15 17-51-56 的屏幕截图](https://github.com/mmdxm/Sevice-Computing/blob/master/HW2%E5%AE%89%E8%A3%85%20go%20%E8%AF%AD%E8%A8%80%E5%BC%80%E5%8F%91%E7%8E%AF%E5%A2%83/img/2019-09-15%2017-51-56%20%E7%9A%84%E5%B1%8F%E5%B9%95%E6%88%AA%E5%9B%BE.png)



## 六、完成第一个包的创建

1.环境变量已经在前面配好，这里不再描述.

2.包路径

![2019-09-15 20-36-29 的屏幕截图](https://github.com/mmdxm/Sevice-Computing/blob/master/HW2%E5%AE%89%E8%A3%85%20go%20%E8%AF%AD%E8%A8%80%E5%BC%80%E5%8F%91%E7%8E%AF%E5%A2%83/img/2019-09-15%2020-36-29%20%E7%9A%84%E5%B1%8F%E5%B9%95%E6%88%AA%E5%9B%BE.png)

3.第一个程序

![2019-09-15 20-37-31 的屏幕截图](https://github.com/mmdxm/Sevice-Computing/blob/master/HW2%E5%AE%89%E8%A3%85%20go%20%E8%AF%AD%E8%A8%80%E5%BC%80%E5%8F%91%E7%8E%AF%E5%A2%83/img/2019-09-15%2020-37-31%20%E7%9A%84%E5%B1%8F%E5%B9%95%E6%88%AA%E5%9B%BE.png)

4.第一个库

![2019-09-15 20-38-55 的屏幕截图](https://github.com/mmdxm/Sevice-Computing/blob/master/HW2%E5%AE%89%E8%A3%85%20go%20%E8%AF%AD%E8%A8%80%E5%BC%80%E5%8F%91%E7%8E%AF%E5%A2%83/img/2019-09-15%2020-38-55%20%E7%9A%84%E5%B1%8F%E5%B9%95%E6%88%AA%E5%9B%BE.png)



5.文件结构

![2019-09-15 20-40-05 的屏幕截图](https://github.com/mmdxm/Sevice-Computing/blob/master/HW2%E5%AE%89%E8%A3%85%20go%20%E8%AF%AD%E8%A8%80%E5%BC%80%E5%8F%91%E7%8E%AF%E5%A2%83/img/2019-09-15%2020-40-05%20%E7%9A%84%E5%B1%8F%E5%B9%95%E6%88%AA%E5%9B%BE.png)

6.测试
![2019-09-15 20-41-13 的屏幕截图](https://github.com/mmdxm/Sevice-Computing/blob/master/HW2%E5%AE%89%E8%A3%85%20go%20%E8%AF%AD%E8%A8%80%E5%BC%80%E5%8F%91%E7%8E%AF%E5%A2%83/img/2019-09-15%2020-41-13%20%E7%9A%84%E5%B1%8F%E5%B9%95%E6%88%AA%E5%9B%BE.png)

7.远程包



![2019-09-15 20-51-13 的屏幕截图](https://github.com/mmdxm/Sevice-Computing/blob/master/HW2%E5%AE%89%E8%A3%85%20go%20%E8%AF%AD%E8%A8%80%E5%BC%80%E5%8F%91%E7%8E%AF%E5%A2%83/img/2019-09-15%2020-51-13%20%E7%9A%84%E5%B1%8F%E5%B9%95%E6%88%AA%E5%9B%BE.png)



# 实验就完成了！
