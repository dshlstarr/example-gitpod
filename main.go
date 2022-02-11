package main

import (
	"fmt"
)

func main() {
	print("test generic")
	PrintIDAndSum("John", Sum[int32], 2, 3, 4)
}

func print[T any](arg T) {
	fmt.Println(arg)
}

type Numeric interface {
	int | int32 | int64 | float32 | float64
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