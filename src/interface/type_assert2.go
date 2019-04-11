// 类型断言第二种类型:判断一个接口值是否实现了另一种接口
package main 

import (
	"fmt"
)

type Stringer interface {
	String() string
}

type Stringer2 interface {
	String() string
}

type Type1 struct {
	data string
}

func (t Type1) String() string {
	return t.data
}

func main() {
	t := Type1{"test data"}
	
	var st Stringer2
	st = t

	if tv, ok := st.(Stringer); ok { //st必须是个接口值,如果是个具体类型值会报错
		fmt.Printf("v implement String(): %s\n", tv.String()) //成功的话,tv的接口类型是Stringer
	}
}