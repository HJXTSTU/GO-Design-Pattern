## Builder Pattern

&emsp;&emsp;生成器模式将复杂对象的构造与表示分开，以便相同的构建过程可以创建不同的表示形式。<br>

&emsp;&emsp;然而在Go语言中，如果向Builder传递一个结构，代码中就会充满检查代码，判断某个字段是否为空。为了避免这种现象，我们只需要在Builder中放置一个配置结构。<br>

## 实现

&emsp;&emsp;我们来实现一个构造器构造一个灯的过程。<br>

```go
// 先定义几个类型
type Color string
type LampStatus bool
type Brand string

// 定义颜色常量
const (
	BlueColor  Color = "blue"
	GreenColor       = "green"
	RedColor         = "red"
)

// 定义品牌常量
const (
	OppleBulb Brand = "OPPLE"
	Osram           = "OSRAM"
)

//	Lamp Builder define
type Builder interface {
	Color(Color) LampBuilder
	Brand(Brand) LampBuilder
	Build() LampOperation
}

type LampBuilder struct {
	Lamp	// 配置结构
}


func (lb LampBuilder) Color(c Color) LampBuilder {
	lb.color=c
	return lb
}

func (lb LampBuilder) Brand(b Brand) LampBuilder {
	lb.brand=b
	return lb
}

func (lb LampBuilder) Build() LampOperation {
    // 新的产品产生过程
	lamp := Lamp{color: lb.color, brand: lb.brand, status: false}
	return lamp
}

func NewBuilder() Builder {
	return LampBuilder{}
}

type LampOperation interface {
	Open() error
	Close() error
	ProductionIllustrative()
}

// 灯的定义
type Lamp struct {
	color  Color
	brand  Brand
	status LampStatus
}

func (l Lamp) Open() error {
	if l.status {
		return errors.New("Lamp is opened")
	}
	fmt.Println("Open lamp.")
	l.status = true
	return nil
}

func (l Lamp) Close() error {
	if !l.status {
		return errors.New("Lamp is closed")
	}
	fmt.Println("Close lamp.")
	l.status = true
	return nil;
}

func (l Lamp) ProductionIllustrative() {
	fmt.Println("I'm a lamp.")
	fmt.Println("Color:" + l.color)
	fmt.Println("Brand:" + l.brand)
}
```

<br>

&emsp;&emsp;使用:<br>

```go
func main(){
	b := builder.NewBuilder()
	lamp_1 := b.Color(builder.BlueColor).Brand(builder.Osram).Build()
	lamp_1.Open()
	lamp_1.ProductionIllustrative()

	lamp_2 := b.Color(builder.GreenColor).Brand(builder.OppleBulb).Build()
	lamp_2.ProductionIllustrative()
}

```