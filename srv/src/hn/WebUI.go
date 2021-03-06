package hn

import (
	"encoding/json"
	"fmt"
	"huser"
	"net/http"
	"sync"
	"sync/atomic"
)

type TWebUI struct {
	Blocked   int32
	WaitGroup sync.WaitGroup
	AppUrl    string
	DataMan   *TDataMan
	UserMan   *huser.TUserMan
}

func (this *TWebUI) Create() *TWebUI {
	return this
}

func (this *TWebUI) Start() {
	this.AddHandlers()
}

func (this *TWebUI) Stop() {
	atomic.AddInt32(&this.Blocked, 1)
	this.WaitGroup.Wait()
}

func (this *TWebUI) AddHandlers() {
	this.InstallFileHandler("")
	this.InstallFileHandler("/static/css")
	this.InstallFileHandler("/static/js")
	this.InstallFileHandler("/static/media")
	this.AddHandler("/notes", this.GetNotes)
}

func (this *TWebUI) AddHandler(subUrl string, function func(response http.ResponseWriter, request *http.Request)) {
	var url = this.AppUrl + subUrl
	http.HandleFunc(url, func(response http.ResponseWriter, request *http.Request) {
		this.WrapHandler(response, request, function)
	})
}

func (this *TWebUI) WrapHandler(response http.ResponseWriter, request *http.Request,
	function func(response http.ResponseWriter, request *http.Request),
) {
	this.WaitGroup.Add(1)
	defer this.WaitGroup.Done()
	if atomic.CompareAndSwapInt32(&this.Blocked, 1, 1) {
		return
	}
	response.Header().Set("Access-Control-Allow-Origin", "*")
	function(response, request)
}

func (this *TWebUI) InstallFileHandler(subDir string) {
	var directoryPath = AppDir + "/../ui/build" + subDir
	var url = this.AppUrl + subDir + "/"
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
