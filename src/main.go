package main

import (
	"fmt"
	"time"
	"juego"
)

func main() {
	j := new(juego.Juego)
	j.Fill()
	j.FillInitTest()
	for {
		fmt.Println(j)
		j.Check()
		time.Sleep(juego.TIEMPO_ACTUALIZACION)
	}
}
