package main

import "fmt"

func main() {
	//defer栈 （先进后出）
	defer d1()
	defer d2()
	defer d3()
	//fmt.Println(f1()) //5
	//fmt.Println(f2()) //5
	//fmt.Println(f3()) //5
	//fmt.Println(f4()) //6
	//fmt.Println(f5()) //5

}

//defer执行顺序
func d1() {
	fmt.Println(1)
}

func d2() {
	fmt.Println(2)
}
func d3() {
	fmt.Println(3)
}

// defer执行机制
//1. 返回值赋值
//2. 执行defer （如果对返回值操作，则涉及值传递和地址传递）
//3.执行return指令

func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f2() int {
	x := 5
	defer func() {
		x += 1
	}()
	return x
}
func f3() int {
	x := 5
	defer func() {
		x = 6
	}()
	return x
}

func f4() (x int) {
	defer func() {
		x++
	}()
	return 5
}

func f5() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}

//结论
//未命名的返回值，defer中对返回值的操作指向不同的内存地址
//命名且同名的返回值，defer中对返回值的操作指向相同的内存地址
