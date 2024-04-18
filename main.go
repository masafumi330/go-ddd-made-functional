package main

import "fmt"

func main()  {
	fmt.Println(Sum([]int{1, 2, 3, 4, 5}))
	fmt.Println(Sum([]float64{1.1, 2.2, 3.3, 4.4, 5.5}))
	// fmt.Println(Sum([]string{"a", "b", "c", "d", "e"})) <- compile errorになる
}

// Sumの型パラメータをint or float64だけに制約したい場合、インターフェースを定義して、それを型パラメータに指定することで実現できる。
type Number interface {
	int | float64
}

func Sum[T Number](nums []T) T{
	var sum T
	for _, num := range nums {
		sum += num
	}
	return sum
}
