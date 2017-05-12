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

}

func (this *TWebUI) AddHandlers() {

}

func (this *TWebUI) AddHandler(subUrl string, function func(response http.ResponseWriter, request *http.Request)) {
	var url = this.RootURL + subUrl
	http.HandleFunc(url, function)
}

func (this *TWebUI) InstallFileHandler(folderName string) {
	var url = this.RootURL + "/" + folderName + "/"
	var directoryPath = this.Directory + "/" + folderName
	GlobalLog.Write("'" + url + "'->'" + directoryPath + "'")
	var fileDirectory = http.Dir(directoryPath)
	var fileServerHandler = http.FileServer(fileDirectory)
	http.Handle(url, http.StripPrefix(url, fileServerHandler))
}