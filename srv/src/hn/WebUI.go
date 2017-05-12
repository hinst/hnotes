package hn

import (
	"net/http"
	"fmt"
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
	this.InstallFileHandler("")
	this.InstallFileHandler("/static/css")
	this.InstallFileHandler("/static/js")
	this.InstallFileHandler("/static/media")
}

func (this *TWebUI) AddHandler(subUrl string, function func(response http.ResponseWriter, request *http.Request)) {
	var url = this.RootURL + subUrl
	http.HandleFunc(url, function)
}

func (this *TWebUI) InstallFileHandler(subDir string) {
	var directoryPath = AppDir + "/../ui/build" +subDir
	var url = this.RootURL + subDir + "/"
	var fileDirectory = http.Dir(directoryPath)
	var fileServerHandler = http.FileServer(fileDirectory)
	fmt.Println(url + " -> " + directoryPath)
	http.Handle(url, http.StripPrefix(url, fileServerHandler))
}