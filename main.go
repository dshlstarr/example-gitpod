package main

import (
	"fmt"
)

func main() {
	fmt.Println("---- general go generic usage ----")
	PrintIDAndSum("John", Sum[float32], 23.2, 3, 4)

	fmt.Println("---- declare with \"var\" ----")
	fmt.Println(var01[int32](2, 3, 4))
	
	fmt.Println("---- declare with \"new\" ----")
	fmt.Println(var02[int32](2, 3, 4))
	
	fmt.Println("---- var vs new on pointer value ----")
	fmt.Println(newVsVarPtr[int]())
}

// Numeric expresses a type constraint satisfied by numeric types.
type Numeric interface {
	~int | ~int32 | ~int64 | ~float32 | ~float64
}

type SumFn[T Numeric] func(...T) T

func Sum[T Numeric](args ...T) T {
	var sum T
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return sum

}

func PrintIDAndSum[T string, K Numeric] (id T, sum SumFn[K], values ...K) {
	fmt.Printf("%s has %v values!\n", id, sum(values...))
}

func var01[T Numeric](args ...T) T {
	var defaultT T
	var sum = &defaultT

	for i := 0; i < len(args); i++ {
		*sum += args[i]	
	}
	return *sum
}

func var02[T Numeric](args ...T) T {
	sum := new(T)

	for i := 0; i < len(args); i++ {
		*sum += args[i]	
	}
	return *sum
}

// newVsVarPtr returns two pointers on different declarations: var vs new()
// var with return with <nil> since no instance of T allocated, 
// new() would be new instance addr of T
func newVsVarPtr[T any]() (*T, *T){
	var v *T
	n := new(T)
	return v, n
}