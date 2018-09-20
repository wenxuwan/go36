# 第一章重点：#

## 1.GO 语言安装的时候主要有三个环境变量： ##

    GOROOT、GOPATH 和 GOBIN
  
  通过 go env 可以查看go语言的一系列的环境变量的配置.

	go env

  **GOROOT**：代表go的安装路径，一般你安装完go语言之后就有了。
  
  **GOPATH**：一个目录路径，也可以包含多个目录路径，每个目录都代表了go语言的一个“工作区”。目录下面有三个目录，src,bin,pkg。

	src : 用来存放源代码文件。
	
	bin: 目录里面存放的都是通过 go install 命令安装后，由 Go 命令源码文件生成的可执行文件。
	有两种情况下，bin 目录会变得没有意义。
	当设置了有效的 GOBIN 环境变量以后，bin 目录就变得没有意义。
	如果 GOPATH 里面包含多个工作区路径的时候，必须设置 GOBIN 环境变量，否则就无法安装 Go 程序的可执行文件。
	
	pkg:用来存放通过 go install 命令安装后的代码包的归档文件(.a 文件)

  **GOBIN**：存放可执行文件的文件目录。


**Go 的源码文件分类：**

	
	（1）命令源码文件：
		声明自己属于 main 代码包、包含无参数声明和结果声明的 main 函数。
		命令源码文件被安装以后，GOPATH 如果只有一个工作区，那么相应的
		可执行文件会被存放当前工作区的 bin 文件夹下；如果有多个工作区，
		就会安装到 GOBIN 指向的目录下。
	
	（2）库源码文件
		
		库源码文件就是不具备命令源码文件上述两个特征的源码文件。存在于某个代码包中的普通的源码文件。

	（3）测试源码文件

		名称以 _test.go 为后缀的代码文件，并且必须包含 Test 或者 Benchmark 名称前缀的函数。
  
详情可以参照[https://studygolang.com/articles/10572](https://studygolang.com/articles/10572)

## 2. Go 语言源码的组织方式 ##


代码包的名字和代码包下面的代码的package name 一般是同名的，如果不同，在构建的过程中就以代码包为准。


GO语言也是以代码包为基本的组织单位。比如：

	import proto "github.com/golang/protobuf/proto"

在工作区内，一个代码包的引入，其实就是以GOPATH/src为基准的相对路径。

## 3.go build 的使用 ##

go build 默认不会重新编译目标代码所依赖的代码包。当然如果依赖的的.a文件不存在或者源代码变化还是会被编译。

	go build -a //强制编译依赖的库
	go build -x //显示执行了哪些操作
	go build -v //可以看到编译的所有代码包的名字和 -a 搭配很好
	
	
# 总结 #

GOPATH 个人认为主要是给GO语言的管理提供了一个相对的“绝对路径”。这样保证代码的管理是有一个标准目录作为“根目录的”。




# 问题 #

1. Go 语言在多个工作区中查找依赖包的时候是以怎样的顺序进行的？
	
	如果有多个工作区，那么是按照GOPATH里面的先后顺序查找的

2. 如果在多个工作区中都存在导入路径相同的代码包会产生冲突吗？
	
	不会冲突，像上面说的会按照顺序查找执行

# 拓展 #


如果我的src下面的文件夹名字叫做hello，但package的名字叫做HelloGO

	package HelloGo
	
	import "fmt"
	
	func PrintHello(){
		fmt.Println("Hello Go")
	}

这时候如果你想调用这个包，代码应该怎么写，
		
		import HelloGo？
		还是import hello

答案肯定是import hello的，因为GOPATH的特性，只会按照目录来查找包，但你调用的时候就不能用目录名字来调用了：
	
	package main

	import "fmt"
	import "hello"
	
	func main() {
		fmt.Println("Hello, world!")
		HelloGo.PrintHello() //导入包的时候用文件夹名字，调用函数用package的名字
	}

	
