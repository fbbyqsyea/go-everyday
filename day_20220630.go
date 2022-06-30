package main

import "fmt"

func main() {

	fmt.Println(sum(1, 2))
	fmt.Println(sum(1.1, 1.21))

	// 泛型变量的实例化
	var MySlice1 Slice1[int] = []int{1, 2, 3, 4}
	MySlice2 := Slice1[int]{1, 2, 3, 4}
	fmt.Println(MySlice1, MySlice2)

	var MyMap1 MAP1[string, int] = map[string]int{
		"aaa": 1000000,
	}

	MyMap2 := map[string]int{
		"bbbb": 1000000,
	}
	fmt.Println(MyMap1, MyMap2)

}

// 泛型的定义
type Computable interface {
	int | float64
}

type IFS interface {
	int | float64 | string
}

type Slice1[T IFS] []T

type MAP1[KEY string, VALUE IFS] map[KEY]VALUE

type Struct1[T IFS] struct {
	Title   string
	Content T
}

func sum[T Computable](a, b T) T {
	return a + b
}
