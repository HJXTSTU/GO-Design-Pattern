## 简单工厂

&emsp;&emsp;简单工厂的实现思想，即创建一个工厂，将产品的实现逻辑集中在这个工厂中。<br>

```go
/***
 *	Simple Factory
 */
type FoodFactory struct {
}

func (ff FoodFactory) CreateFood(name string) Food {
	var s Food;
	switch name {
	case "Meat":
		s = new(Meat);
	case "Hamberger":
		s = new(Hamberger)
	}
	return s;
}

type Food interface {
	Eat()
}

type Meat struct {
}

type Hamberger struct {
}

func (m Meat) Eat() {
	fmt.Println("Eat meat.")
}

func (h Hamberger) Eat() {
	fmt.Println("Eat Hamberger.")
}

func main(){

	// Simple Factory
	f := FoodFactory{}
	f.CreateFood("Meat").Eat()
	f.CreateFood("Hamberger").Eat()
}
```

<br>

## 工厂方法

&emsp;&emsp;工厂方法用于弥补简单工厂的不足之处:<br>

* 当新增一种产品的时候，需要修改工厂逻辑<br>
* 简单工厂没有继承的结构<br>
* 一旦工厂瘫痪，整个系统都瘫痪<br>

&emsp;&emsp;工厂方法将产品的创建逻辑写在子类之中,代码如下<br>

```go
/***
 *	Factory Method
 */
type Factory interface {
	Create() Food
}

type MeatFactory struct {
}

func (mf MeatFactory) Create() Food {
	m := Meat{}
	return m
}

type HambergerFactory struct{

}

func (hf HambergerFactory) Create() Food {
	h := Hamberger{}
	return h
}

type Food interface {
	Eat()
}

type Meat struct {
}

type Hamberger struct {
}

func (m Meat) Eat() {
	fmt.Println("Eat meat.")
}

func (h Hamberger) Eat() {
	fmt.Println("Eat Hamberger.")
}

func main(){
	// Factory Method
	mf := MeatFactory{}
	mf.Create().Eat()
	hf := HambergerFactory{}
	hf.Create().Eat()
}
```

<br>

## 抽象工厂

&emsp;&emsp;抽象工厂是针对一个产品族而言的.<br>

&emsp;&emsp;产品族就好像套餐.一个套餐包含了好几种食品，而每一种食品都是一种类型的食物。举个例子.<br>

&emsp;&emsp;一个套餐定义为食品和饮料.而有一家餐馆的食品包括肉和汉堡，饮料包括CoCo和茶.这家餐馆不想单独出售某种食品和饮料，只卖套餐.<br>

&emsp;&emsp;于是老板定义:<br>

* 套餐A:肉和CoCo<br>
* 套餐B:汉堡和茶<br>

&emsp;&emsp;定好了之后,把套餐A外包给工厂A负责生产，套餐B外包给工厂B负责生产。两个工厂根据这家店的需求，实习那了生产食品和生产饮料的方法。A工厂就负责A套餐，那么他就需要实现生产肉的逻辑和生产CoCo的逻辑即可.而B工厂只需要实现生产汉堡和生产茶的逻辑即可.<br>

&emsp;&emsp;这样以来，来到店里的客人，只需要订购套餐，服务员通知工厂生产并送达即可。假设一个客人要套餐A，服务员通知工厂A，先生产一个肉，再来一杯CoCo，服务员负责把这些产品递给客人食用即可。<br>

```go
/***
 *	Abstract Factory
 */

type HJXFactory interface{
	CreateFood()Food
	CreateDrink()Drink
}

type FactoryA struct {

}

func (af FactoryA)CreateFood()Food{
	f := Meat{}
	return f
}

func (af FactoryA)CreateDrink()Drink{
	d := CoCo{}
	return d
}

type FactoryB struct {

}

func (bf FactoryB)CreateFood()Food{
	f := Hamberger{}
	return f
}

func (bf FactoryB)CreateDrink()Drink{
	d := Tea{}
	return d
}
type Food interface {
	Eat()
}

type Meat struct {
}

type Hamberger struct {
}

func (m Meat) Eat() {
	fmt.Println("Eat meat.")
}

func (h Hamberger) Eat() {
	fmt.Println("Eat Hamberger.")
}

type Drink interface{
	Drink()
}

type CoCo struct{

}

func (cc CoCo)Drink(){
	fmt.Println("Drink CoCo")
}

type Tea struct{

}

func (t Tea)Drink(){
	fmt.Println("Drink Tea")
}
func main(){
	// Abstract Factory
	fa := FactoryA{}
	fa.CreateFood().Eat()
	fa.CreateDrink().Drink()

	fb := FactoryB{}
	fb.CreateFood().Eat()
	fb.CreateDrink().Drink()
}
```

<br>

<center>

Author: 寒江雪<br>

Date:2018-03-08<br>

</center>