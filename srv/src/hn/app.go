package hn

import "os"
import "path/filepath"
import "sync"
import "net/http"

type TApp struct {
	Holder sync.WaitGroup
	WebUI *TWebUI
}

var AppDir, _ = filepath.Abs(filepath.Dir(os.Args[0]))

func (this *TApp) Run() {
	this.Holder.Add(1)
	this.WebUI = (&TWebUI{}).Create()
	this.WebUI.Start()
	go http.ListenAndServe(":9001", nil)
	InstallShutdownReceiver(this.stop)
	this.Holder.Wait()
}

func (this *TApp) stop() {
	this.Holder.Done()
}