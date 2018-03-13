## Proxy Pattern

&emsp;&emsp;代理模式使得一个对象可以给另一个对象提供访问控制。截取所有访问。<br>

## 实现

&emsp;&emsp;代理模式可以使用在很多地方，例如网络连接，内存中大的对象，一个文件，或者其他消耗大的对象，或者是不可能被复制的对象。<br>

下面是一个简单的例子。<br>

```go
package proxy

import (
	"fmt"
	"sync"
)

type IObject interface {
	ObjDo(action string)
}

type Object struct {
	action string
}

func (this *Object) ObjDo(action string) {
	fmt.Println("I can do:" + action)
}

type ProObject struct {
	obj *Object
}

var one  = new(sync.Once)
func (this *ProObject) ObjDo(action string) {
	one.Do(func() {
		if this.obj==nil{
			this.obj=new(Object)
		}
	})
	this.obj.ObjDo(action)
}
```

<br>

```go
func main(){
	proxy := new(proxy.ProObject)
	proxy.ObjDo("Well")
}
```

<br>