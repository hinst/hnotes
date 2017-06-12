package huser

import (
	"sync"
	"net/http"
	"sync/atomic"
)

type TUserMan struct {
	Blocked int32
	AppUrl string
	FilePath string
	WebWaitGroup *sync.WaitGroup
	DataMan TDataMan
}

func (this *TUserMan) Create() *TUserMan {
	this.DataMan.FilePath = this.FilePath
	return this
}

func (this *TUserMan) Start() {

}

func (this *TUserMan) Stop() {
	atomic.AddInt32(&this.Blocked, 1)
	this.WebWaitGroup.Wait()
}

func (this *TUserMan) AddHandler(subUrl string, function func(response http.ResponseWriter, request *http.Request)) {
	var url = this.AppUrl + subUrl
	http.HandleFunc(url, func(response http.ResponseWriter, request *http.Request) {
		this.WrapHandler(response, request, function)
	})
}

func (this *TUserMan) WrapHandler(response http.ResponseWriter, request *http.Request,
	function func(response http.ResponseWriter, request *http.Request),
) {
	this.WebWaitGroup.Add(1)
	defer this.WebWaitGroup.Done()
	if atomic.CompareAndSwapInt32(&this.Blocked, 1, 1) { return }
	response.Header().Set("Access-Control-Allow-Origin", "*")
	function(response, request)
}


