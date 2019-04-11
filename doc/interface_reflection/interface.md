# 接口

## 1. 什么是接口

接口是一种约定，说明对象可以完成哪些行为（通过定义方法（集））。



**接口类型**定义：

```go
type Namer interface {
    Method1(param_list) return_type
    Method2(param_list) return_type
    ...
} //这里定义了一个接口类型，取名为Namer
```

Go语言的接口类型可以定义变量，变量的值可以是实现了该接口的具体类型的实例或另一个接口值。试想一下，接口要完成其功能：<u>隐藏其所引用的真实具体类型实例，又准确调用其方法</u>，如何实现？接口值很明显要保存两样：**接收者（receiver）和方法表指针（method table ptr）**。metohd table ptr是不是与C++中类的虚函数表非常类似，这正是interface类型的变量具有多态特性的关键。



## 2. 类型断言之判断接口值的动态类型并转换为具体类型的值

接口变量所存储值的具体类型（**动态**类型）从接口是看不出来的，需要设计某种方式来获取变量中存储值的具体类型。通常使用**类型断言**来测试某个时刻`varI`是否包含具体类型`T`的值，这是类型的第一种情况。

```go
v := varI.(T) //varI必须是个接口变量，这种方式未检查运行时错误
```

如果转换合法，`v`是`varI`转换到类型`T`的值，否则，`v`是类型`T`的零值

更安全、更常用的类型断言方式：

```go
if v, ok := varI.(T); ok {
    Process(v)
    return
}
```

**T是接收者类型，是区分*T和T的，**

```go

type Square struct {
    side float32
}

type Shaper interface {
    Area() float32
}

func main() {
	var areaIntf Shaper
    sq1 := new(Square)
    sq1.size = 5
    
    areaIntf = sq1
    if t, ok := areaIntf.(*Square); ok {  //省略*，会引发编译错误
        // ...
    }
}

func (sq *Square) Area() float32 { //是*Square类型实现了接口，而不是Square实现了接口
    return sq.side * sq.side
} 
```



另一种形式进行类型判断，采用特殊形式switch：**type-switch**

```go
switch t := areaIntf.(type) {  //或者switch areaIntf.(type)
	case *Square:
	// ...
    case *Circle:
    //...
    case nil:
    //....
    default:
    //...
}
```



## 3. 类型断言之判断接口值是否实现了某个接口并转换为目标接口的值

类型断言的第二种情况：`v.(T)`的T是接口类型，类型断言检测v的动态类型是否满足T接口。这里的v也必须是个接口，不然编译器报错。

```go
type Stringer interface {
    String() string
}

if sv, ok := v.(Stringer); ok {
    fmt.Printf("v implements String(): %s\n", sv.String()) //note: sv, not v
} //判断值v是否实现了接口Stringer
```

如果这个检查成功，则检查结果的接口值sv的动态类型和动态值不变，但是接口值的接口类型被转换为接口类型Stringer（T）。换句话说，对一个接口类型的类型断言改变了类型的表述方式，改变了可以获取的方法集合，但是它保护了接口值内部的动态类型和值的部分。参考[Go的类型断言解析](https://studygolang.com/articles/11419)

如果检查失败，接下来这个操作会抛出panic，除非两个变量来接受检查结果。



### 语法糖：方法与接口

作用于**具体类型**的变量上的方法，不区分变量是指针类型还是值类型。

- 接收者是值类型，可以通过值调用，也可以通过指针调用（编译器自动先解引用）
- 接收者是指针类型，只能通过指针类型调用。

**但是，接口变量存储的具体值是不可寻址的。**

即，针对具体类型T，

- 存储了类型*T的接口，其可调用的方法集包括接收者为\*T或T的所有方法集
- 存储了类型T的接口，其可调用的方法集只包含接收者为T的方法，不包含接收者为*T的方法