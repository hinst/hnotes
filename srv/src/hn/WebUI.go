package hn

import (
	"net/http"
	"fmt"
	"encoding/json"
	"captcha"
)

type TWebUI struct {
	RootURL string
	DataMan *TDataMan
}

func (this *TWebUI) Create() *TWebUI {
	this.RootURL = "/hnotes"
	return this
}

func (this *TWebUI) Start() {
	this.AddHandlers()
}

func (this *TWebUI) AddHandlers() {
	this.InstallFileHandler("")
	this.InstallFileHandler("/static/css")
	this.InstallFileHandler("/static/js")
	this.InstallFileHandler("/static/media")
	this.AddHandler("/notes", this.GetNotes)
	this.AddHandler("/getCaptcha", this.GetCaptcha)
	http.Handle(this.RootURL + "/captcha/", captcha.Server(captcha.StdWidth, captcha.StdHeight))
}

func (this *TWebUI) AddHandler(subUrl string, function func(response http.ResponseWriter, request *http.Request)) {
	var url = this.RootURL + subUrl
	http.HandleFunc(url, func(response http.ResponseWriter, request *http.Request) {
		this.WrapHandler(response, request, function)
	})
}

func (this *TWebUI) WrapHandler(response http.ResponseWriter, request *http.Request,
	function func(response http.ResponseWriter, request *http.Request),
) {
	response.Header().Set("Access-Control-Allow-Origin", "*")
	function(response, request)
}

func (this *TWebUI) InstallFileHandler(subDir string) {
	var directoryPath = AppDir + "/../ui/build" +subDir
	var url = this.RootURL + subDir + "/"
	var fileDirectory = http.Dir(directoryPath)
	var fileServerHandler = http.FileServer(fileDirectory)
	fmt.Println(url + " -> " + directoryPath)
	http.Handle(url, http.StripPrefix(url, fileServerHandler))
}

func (this *TWebUI) GetNotes(response http.ResponseWriter, request *http.Request) {
	var notes = GetSampleNoteArray()
	var data, _ = json.Marshal(&notes)
	response.Write(data)
}

func (this *TWebUI) GetCaptcha(response http.ResponseWriter, request *http.Request) {
	var id = captcha.New()
	response.Write([]byte(id))
}

func (this *TWebUI) RegisterNewUser(response http.ResponseWriter, request *http.Request) {
	var args struct {captchaId, captcha, user, password string}
	if json.NewDecoder(request.Body).Decode(&args) == nil {
		if captcha.VerifyString(args.captchaId, args.captcha) {
			this.DataMan.RegisterUser(TUser{name: args.user, password: args.password})
		}
	}
}
