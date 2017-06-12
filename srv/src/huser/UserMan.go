package huser

import (
	"sync"
	"net/http"
	"sync/atomic"
	"captcha"
	"encoding/json"
)

type TUserMan struct {
	WebBlocked int32
	AppUrl string
	FilePath string
	WebWaitGroup *sync.WaitGroup
	DataMan *TDataMan
}

func (this *TUserMan) Create() *TUserMan {
	this.DataMan = (&TDataMan{}).Create()
	this.DataMan.FilePath = this.FilePath
	return this
}

func (this *TUserMan) Start() {
	this.DataMan.Start()
	this.AddHandler("/getCaptcha", this.GetCaptcha)
	this.AddHandler("/registerNewUser", this.RegisterNewUser)
	this.AddHandler("/login", this.Login)
	http.Handle(this.AppUrl + "/captcha/", captcha.Server(captcha.StdWidth, captcha.StdHeight))
}

func (this *TUserMan) Stop() {
	atomic.AddInt32(&this.WebBlocked, 1)
	this.WebWaitGroup.Wait()
	this.DataMan.Stop()
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
	if atomic.CompareAndSwapInt32(&this.WebBlocked, 1, 1) { return }
	response.Header().Set("Access-Control-Allow-Origin", "*")
	function(response, request)
}

func (this *TUserMan) GetCaptcha(response http.ResponseWriter, request *http.Request) {
	var id = captcha.New()
	response.Write([]byte(id))
}

func (this *TUserMan) RegisterNewUser(response http.ResponseWriter, request *http.Request) {
	var args struct {CaptchaId, Captcha, User, Password string}
	var responseObject struct { CaptchaSuccess, Success bool }
	if json.NewDecoder(request.Body).Decode(&args) == nil {
		if captcha.VerifyString(args.CaptchaId, args.Captcha) {
			responseObject.CaptchaSuccess = true
			responseObject.Success = this.DataMan.RegisterUser(TUser{name: args.User, password: args.Password})
		}
	}
	response.Write(JsonMarshal(&responseObject))
}

func (this *TWebUI) Login(response http.ResponseWriter, request *http.Request) {
	var args struct { User, Password string }
	var responseObject struct { SessionKey string }
	if json.NewDecoder(request.Body).Decode(&args) == nil {
		var user = TUser{name: args.User, password: args.Password}
		responseObject.SessionKey = this.DataMan.Login(user)
	}
	response.Write(JsonMarshal(&responseObject))
}
