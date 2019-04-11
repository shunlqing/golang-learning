/*
使用接口传参
排序函数接收Sorter接口类型的参数,即只要实现了该接口的类型变量,该排序函数可以对其排序
*/ 
package main

import (
	"fmt"
)

type Sorter interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

func Sort(data Sorter) {
	for i := 1; i < data.Len(); i++ {
		for j := 0; j < data.Len() - i; j++ {
			if data.Less(j+1, j) {
				data.Swap(j, j+1)
			}
		}
	}
}

func IsSorted(data Sorter) bool {
	n := data.Len()
	for i := n - 1; i > 0; i-- {
		if data.Less(i, i-1) {
			return false
		}
	}
	return true
}

type IntArray []int

func (p IntArray) Len() int 	{ return len(p) }
func (p IntArray) Less(i, j int) bool { return p[i] < p[j]}
func (p IntArray) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

func main() {
	data := []int{75, 48, 28, 39, 10, 999}
	a := IntArray(data)
	Sort(a)
	if !IsSorted(a) {
		panic("fail")
	}
	fmt.Print(a)
}