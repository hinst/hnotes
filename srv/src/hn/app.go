package hn

import "os"
import "path/filepath"
import "sync"
import "net/http"

type TApp struct {
	Holder sync.WaitGroup
	WebUI *TWebUI
	DataMan *TDataMan
}

var AppDir, _ = filepath.Abs(filepath.Dir(os.Args[0]))

func (this *TApp) Run() {
	this.Holder.Add(1)
	this.DataMan = (&TDataMan{}).Create()
	this.DataMan.Start()
	this.WebUI = (&TWebUI{DataMan: this.DataMan}).Create()
	this.WebUI.Start()
	go http.ListenAndServe(":9001", nil)
	InstallShutdownReceiver(this.stop)
	this.Holder.Wait()
	this.DataMan.Stop()
}

func (this *TApp) stop() {
	this.Holder.Done()
}