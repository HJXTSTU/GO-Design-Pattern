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
