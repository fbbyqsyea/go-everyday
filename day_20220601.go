package main

import "fmt"

type People interface {
	Speak(string) string
}

type Student struct{}

func (stu Student) Speak(think string) (talk string) {
	if think == "love" {
		talk = "you are a good boy"
	} else {
		talk = "hi"
	}
	return talk
}

func main() {

	// 	也就是说, Student不是People类型, 但是*Student是. 就这么简单. 只是当写下

	// peo := Student{}
	// peo.Speak("hello")
	// 时,编译器会隐式替换为(&peo).Speak("Hello")而已, 给人以Student也有func Speak(string) string方法的错觉.
	// cannot use (Student literal) (value of type Student) as People value in variable declaration: Student does not implement People (method Speak has pointer receiver)
	var peo People = Student{}
	think := "love"
	fmt.Println(peo.Speak(think))
}
