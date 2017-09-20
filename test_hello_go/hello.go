package main

import (
	//package 别名调用
	ios "fmt"
	"strings"
)
///结构体声明
type node struct {
	name string
	next *node
}

func main() {
	//常量
	const pI float32 = 3.1415926;

	//全局变量，声明变量１
	var name = "yzp";

	//类型重命名
	type i int;

	//声明变量２
	var age i; age = 10;
	var a i =12;
	ios.Println(age);

	//声明变量３
	c := "asdad"
	d := pI;

	//多变量声明１
	var vname1, vname2, vname3 i;
	vname1, vname2, vname3 = 10,11,13

	//多变量声明2
	v1,v2,v3 := 10,11,"aaa";

	//多变量多类型声明,这种因式分解关键字的写法一般用于声明全局变量
	var (
		vs1 i
		vs2 float32
		vs3 string
	)
	vs1,vs2,vs3=12,23.3,"bbb";


	//链表
	var n node;
	n.name="yzpp";
	n.next=nil;

	ios.Println(d,a,vname1, vname2, vname3,v1,v2,v3,vs1,vs2,vs3,n.name,n.next);
	ios.Println(name + " hello go! " + c);
	ios.Println(strings.Compare(c,name));
	ios.Println(strings.HasPrefix(name,"y"));
}
