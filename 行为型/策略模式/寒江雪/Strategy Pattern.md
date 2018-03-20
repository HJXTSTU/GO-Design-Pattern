## Strategy Pattern

&emsp;&emsp;策略模式在运行时动态地装配算法行为到对象中。

&emsp;&emsp;我们可以定义算法，封装它们，动态地切换它们。

## 实现

```go
type Operator interface{
	Apply(int,int)int
}

type Operation struct{
	Operator Operator
}

func (this *Operation)Operate(l,r int)int{
	return this.Operator.Apply(l,r)
}
```

&emsp;&emsp;定义具体的对象

```go
type Addition struct{}

func (Addition) Apply(lval, rval int) int {
	return lval + rval
}
```

&emsp;&emsp;使用

```go
func main() {
	operation := strategy.Operation{strategy.Addition{}}

	res  := operation.Operate(1,1)

	fmt.Println(res)
}
```

## 注意

* 策略模式可以让更换对象的内脏，而装饰者模式可以更换对象的外表。