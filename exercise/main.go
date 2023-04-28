package main

import (
	"fmt"
)

/*
var skillParam = flag.String("skill", "", "skill to perform")

func main() {
	flag.Parse()

	var skill = map[string]func(){
		"file": func() {
			fmt.Println(("chicked fire"))
		},
		"run": func() {
			fmt.Println("solider run")
		},
		"fly": func() {
			fmt.Println("angel fly")
		},
	}

	if f, ok := skill[*skillParam]; ok {
		f()
	} else {
		fmt.Println("skill not found")
	}
}
*/

//函数类型实现接口

//调用器接口
/*
type Invoker interface {
	Call(interface{})
}

//结构体类型

type Struct struct {
}

// 实现Invoker的Call
func (s *Struct) Call(p interface{}) {
	fmt.Println("from struct", p)
}

// 函数定义为类型
type FuncCaller func(interface{})

func (f FuncCaller) Call(p interface{}) {
	f(p)
}
func main() {
	//声明接口变量
	var invoker = Invoker

	//实例化结构体

	s := new(Struct)

	//将实例化的结构体赋值到接口
	invoker = s

	invoker.Call("hello")
}
*/
type X struct {
	x, by int
}
type A struct {
	ax, ay int
	X
}

type B struct {
	A
	bx, by float32
}

func main() {
	b := B{
		bx: 3.0,
		by: 4.0,
		A: A{
			ax: 1,
			ay: 2,
			X: X{
				x:  1,
				by: 2,
			},
		},
	}
	fmt.Println(b.ax, b.ay, b.bx, b.by)
	fmt.Println(b.A)
}
