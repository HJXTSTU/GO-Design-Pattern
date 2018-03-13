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
