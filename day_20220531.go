package main

func main() {
	// A var w = nil
	// B var x interface{} = nil
	// C var y string = nil
	// D var z error = nil

	// 问题:上面的ABCD四个赋值操作那些是正确的?
	// 答案: BD nil只能赋值给指针、chan、func、interface、map、slice等引用类型的变量， nil仅仅声明了变量 但是没有对变量进行赋值 但是值类型的变量需要初始化为对应的零值 所以值类型的变量不能赋值为nil
}
