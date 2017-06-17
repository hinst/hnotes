package main

import "fmt"
import "hn"

func main() {
	fmt.Println("STARTING...")
	var app = (&hn.TApp{}).Create()
	app.Run()
	fmt.Println("EXITING...")
}
