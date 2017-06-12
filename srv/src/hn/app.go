package hn

import "os"
import "path/filepath"
import "sync"
import "net/http"
import "huser"

type TApp struct {
	AppUrl  string
	Holder  sync.WaitGroup
	WebUI   *TWebUI
	UserMan *huser.TUserMan
	DataMan *TDataMan
}

var AppDir, _ = filepath.Abs(filepath.Dir(os.Args[0]))

func (this *TApp) Create() *TApp {
	this.AppUrl = "/hnotes"
	return this
}

func (this *TApp) Run() {
	this.Holder.Add(1)
	this.RunStart()
	go this.startWebServer()
	InstallShutdownReceiver(this.ReceiveStopSignal)
	this.Holder.Wait()
	this.RunStop()
}

func (this *TApp) startWebServer() {
	var result = http.ListenAndServe(":9001", nil)
	AssertResult(result)
}

func (this *TApp) ReceiveStopSignal() {
	this.Holder.Done()
}

func (this *TApp) RunStart() {
	this.DataMan = (&TDataMan{}).Create()
	this.DataMan.Start()
	this.UserMan = (&huser.TUserMan{
		AppUrl:   this.AppUrl,
		FilePath: AppDir + "/data/users.db",
	}).Create()
	this.UserMan.Start()
	this.WebUI = (&TWebUI{
		AppUrl:  this.AppUrl,
		DataMan: this.DataMan,
		UserMan: this.UserMan,
	}).Create()
	this.WebUI.Start()
}

func (this *TApp) RunStop() {
	this.WebUI.Stop()
	this.UserMan.Stop()
	this.DataMan.Stop()
}
