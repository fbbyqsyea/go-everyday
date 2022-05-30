package main

import "fmt"

const (
	x = iota // 0
	_        // 赋值  1  丢弃该值
	y        // 2
	z = "kk" // kk   index 3
	k        // 没赋值 同上一个变量  index 4
	p = iota // 重新回到iota  5
	q        // 6
)

const (
	_  = iota
	KB = 1 << (10 * iota) // 1 10
	MB = 1 << (10 * iota) // 2 20
	GB = 1 << (10 * iota) // 3 30
	TB = 1 << (10 * iota) // 4 40
	PB = 1 << (10 * iota) // 5 50
)

func main() {
	fmt.Println(x, y, z, k, p, q)   // 0 2 kk kk 5 6
	fmt.Println(KB, MB, GB, TB, PB) // 1024 1048576 1073741824 1099511627776 1125899906842624
}

// 说明: iota是golang中的常量计数器  只能在golang的常量中使用。iota在const中出现时被重置为0，const每新增一行变量申明 不管是否使用iota iota都将计数一次。
