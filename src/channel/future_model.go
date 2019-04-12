// Future模式

package main

import (
	"fmt"
)

// 普通模式:进行Product计算时,计算b的逆矩阵时,需要等待a的逆计算完成
// func InverseProduct(a Matrix, b Matrix) {
// 	a_inv := Inverse(a)
// 	b_inv := Inverse(b)
// 	return Product(a_inv, b_inv)
// }

func InverseProduct(a Matrix, b Matrix) Matrix{
	a_inv_future := InverseFuture(a)
	b_inv_future := InverseFuture(b) //两者计算不用相互等待
	return Product(<-a_inv_future, <-b_inv_future)
}

func InverseFuture(a Matrix) chan Matrix {
	future := make(chan Matrix)
	go func() {
		future <- Inverse(a)
	}()
	return future
}

//fake code
type Matrix struct {
	data string
}

func Inverse(a Matrix) Matrix {
	a.data += "+inv"
	return a
}

func Product(a Matrix, b Matrix) Matrix {
	d := a.data + "x" + b.data
	var m Matrix = Matrix{d}
	return m
}

func main() {
	var m1, m2 Matrix
	m1.data = "AAA"
	m2.data = "BBB"
	fmt.Println(InverseProduct(m1, m2))
}