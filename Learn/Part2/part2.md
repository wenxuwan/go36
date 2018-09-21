# 第二章重点：#

**Go 的源码文件分类：**
	
在第一章的时候已经提及到了这个关键问题：

	（1）命令源码文件：
		声明自己属于 main 代码包、包含无参数声明和结果声明的 main 函数。
		命令源码文件被安装以后，GOPATH 如果只有一个工作区，那么相应的
		可执行文件会被存放当前工作区的 bin 文件夹下；如果有多个工作区，
		就会安装到 GOBIN 指向的目录下。
	
	（2）库源码文件
		
		库源码文件就是不具备命令源码文件上述两个特征的源码文件。存在于某个代码包中的普通的源码文件。

	（3）测试源码文件

		名称以 _test.go 为后缀的代码文件，并且必须包含 Test 或者 Benchmark 名称前缀的函数。


**GO语言的flag包：**

微课堂主要介绍了一下flag包的用法。并且付上了官方文档。所以在这里稍微列举一下用法。

# flag的定义： #

	flag.String(), Bool(), Int() //这里这是列举了几个

然后就是两种定义的方式：

	var ip = flag.Int("flagname", 1234, "help message for flagname") //ip 为指针类型，Int或者
	String返回的都是指针类型
	
	flag.IntVar(&flagvar, "flagname", 1234, "help message for flagname")//此处需要在外面手动定义flagvar

第一种方式是返回一个值，第二种方式是把变量的传进去。


通过flag.Var()绑定自定义类型，自定义类型需要实现Value接口(Receives必须为指针)，如：

	flag.Var(&flagVal, "name", "help message for flagname")
对于这种类型的flag，默认值为该变量类型的初始值

## flag的解析 ##

	flag.Parse()

解析函数将会在碰到第一个非flag命令行参数时停止，非flag命令行参数是指不满足命令行语法的参数，如命令行参数为cmd --flag=true abc则第一个非flag命令行参数为“abc”

## flag解析后的参数使用 ##

	fmt.Println("ip has value ", *ip)
	fmt.Println("flagvar has value ", flagvar)

#命令行语法#

	-flag //只支持bool类型
	-flag=x
	-flag x //只支持非bool类型

注意第三种只支持非bool类型，原因就是如果文件的名字为false的话那么就会产生歧义。看例子：

	package main

	import (
		"flag"
		"fmt"
	)
	
	var name string
	var right bool
	func init() {
		flag.StringVar(&name, "name", "everyone", "The greeting object.")
		flag.BoolVar(&right, "check", false, "The right ob")
	}
	
	func main() {
		flag.Parse()
		Hello(name)
		fmt.Println(right)
	}

	这里采用的还是老师的例子，只多加了个bool类型。
	我们编译之后生成exe

	demo4.exe -check false
	demo4.exe -check
	
	对于这两种都会采用默认的bool类型true。

	demo4.exe -check=false

	只有这样的时候，才会采用false


#定义自己的参数使用说明#
	
	第一种：

	最简单就是修改Usages这个函数变量：

	var Usage = func() {
		fmt.Fprintf(CommandLine.Output(), "Usage of %s:\n", os.Args[0])
		PrintDefaults()
	}

	我们可以看一下源码里面采用的是第一个参数的名字也就是我们的exe文件。

	var Usage = func() {
		fmt.Fprintf(CommandLine.Output(), "HaoLin guid%s:\n", "wenxuwan")
		PrintDefaults()
	}
	这样你就可以定义输出了

	郝林老师还讲了更深入的定义自己的类型。主要是更深入的定义flag包里面的一些对象实现自己的方法，具体可以看下源码。
	

#思考题#

	1. 默认情况下，我们可以让命令源码文件接受哪些类型的参数值？
	这个从flag包的代码包里面就可以看到，这里贴一下班里一个同学的答案

		int(int|int64|uint|uint64),
		float(float|float64)
		string,
		bool,
		duration(时间),
		var(自定义)


	2. 我们可以把自定义的数据类型作为参数值的类型吗？如果可以，怎样做？
	
		当然是可以的：
		// Var defines a flag with the specified name and usage string. The type and
		// value of the flag are represented by the first argument, of type Value, which
		// typically holds a user-defined implementation of Value. For instance, the
		// caller could create a flag that turns a comma-separated string into a slice
		// of strings by giving the slice the methods of Value; in particular, Set would
		// decompose the comma-separated string into the slice.
		func Var(value Value, name string, usage string) {
			CommandLine.Var(value, name, usage)
		}

		type pepoInfo struct{
			name string
			age int
			homeAdress string
		}

		func (p * pepoInfo) Set(val string) error{
			peopleList := strings.Split(val, ",")
			p.name = peopleList[0]
			p.age,_ =  strconv.Atoi(peopleList[1])
			p.homeAdress = peopleList[2]
			return nil
		}	

		func (s *pepoInfo) String() string {
			infoList := strings.Split("wenxuwan,20,山东", ",")
			s.name = infoList[0]
			s.age,_ = strconv.Atoi(infoList[1])
			s.homeAdress = infoList[2]
			return "It's none of my business"
		}
		func main(){

		var people pepoInfo
		flag.Var(&people, "info", "hao ma fan")
		flag.Parse()
		//打印结果slice接收到的值
		fmt.Println(people)
		}

	demo4.exe --info wenxuwan,10,shouguang

