package main

func GetValue() int {
	return 1
}

func main() {
	i := GetValue()
	switch i.(type) {
	case int:
		println("int")
	case string:
		println("string")
	case interface{}:
		println("interface")
	default:
		println("unknown")
	}
}

// 下面这段代码能否编译通过？如果可以，输出什么？

// 答案及解析：编译失败。考点：类型断言，类型断言的语法形如：i.(type)，其中 i 是接口，type 是固定关键字，需要注意的是，只有接口类型才可以使用类型断言。
