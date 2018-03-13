## Chain of Responsibility Pattern

&emsp;&emsp;责任链模式为请求创建了一个接收者对象的链。处理请求的对象是一系列链条，或者一个集合。每个处理节点都保存有其他处理节点的引用。接收到请求的时候，会判断该请求是否需要自己处理，处理完毕传给下一个处理节点，有点像流水线作业。<br>

## 例子

&emsp;&emsp;让我举一个复杂一点的例子。我就不用链条的形式来写了，我用一个map来保存处理器的引用。<br>

&emsp;&emsp;单机情况下模拟一个Request的发送。<br>

&emsp;&emsp;结构如下:<br>

![](https://raw.githubusercontent.com/lkysyzxz/pictureForMD/csdn_blog_golang/chain%20of%20response%20pattern.png)

&emsp;&emsp; 这个系统中,定义一个IProcess接口，每一个处理节点都应该实现这个接口。<br>

&emsp;&emsp;创建一个Processer结构实现这个接口，这个结构也作为父类被RootProcesser和PostProcesser结构继承。<br>

&emsp;&emsp;RootProcesser作为根节点，它的职责是把收到的请求，根据请求的方法传递给相应方法的处理节点，这里只有Post方法。<br>

&emsp;&emsp; PostProcesser收到请求后，根据Url和多路复用器，将请求传递给相应Url的处理函数。<br>

&emsp;&emsp;处理函数在Mux结构中注册，以便被调用。<br>

下面是关于Request的定义<br>

```go
package chainofresponsibility


type Values map[string][]string

type Request struct {
	Method   string
	Url		 string
	PostForm Values
}

func MakeRequest() *Request {
	r := Request{}
	r.PostForm = make(Values)
	return &r
}

func SendRequst(request *Request,mux *Mux){
	mux.root.Process(request)
}

func (this *Request) SetValues(key string, values ...string) {
	this.PostForm[key] = append(this.PostForm[key], values...)
}

func (this *Request) Post(url string,mux *Mux) {
	this.Method = "POST"
	this.Url=url
	SendRequst(this,mux)
}
```

<br>

下面是处理器和多路复用器的定义<br>

```go
package chainofresponsibility

type IProcess interface {
	Process(request *Request)
}

type HandlersCollection map[string]IProcess

type Processer struct {
	Handlers HandlersCollection
}

func (this *Processer) SetHandler(key string, process IProcess) {
	this.Handlers[key] = process
}

func (this *Processer) Init() {
	this.Handlers = make(HandlersCollection)
}

type RootProcesser struct {
	Processer
}

func (this *RootProcesser) Process(request *Request) {
	this.Handlers[request.Method].Process(request)
}

func newRootProcesser() *RootProcesser {
	root := new(RootProcesser)
	root.Processer.Init()
	return root
}

type PostProcesser struct {
	Processer
	PMux *Mux
}

func newPostProcesser(mux *Mux) *PostProcesser {
	post := new(PostProcesser)
	post.Processer.Init()
	post.PMux = mux
	return post
}

func (this *PostProcesser) Process(request *Request) {
	this.PMux.mux[request.Url](request)
}

type HandlerFunc func(request *Request)
type muxEntry map[string]HandlerFunc

type Mux struct {
	mux  muxEntry
	root IProcess
}

func (this *Mux) Handle(url string, handlerFunc HandlerFunc) {
	this.mux[url] = handlerFunc
}

func (this *Mux) SetRootProcess(root IProcess) {
	this.root = root;
}

func NewMux() *Mux {
	mux := Mux{}
	mux.mux = make(muxEntry)

	root := newRootProcesser()
	post := newPostProcesser(&mux)
	root.SetHandler("POST", post)

	mux.root = root;
	return &mux
}

```

<br>

下面是main<br>

```go
package main

import (
	"projects/DesignPatternsByGo/behavioralPatterns/chainofresponsibility"
	"fmt"
)

func Login(request *chainofresponsibility.Request){
	username := request.PostForm["username"][0]
	password := request.PostForm["password"][0]

	fmt.Println(username)
	fmt.Println(password)
}

func main(){
	mux := chainofresponsibility.NewMux()

	mux.Handle("login",Login)


	req := chainofresponsibility.MakeRequest()
	req.Method="POST"
	req.SetValues("username","111")
	req.SetValues("password","222")
	req.Post("login",mux)
}
```

<br>

## 优点

* 降低耦合度。它将请求的发送者和接收者解耦。
* 简化了对象。使得对象不需要知道链的结构。
* 增强给对象指派职责的灵活性。通过改变链内的成员或者调动它们的次序，允许动态地新增或者删除责任。
* 增加新的请求处理类很方便。

## 缺点

* 不能保证请求一定被接收。
* 系统性能将受到一定影响，而且在进行代码调试时不太方便，可能会造成循环调用。 
* 可能不容易观察运行时的特征，有碍于除错。<br>