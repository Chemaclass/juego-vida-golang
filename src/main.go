package main

import (
	"fmt"
	"juego"
)

func main() {
	j := new(juego.Juego)
	j.Fill()
	j.FillInitTest()
	go j.Run()
	end()
}

func end() {
	var s string
	fmt.Scanf("%s", &s)
	fmt.Println("End")
}
