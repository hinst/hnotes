package hn

import (
	"net/http"
)

type TWebUI struct {
	RootURL string
}

func (this *TWebUI) Create() *TWebUI {
	this.RootURL = "/hnotes"
	return this
}

func (this *TWebUI) Start() {
	this.AddHandlers()
}

func (this *TWebUI) AddHandlers() {
	this.InstallFileHandler()
}

func (this *TWebUI) AddHandler(subUrl string, function func(response http.ResponseWriter, request *http.Request)) {
	var url = this.RootURL + subUrl
	http.HandleFunc(url, function)
}

func (this *TWebUI) InstallFileHandler() {
	var directoryPath = AppDir + "/ui/build"
	var fileDirectory = http.Dir(directoryPath)
	var fileServerHandler = http.FileServer(fileDirectory)
	http.Handle(this.RootURL, http.StripPrefix(this.RootURL, fileServerHandler))
}