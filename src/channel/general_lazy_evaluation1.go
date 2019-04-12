// 通用的惰性生成器

// 工厂模式

//练习: 使用工厂模式生成前10个斐波那契数

package main

import (
	"fmt"
)

type Any interface{}
type EvalFunc func(Any) (Any, Any)

type fibState struct {
	ex uint64
	exx uint64
}

func main() {
	evenFunc := func(state Any) (Any, Any) {
		os := state.(int)
		ns := os + 2
		return os, ns
	}

	fibFunc := func(state Any) (Any, Any) {
		fs := state.(fibState)

		if fs.ex == 0 && fs.exx == 0 {
			fs.ex, fs.exx = 1, 0
		} else {
			fs.ex, fs.exx = fs.ex + fs.exx, fs.ex
		}

		return fs.ex, fs
	}

	even := BuildLazyIntEvaluator(evenFunc, 0)

	fib := BuildLazyFibEvaluator(fibFunc, fibState{0, 0})

	for i := 0; i < 10; i++ {
		fmt.Printf("%vth even: %v\n", i, even())
	}

	for i := 0; i < 10; i++ {
		fmt.Printf("%vth fib: %v\n", i, fib())
	}
}

// 工厂函数
// 参数evalFunc: 计算下一个返回值和下一个状态参数
// 参数initState: 一个初始状态
// 返回值: 无参,返回值是生成序列的函数
func BuildLazyEvaluator(evalFunc EvalFunc, initState Any) func() Any {
	retValChan := make(chan Any)
	loopFunc := func() {
		var actState Any = initState
		var retVal Any
		for {
			retVal, actState = evalFunc(actState)
			retValChan <- retVal
		}
	}
	retFunc := func() Any {
		return <- retValChan
	}
	go loopFunc()
	return retFunc
}

func BuildLazyIntEvaluator(evalFunc EvalFunc, initState Any) func() int {
	ef := BuildLazyEvaluator(evalFunc, initState)
	return func() int {
		return ef().(int)
	}
}

func BuildLazyFibEvaluator(evalFunc EvalFunc, initState Any) func() uint64 {
	ef := BuildLazyEvaluator(evalFunc, initState)
	return func() uint64 {
		return ef().(uint64)
	}
}