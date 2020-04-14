package fibonacci

import (
	"log"
	"testing"
)

/*
go test -bench=.
go test    -test.bench

go test    -test.bench  ^BenchmarkBubbleSort$
*/
func TestName(t *testing.T) {

	num := 10

	log.Println(Fibonacci(num))
	log.Println(Fibonacci2(num))
	log.Println(Fibonacci3(num))
}

func Fibonacci(x int) int {
	if x == 0 {
		return 1
	}

	if x == 1 {
		return 1
	}

	return Fibonacci(x-1) + Fibonacci(x-2)

}

func BenchmarkName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fibonacci(10)

	}
}

func BenchmarkName2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fibonacci2(10)
	}
}

func BenchmarkName3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fibonacci3(10)
	}
}

func Fibonacci2(x int) int {
	if x == 0 {
		return 1
	}
	if x == 1 {
		return 1
	}
	var array = make([]int, x+1)
	//array := []int{1, 1}
	array[0] = 1
	array[1] = 1
	for i := 2; i <= x; i++ {
		array[i] = array[i-1] + array[i-2]
	}
	return array[x]
}

func Fibonacci3(x int) int {
	if x == 0 {
		return 1
	}
	if x == 1 {
		return 1
	}

	num1 := 1
	num2 := 1
	var num int

	for i := 2; i <= x; i++ {

		num = num1 + num2

		num1 = num2
		num2 = num

	}
	return num
}
