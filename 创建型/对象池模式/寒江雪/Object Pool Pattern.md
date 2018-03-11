## Object Pool Pattern

&emsp;&emsp;对象池模式是一种创建型模式，根据需求来预测将要使用的对象，提前创建并保存在内存中。<br>

## 实现

### 对象池定义

```go
package pool

import (
	"fmt"
	"strconv"
)

type Object struct{

}

func (Object)Do(index int){
	fmt.Println("Object Do:"+strconv.Itoa(index))
}


type Pool chan *Object

func NewPool(total int)*Pool{
	p := make(Pool,total)
	for i := 0;i<total;i++{
		p <- new(Object)
	}
	return &p
}
```

<br>

### 使用

```go
func main(){
	p := pool.NewPool(5)
	wait := sync.WaitGroup{}
	for i:=0;i<100;i++ {
		index := i
		wait.Add(1)
		go func(pool pool.Pool, ind int) {
			select {
			case Obj := <-pool:
				Obj.Do(ind)
				pool <- Obj
			default:
				fmt.Println("No Object:"+strconv.Itoa(ind))
			}
			wait.Done()
		}(*p, index)
	}
	wait.Wait()
}
```

<br>

&emsp;&emsp;这里使用goroutines来并发地读取pool中的对象.<br>

## 注意

* 当创建对象的代价比维护代价更高的时候，使用对象池模式是极好的。
* 如果需求相对固定，那么维护对象的代价可能得不偿失
* 提前初始化对象对性能有积极的影响<br>