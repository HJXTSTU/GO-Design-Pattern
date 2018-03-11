##  Singleton Pattern

&emsp;&emsp;单例模式限制了一个类型只有一个对象。<br>

## 实现

###  单例定义

```go
package singleton

type singleton map[string]string

var (
    once sync.Once

    instance singleton
)

func New() singleton {
    once.Do(func() {
        instance = make(singleton)
    })

    return instance
}
```

<br>

### 使用

```go
func main(){
	instance_1 := singleton.GetInstance()
	instance_1["this"]="that"

	instance_2 := singleton.GetInstance()
	s := instance_2["this"]
	fmt.Println(s)
}
```

<br>

<center>

Author:寒江雪<br>

Date:2018 03 10<br>

</center>