package main

import (
	"fmt"
)

func main() {
	PrintIDAndSum("John", Sum[float32], 23.2, 3, 4)
}

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