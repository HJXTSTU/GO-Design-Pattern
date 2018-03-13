## Facade Pattern

&emsp;&emsp; 外观模式其实就是把复杂的东西给封装，由统一的接口进行操作。这样可以简化用户的使用。<br>

## 例子

&emsp;&emsp;如果要开一家饭馆，一个饭馆分为采购的，管仓库的，切菜的，炒菜的，吃菜的。<br>

&emsp;&emsp;而采购的，管仓库的，切菜的，炒菜的，是厨房。吃菜的，是客户。客户并不希望知道厨房工作的细节，只希望告诉厨房自己想吃什么就行了，他不想告诉厨房要去买什么菜，怎么切，怎么炒。而巧的是，厨房也只希望给客户提供有限的选择即可，具体工作的细节由自己控制就行。因此厨房和客户之间只隔着一个菜单和服务员。<br>

### 实现

&emsp;&emsp;先来实现厨房的各个部门<br>

```go
package facade

// 虽然把蔬菜定义在厨房的采购部并不合理
// 但是我还是这么做了
const (
	STATUS_BUYED  = iota
	STATUS_STORED
	STATUS_CUTED
	STATUS_COOKED
	STATUS_EATED
)

type VegStatus int

type Vegetable struct {
	name   string
	status VegStatus
}

func BuyVegetable(name string) *Vegetable {
	return &Vegetable{name, STATUS_BUYED}
}

func Eat(veg *Vegetable) {
	veg.status = STATUS_EATED
}
```

<br>

```go
package facade
//	买回来就要保存
var storage []*Vegetable = make([]*Vegetable, 0)

func SaveVegetables(veg ...*Vegetable) {
	storage = append(storage,veg...)
	for _,v := range veg{
		v.status=STATUS_STORED
	}
}

func GetVegetables() *Vegetable {
	l := len(storage)
	res := storage[l-1]
	storage = storage[:l-1]
	return res
}
```

<br>

```go
package facade
//	切菜
//	具体拿去做什么不知道
//	只管切
func CutVegtable(veg ...*Vegetable)[]*Vegetable{
	for _,v := range veg{
		v.status=STATUS_CUTED
	}
	return veg
}
```

<br>

```go
package facade
//	炒菜
//	给谁吃不知道
//	炒就是了
func CookVegtable(vec ...*Vegetable)[]*Vegetable{
	for _,v := range vec{
		v.status=STATUS_COOKED
	}
	return vec
}
```

<br>

```go
package facade

//	菜单
//	客户随意选择
//	制作方式可以由主厨决定
//	只要客户喜欢
func SauteVegtable()[]*Vegetable{
	qc := BuyVegetable("青菜")
	suan := BuyVegetable("蒜")
	jiang := BuyVegetable("姜")
	SaveVegetables(qc,suan,jiang)

	vegs := CookVegtable(CutVegtable(storage...)...)
	return vegs
}

```

<br>

### 使用

```go
func EatVegtables(veg ...*Vegetable){
	for _,v :=range veg{
		Eat(v)
	}
}

func main(){

	// No Facade
	bc := BuyVegetable("白菜")
	SaveVegetables(bc)
	vecs := CookVegtable(GetVegetables())
	EatVegtables(vecs...)
	for _,v := range vecs{
		fmt.Println(*v)
	}

	// Favade
	sauteVegtable := SauteVegtable()
	EatVegtables(sauteVegtable...)
	for _,v := range sauteVegtable{
		fmt.Println(*v)
	}
}
```

<br>

## 优缺点

### 优点

* 它对客户屏蔽子系统组件，因而减少了客户处理的对象的数目并使得子系统使用起来更加方便。 
* 它实现了子系统与客户之间的松耦合关系，而子系统内部的功能组件往往是紧耦合的。 <br>

### 缺点

* 增加新的子系统可能需要修改外观类或客户端的源代码，违背了“开闭原则”。
* 对客户访问子系统类做太多的限制则减少了可变性和灵活性。

## 应用情景

* 当要为一个复杂子系统提供一个简单接口时可以使用外观模式。 
* 客户程序与多个子系统之间存在很大的依赖性。 
* 在分层结构中，可以使用外观模式定义系统中每一层的入口，层与层之间不直接产生联系，而通过外观类建立联系，降低层之间的耦合度。<br>

<br>

