package hn

import "os"
import "path/filepath"

type TApp struct {
}

var AppDir, _ = filepath.Abs(filepath.Dir(os.Args[0]))

func (this *TApp) Run() {

}