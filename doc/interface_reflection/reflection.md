# 反射

参考文章

- [Golang的反射reflect深入理解和示例](https://juejin.im/post/5a75a4fb5188257a82110544)

## 1. 反射理解

**变量的最基本信息就是类型type和值value。** 在Golang实现中，每个interface变量都有一个对应的pair，用于记录实际变量的值和类型（value, type）。反射就是用来检测存储在接口变量内部（值value; 类型concrete type）pair对的一种机制。



## 2. 反射包reflect

反射包的`Type`表示一个Go类型。

```go
type Type interface {
    
}

func TypeOf(i interface{}) Type
```

`reflect.TypeOf`直接返回我们想要的type类型，如float64，int，各种pointer、struct等真实的类型。



反射包的Value表示

```go
type Value struct {
    // contains filtered or unexported fields
}
func ValueOf(i interface{}) Value
```

`reflect.ValueOf`直接返回我们想要的具体的值，如1.23或者类似&{1 ”Allen Wu“ 25}这样的结构体值



也就是说反射可以将”接口类型变量“转换为”反射类型对象“，反射类型就是`reflect.Type`和`reflect.Value`这两种。



后续：

- 获取接口信息的使用（已知原有类型和未知原有类型）
- 通过reflect.Value设置实际变量的值
- 通过relfect.ValueOf来进行方法的调用
- reflect的性能