#container里面的容器#


****


主要介绍了list的使用规则,主要重点都是围绕使用来的。注意下面几点即可：

	1. 链表开箱即用
	2. 不能把自己定义的Element类型给库函数使用。

****

##链表的开箱即用##

	个人理解就是指的你在调用一些函数的时候，会去check一下链表的状态。
	如果元素和链表不属于同一个，那么后面的操作就没必要了，还有一些元素
	添加的函数，会对判断链表状态，然后有目的性的Init操作。

	l.lazyInit()// 惰性初始化函

****

##链表的基本操作##

	基本上函数自己看就可以用了。

***
##ring的用法##

	func doubleValue(in interface{}){
	fmt.Println(in)
	}
	
	func testRing(){
		r := ring.New(10)
		for i :=0;i < r.Len();i++{
			r.Value = i
			r = r.Next()
		}
	
		r.Do(doubleValue) //func (r *Ring) Do(f func(interface{}))  //对链表中任意元素执行f操作，如果f改变了r，则该操作造成的后果是不可预期的。
		r = r.Move(2) ////返回移动n个位置（n>=0向前移动，n<0向后移动）后的元素，r不能为空。
		fmt.Println(r.Value)
		r = r.Next() //获取下一个
		fmt.Println(r.Value)
		r = r.Prev() //获取上一个
		fmt.Println(r.Value)
		r.Unlink(2) //从当前node的next开始删除n个元素
		for i :=0;i < r.Len();i++{
			fmt.Println("After remove 2 nodes:",r.Value)
			r = r.Next()
		}
	}

***
##思考题##

1.ring包中的循环链表在哪些场景中用

	循环链表个人感觉主要用于保存固定的数据，比如保存最近100次的操作。

2.你使用过container/heap包中的堆吗？它的适用场景又有哪些？

	预留处理！！！！！！！！！！！！！！！！
		
