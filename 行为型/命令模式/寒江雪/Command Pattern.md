## Command Pattern

&emsp;&emsp;命令模式将命令包裹在对象中，并传给调用对象。调用对象寻找处理该命令的合适对象，并把该命令传给合适的对象，该对象执行命令。<br>

## 例子

&emsp;&emsp;举一个例子，在go里边调用C语言的代码，来监听键盘的输入，并调用与输入绑定的函数。<br>

&emsp;&emsp;定义事件码和事件结构。<br>

```go
const (
	EVENT_CODE_KEY = iota
)

type EventCode rune
type EventData interface{}
type Event struct{
	Code EventCode
	Data EventData
}
```

&emsp;&emsp;为了简单，只写一个按键的事件。<br>

&emsp;&emsp;定义一个抽象的接口。<br>

```go
type Command interface {
	Execute(EventData)
}
```

&emsp;&emsp;这个接口定义了命令的形式。<br>

&emsp;&emsp;我们想让用户可以在外部传入回调函数，那么就需要定义一个接收事件的回调函数。<br>

```go
type EventFunc func(Event)
```

&emsp;&emsp;有了回调函数形式，我们还需要一个对象，包含这种函数。这种函数不是Command接口对象。但我们可以把它包装成Command对象，这里就用到了适配器模式。<br>

```go
type EventFunder struct {
	f EventFunc
}

func (this EventFunder) Execute(data Event) {
	this.f(data)
}

func eventFunder(eventFunc EventFunc) EventFunder {
	return EventFunder{eventFunc}
}
```

&emsp;&emsp;eventFunder就是把一个EventFunc类型的对象转换为EventFunder,EventFunder实现了Command接口，它属于Command。<br>

&emsp;&emsp;最后是EventSystem的实现。<br>

```go
type Hash map[EventCode]Command

type EventSystem struct {
	hash Hash
}

func (this *EventSystem) Init() *EventSystem {
	this.hash = make(Hash)
	return this
}

func (this *EventSystem) Map(code EventCode, eventFunc EventFunc) {
	this.hash[code] = eventFunder(eventFunc)
}

func (this *EventSystem) InspectKeyboard() {
	go func() {
		C.init_keyboard()
		for {
			if C.kbhit() > 0 {
				ch := EventCode(C.readch())
				// TODO::generate key event
				this.generateEvent(Event{EVENT_CODE_KEY, ch})
			}
		}
	}()
}

func (this *EventSystem) generateEvent(e Event) {
	v, ok := this.hash[e.Code]
	if ok {
		v.Execute(e)
	}
}

func NewEventSystem() *EventSystem {
	return (&EventSystem{}).Init()
}
```

&emsp;&emsp;EventSystem实际上是一个调用者，它只接收Command对象。所以要把传入的函数做适配。<br>

## 优点

* 降低了系统耦合度。
* 新的命令可以很容易添加到系统中去。

## 缺点

* 使用命令模式可能会导致某些系统有过多的具体命令类。

## 注意事项

&emsp;&emsp;系统需要支持命令的撤销(Undo)操作和恢复(Redo)操作，也可以考虑使用命令模式，见命令模式的扩展。<br>

&emsp;&emsp;